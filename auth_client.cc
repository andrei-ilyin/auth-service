#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>

#include "auth.pb.h"
#include "auth.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

class AuthenticatorClient {
 public:
  explicit AuthenticatorClient(const std::shared_ptr<Channel>& channel)
      : stub_(auth::Authenticator::NewStub(channel)) {}

  bool Login(const std::string& username, const std::string& password) {
    auth::LoginRequest request;
    request.mutable_credentials()->set_user_name(username);
    request.mutable_credentials()->set_password(password);

    auth::LoginResponse response;
    ClientContext context;
    Status status = stub_->Login(&context, request, &response);

    if (!status.ok()) {
      std::cerr << "Error " << status.error_code() << ": "
                << status.error_message() << std::endl;
      return false;
    }

    if (response.status().code() != auth::Status_Code_OK) {
      std::cerr << "Login attempt failed with status "
                << auth::Status_Code_Name(response.status().code())
                << std::endl;
      return false;
    }

    cookie_ = std::move(*response.mutable_cookie());
    return true;
  }

  bool Logout() {
    auth::LogoutRequest request;
    *request.mutable_cookie() = cookie_;

    auth::LogoutResponse response;
    ClientContext context;
    Status status = stub_->Logout(&context, request, &response);

    if (!status.ok()) {
      std::cerr << "Error " << status.error_code() << ": "
                << status.error_message() << std::endl;
      return false;
    }

    return true;
  }

  bool ValidateAccess(const std::string& resource_name) {
    auth::ValidationRequest request;
    *request.mutable_cookie() = cookie_;
    request.set_resource(resource_name);

    auth::ValidationResponse response;
    ClientContext context;
    Status status = stub_->Validate(&context, request, &response);

    if (!status.ok()) {
      std::cerr << "Error " << status.error_code() << ": "
                << status.error_message() << std::endl;
      return false;
    }

    if (response.status().code() != auth::Status_Code_OK) {
      std::cerr << "Access rights verification failed for resource '"
                << resource_name << "' with status "
                << auth::Status_Code_Name(response.status().code())
                << std::endl;
      return false;
    }

    return true;
  }

  const auth::Cookie& GetCookie() const {
    return cookie_;
  }

 private:
  std::unique_ptr<auth::Authenticator::Stub> stub_;
  auth::Cookie cookie_;
};

int main(int argc, char** argv) {
  if (argc != 3) {
    std::cerr << "Usage: " << argv[0] << " {login} {password}" << std::endl;
    return 1;
  }

  AuthenticatorClient client(grpc::CreateChannel(
      "localhost:50051", grpc::InsecureChannelCredentials()));

  if (!client.Login(argv[1], argv[2])) {
    return 1;
  }

  std::cout << "Session ID = " << client.GetCookie().session_id() << '\n';

  if (client.ValidateAccess("public")) {
    std::cout << "Access to 'public' granted!\n";
  }
  if (client.ValidateAccess("secret")) {
    std::cout << "Access to 'secret' granted!\n";
  }
  if (client.ValidateAccess("unknown")) {
    std::cout << "Access to 'unknown' granted!\n";
  }

  client.Logout();

  if (client.ValidateAccess("public")) {
    std::cout << "Access to 'public' granted!\n";
  }
  if (client.ValidateAccess("secret")) {
    std::cout << "Access to 'secret' granted!\n";
  }
  if (client.ValidateAccess("unknown")) {
    std::cout << "Access to 'unknown' granted!\n";
  }

  return 0;
}