CXX = clang++
CXXFLAGS += -std=c++14 -O2 -Wall -Wextra -Wno-unused-parameter

PROTOC = protoc
GRPC_CPP_PLUGIN = grpc_cpp_plugin
GRPC_CPP_PLUGIN_PATH ?= `which $(GRPC_CPP_PLUGIN)`

PROTOS_PATH = .

all: auth_server auth_client

auth_client: auth.pb.o auth.grpc.pb.o auth_client.o
	$(CXX) $^ $(LDFLAGS) `pkg-config --libs protobuf grpc++ grpc` -o $@

auth_server: auth.pb.o auth.grpc.pb.o auth_server.o database.o rpc_server_impl.o
	$(CXX) $^ $(LDFLAGS) `pkg-config --libs protobuf grpc++ grpc libmongocxx` -o $@

auth_client.o: auth_client.cc auth.pb.cc auth.grpc.pb.cc
	$(CXX) -c $< $(CXXFLAGS) `pkg-config --cflags protobuf grpc++ grpc` -o $@

auth_server.o: auth_server.cc auth.pb.cc auth.grpc.pb.cc
	$(CXX) -c $< $(CXXFLAGS) `pkg-config --cflags protobuf grpc++ grpc libmongocxx` -o $@

database.o: database.cc auth.pb.cc auth.grpc.pb.cc
	$(CXX) -c $< $(CXXFLAGS) `pkg-config --cflags protobuf libmongocxx` -o $@

rpc_server_impl.o: rpc_server_impl.cc auth.pb.cc auth.grpc.pb.cc
	$(CXX) -c $< $(CXXFLAGS) `pkg-config --cflags protobuf grpc++ grpc libmongocxx` -o $@

.PRECIOUS: %.grpc.pb.cc
%.grpc.pb.cc: %.proto
	$(PROTOC) -I $(PROTOS_PATH) --grpc_out=. --plugin=protoc-gen-grpc=$(GRPC_CPP_PLUGIN_PATH) $<

.PRECIOUS: %.pb.cc
%.pb.cc: %.proto
	$(PROTOC) -I $(PROTOS_PATH) --cpp_out=. $<

clean:
	rm -f *.o *.pb.cc *.pb.h auth_server auth_client
