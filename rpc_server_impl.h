#ifndef RPC_SERVER_IMPL_H_
#define RPC_SERVER_IMPL_H_

#include "auth.grpc.pb.h"
#include "database.h"

namespace auth {

class AuthenticatorServiceImpl final : public auth::Authenticator::Service {
 public:
  explicit AuthenticatorServiceImpl(
      std::unique_ptr<DatabaseInterface> database, bool print_log);

 private:
  grpc::Status Login(grpc::ServerContext* context,
                     const auth::LoginRequest* request,
                     auth::LoginResponse* response) override;

  grpc::Status Logout(grpc::ServerContext* context,
                      const auth::LogoutRequest* request,
                      auth::LogoutResponse* response) override;

  grpc::Status Validate(grpc::ServerContext* context,
                        const auth::ValidationRequest* request,
                        auth::ValidationResponse* response) override;

  std::unique_ptr<DatabaseInterface> database_;
  bool print_log_;
};

}  // namespace auth

#endif  // RPC_SERVER_IMPL_H_