#include "rpc_server_impl.h"

namespace auth {

AuthenticatorServiceImpl::AuthenticatorServiceImpl(
    std::unique_ptr<DatabaseInterface> database, bool print_log)
    : database_(std::move(database)), print_log_(print_log) {}

grpc::Status AuthenticatorServiceImpl::Login(grpc::ServerContext* context,
                                             const auth::LoginRequest* request,
                                             auth::LoginResponse* response) {
  try {
    const std::string& username = request->credentials().user_name();
    const std::string& password = request->credentials().password();

    if (print_log_) {
      std::cerr << "Login attempt with username '" << username
                << "' and password '" << password << "'\n";
    }

    std::string user_id = database_->FindUser(username, password);
    if (user_id.empty()) {
      response->mutable_status()->set_code(auth::Status_Code_ACCESS_DENIED);
      return grpc::Status::OK;
    }

    if (print_log_) {
      std::cerr << "Registering cookie " << request->cookie().session_id()
                << '\n';
    }

    database_->RegisterCookie(user_id, request->cookie());
    response->mutable_status()->set_code(auth::Status_Code_OK);
    return grpc::Status::OK;

  } catch (const std::exception& exception) {
    std::cerr << "Exception at Login: " << exception.what() << '\n';
    response->mutable_status()->set_code(auth::Status_Code_INTERNAL_ERROR);
    return grpc::Status::OK;
  }
}

grpc::Status AuthenticatorServiceImpl::Logout(grpc::ServerContext* context,
                                              const auth::LogoutRequest* request,
                                              auth::LogoutResponse* response) {
  try {
    if (print_log_) {
      std::cerr << "Removing cookie " << request->cookie().session_id() << '\n';
    }

    database_->RemoveCookie(request->cookie());

    response->mutable_status()->set_code(auth::Status_Code_OK);
    return grpc::Status::OK;

  } catch (const std::exception& exception) {
    std::cerr << "Exception at Login: " << exception.what() << '\n';
    response->mutable_status()->set_code(auth::Status_Code_INTERNAL_ERROR);
    return grpc::Status::OK;
  }
}

grpc::Status AuthenticatorServiceImpl::Validate(
    grpc::ServerContext* context,
    const auth::ValidationRequest* request,
    auth::ValidationResponse* response) {
  try {
    if (print_log_) {
      std::cerr << "Validating access for session "
                << request->cookie().session_id() << " to resource '"
                << request->resource() << "'\n";
    }

    response->mutable_status()->set_code(database_->ValidateAccess(
        request->cookie(), request->resource()));
    return grpc::Status::OK;

  } catch (const std::exception& exception) {
    std::cerr << "Exception at Login: " << exception.what() << '\n';
    response->mutable_status()->set_code(auth::Status_Code_INTERNAL_ERROR);
    return grpc::Status::OK;
  }
}

}  // namespace auth