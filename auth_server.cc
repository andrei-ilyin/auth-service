#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>

#include "database.h"
#include "rpc_server_impl.h"

int main(int argc, char* argv[]) {
  bool print_log = false;
  const std::string& listener_address = "0.0.0.0:50051";

  // auto database = std::make_unique<auth::FakeDatabase>();
  
  auto database = std::make_unique<auth::MongoDatabase>(
      "mongodb://localhost:27017,localhost:27018,localhost:27019/?replicaSet=rs0",
      "auth-db");

  auth::AuthenticatorServiceImpl service(std::move(database), print_log);

  grpc::ServerBuilder builder;
  builder.AddListeningPort(listener_address, grpc::InsecureServerCredentials());
  builder.RegisterService(&service);

  std::unique_ptr<grpc::Server> server(builder.BuildAndStart());

  std::cout << "Server listening on " << listener_address << std::endl;
  server->Wait();
  std::cout << "Server is shutting down..." << std::endl;

  return 0;
}
