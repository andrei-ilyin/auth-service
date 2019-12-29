CXX = g++
CXXFLAGS += -std=c++17
CPPFLAGS += `pkg-config --cflags protobuf grpc grpc libmongocxx`
LDFLAGS += `pkg-config --libs protobuf grpc++ grpc libmongocxx`	

PROTOC = protoc
GRPC_CPP_PLUGIN = grpc_cpp_plugin
GRPC_CPP_PLUGIN_PATH ?= `which $(GRPC_CPP_PLUGIN)`

PROTOS_PATH = .

all: auth_server auth_client

auth_server: auth.pb.o auth.grpc.pb.o auth_server.o
	$(CXX) $^ $(LDFLAGS) -o $@

auth_client: auth.pb.o auth.grpc.pb.o auth_client.o
	$(CXX) $^ $(LDFLAGS) -o $@

.PRECIOUS: %.grpc.pb.cc
%.grpc.pb.cc: %.proto
	$(PROTOC) -I $(PROTOS_PATH) --grpc_out=. --plugin=protoc-gen-grpc=$(GRPC_CPP_PLUGIN_PATH) $<

.PRECIOUS: %.pb.cc
%.pb.cc: %.proto
	$(PROTOC) -I $(PROTOS_PATH) --cpp_out=. $<

clean:
	rm -f *.o *.pb.cc *.pb.h auth_server auth_client
