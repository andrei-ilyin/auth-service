#include <iostream>
#include <memory>
#include <string>

#include <bsoncxx/builder/stream/document.hpp>
#include <bsoncxx/json.hpp>
#include <grpcpp/grpcpp.h>
#include <mongocxx/client.hpp>
#include <mongocxx/instance.hpp>
#include <mongocxx/pool.hpp>
#include <utility>

#include "auth.pb.h"
#include "auth.grpc.pb.h"

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::StatusCode;
using bsoncxx::builder::basic::kvp;

class DatabaseInterface {
 public:
  virtual std::string FindUser(const std::string& name,
                               const std::string& password) = 0;

  virtual auth::Cookie CreateCookie(const std::string& user_id) = 0;

  virtual void RemoveCookie(const auth::Cookie& cookie) = 0;

  virtual auth::Status_Code ValidateAccess(const auth::Cookie& cookie,
                                           const std::string& resource) = 0;
};

class FakeDatabase : public DatabaseInterface {
 public:
  std::string FindUser(const std::string& name,
                       const std::string& password) override {
    return std::string("__user_id__");
  }

  auth::Cookie CreateCookie(const std::string& user_id) override {
    auth::Cookie cookie;
    cookie.set_session_id("__session_id__");
    return cookie;
  }

  void RemoveCookie(const auth::Cookie& cookie) override {
  }

  auth::Status_Code ValidateAccess(const auth::Cookie& cookie,
                                   const std::string& resource) override {
    return auth::Status::ACCESS_DENIED;
  }
};

class MongoDatabase : public DatabaseInterface {
 public:
  explicit MongoDatabase(const std::string& uri, std::string db_name)
      : instance_(), pool_(mongocxx::uri(uri)), db_name_(std::move(db_name)) {
    write_concern_.majority(std::chrono::milliseconds(100));
  }

  std::string FindUser(const std::string& name,
                       const std::string& password) override {
    bsoncxx::builder::basic::document query;
    query.append(kvp("name", name));
    query.append(kvp("password", password));

    core::optional<bsoncxx::document::value> maybe_result;
    {
      auto client = pool_.acquire();
      auto database = (*client)[db_name_];
      maybe_result = database["users"].find_one(query.view());
    }

    if (!maybe_result) {
      return "";
    }

    return maybe_result.value().view()["_id"].get_oid().value.to_string();
  }

  auth::Cookie CreateCookie(const std::string& user_id) override {
    bsoncxx::builder::basic::document document;
    document.append(kvp("user_id", user_id));

    core::optional<mongocxx::result::insert_one> maybe_result;
    {
      auto client = pool_.acquire();
      auto database = (*client)[db_name_];
      database.write_concern(write_concern_);
      maybe_result = database["cookies"].insert_one(document.view());
    }

    auth::Cookie cookie;
    cookie.set_session_id(
        maybe_result->inserted_id().get_oid().value.to_string());

    return cookie;
  }

  void RemoveCookie(const auth::Cookie& cookie) override {
    bsoncxx::builder::basic::document query;
    query.append(kvp("_id", bsoncxx::oid(cookie.session_id())));

    auto client = pool_.acquire();
    auto database = (*client)[db_name_];
    database.write_concern(write_concern_);
    database["cookies"].delete_one(query.view());
  }

  auth::Status_Code ValidateAccess(const auth::Cookie& cookie,
                                   const std::string& resource) override {
    auto client = pool_.acquire();
    auto database = (*client)[db_name_];

    bsoncxx::builder::basic::document cookie_query;
    cookie_query.append(kvp("_id", bsoncxx::oid(cookie.session_id())));

    auto maybe_result = database["cookies"].find_one(cookie_query.view());
    if (!maybe_result) {
      return auth::Status::INVALID_SESSION;
    }

    const std::string& user_id =
        maybe_result->view()["user_id"].get_utf8().value.to_string();

    bsoncxx::builder::basic::document permission_query;
    permission_query.append(kvp("user_id", user_id));
    permission_query.append(kvp("resource", resource));

    if (database["permissions"].find_one(permission_query.view())) {
      return auth::Status::OK;
    } else {
      return auth::Status::ACCESS_DENIED;
    }
  }

 private:
  mongocxx::instance instance_;
  mongocxx::pool pool_;
  mongocxx::write_concern write_concern_;

  std::string db_name_;
};

class AuthenticatorServiceImpl final : public auth::Authenticator::Service {
 public:
  explicit AuthenticatorServiceImpl(std::unique_ptr<DatabaseInterface> database)
      : database_(std::move(database)) {}

 private:
  grpc::Status Login(ServerContext* context, const auth::LoginRequest* request,
                     auth::LoginResponse* response) override {
    const std::string& username = request->credentials().user_name();
    const std::string& password = request->credentials().password();

    std::string user_id = database_->FindUser(username, password);
    if (user_id.empty()) {
      response->mutable_status()->set_code(auth::Status_Code_ACCESS_DENIED);
      return grpc::Status::OK;
    }

    auth::Cookie cookie = database_->CreateCookie(user_id);
    response->mutable_status()->set_code(auth::Status_Code_OK);
    *response->mutable_cookie() = std::move(cookie);

    return grpc::Status::OK;
  }

  grpc::Status Logout(ServerContext* context,
                      const auth::LogoutRequest* request,
                      auth::LogoutResponse* response) override {
    // std::cerr << "Deleting cookie " << request->cookie().session_id() << '\n';
    database_->RemoveCookie(request->cookie());
    return grpc::Status::OK;
  }

  grpc::Status Validate(ServerContext* context,
                        const auth::ValidationRequest* request,
                        auth::ValidationResponse* response) override {
    response->mutable_status()->set_code(database_->ValidateAccess(
        request->cookie(), request->resource()));
    return grpc::Status::OK;
  }

  std::unique_ptr<DatabaseInterface> database_;
};

void RunServer(const std::string& listener_address,
               const std::string& database_server_uri,
               const std::string& database_name) {
  // AuthenticatorServiceImpl service(std::make_unique<FakeDatabase>());

  AuthenticatorServiceImpl service(
      std::make_unique<MongoDatabase>(database_server_uri, database_name));

  ServerBuilder builder;
  builder.AddListeningPort(listener_address, grpc::InsecureServerCredentials());
  builder.RegisterService(&service);

  std::unique_ptr<Server> server(builder.BuildAndStart());

  std::cout << "Server listening on " << listener_address << std::endl;
  server->Wait();
  std::cout << "Server is shutting down..." << std::endl;
}

int main(int argc, char* argv[]) {
  RunServer(
      "0.0.0.0:50051",
      "mongodb://localhost:27017,localhost:27018,localhost:27019/?replicaSet=rs0",
      "auth-db");
  return 0;
}
