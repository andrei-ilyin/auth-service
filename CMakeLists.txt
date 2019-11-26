cmake_minimum_required(VERSION 3.14)
project(rpc_service)

set(CMAKE_CXX_STANDARD 14)

set(CMAKE_CXX_FLAGS ${CMAKE_CXX_FLAGS} "-O2 -Wall -Wextra -Wno-unused-parameter")
set(CMAKE_EXE_LINKER_FLAGS ${CMAKE_EXE_LINKER_FLAGS}
        "-L/usr/local/lib -lprotobuf -pthread -lgrpc++ -lgrpc -Wl,--no-as-needed -lgrpc++_reflection -Wl,--as-needed")

add_executable(
        auth_server
        auth_server.cc auth.pb.cc auth.grpc.pb.cc
)

add_executable(
        auth_client
        auth_client.cc auth.pb.cc auth.grpc.pb.cc
)
