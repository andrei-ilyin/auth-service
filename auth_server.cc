#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>

#include "auth.pb.h"
#include "auth.grpc.pb.h"

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;

class AuthenticatorServiceImpl final : public auth::Authenticator::Service {
  Status Login(ServerContext* context, const auth::LoginRequest* request,
               auth::LoginResponse* response) override {
    const std::string& username = request->credentials().user_name();
    const std::string& password = request->credentials().password();

    std::cout << "Login attempt with login " << username << " and password "
              << password << std::endl;

    if (username != "root" || password != "Pass_w0rd") {
      response->mutable_status()->set_code(auth::Status_Code_ACCESS_DENIED);
      return Status::OK;
    }

    // TODO(andrei_ilyin): Write actual implementation

    response->mutable_status()->set_code(auth::Status_Code_OK);
    response->mutable_cookie()->set_session_id(42);
    response->mutable_cookie()->set_hash_key("Random String Here");

    return Status::OK;
  }

  Status Logout(ServerContext* context, const auth::LogoutRequest* request,
                auth::LogoutResponse* response) override {
    // TODO(andrei_ilyin): Write actual implementation
    return Status::OK;
  }

  Status Validate(ServerContext* context,
                  const auth::ValidationRequest* request,
                  auth::ValidationResponse* response) override {
    // TODO(andrei_ilyin): Write actual implementation
    if (request->resource() != "secret") {
      response->mutable_status()->set_code(auth::Status_Code_OK);
    } else {
      response->mutable_status()->set_code(auth::Status_Code_ACCESS_DENIED);
    }

    return Status::OK;
  }
};

void RunServer() {
  std::string server_address = "0.0.0.0:50051";
  AuthenticatorServiceImpl service;

  ServerBuilder builder;
  builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
  builder.RegisterService(&service);

  std::unique_ptr<Server> server(builder.BuildAndStart());

  std::cout << "Server listening on " << server_address << std::endl;
  server->Wait();

  std::cout << "Server is shutting down..." << std::endl;
}

int main(int argc, char* argv[]) {
  RunServer();
  return 0;
}
