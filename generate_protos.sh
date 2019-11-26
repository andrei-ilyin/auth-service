#!/usr/bin/env bash

protoc --cpp_out=. *.proto
protoc --grpc_out=. --plugin=protoc-gen-grpc=$(which grpc_cpp_plugin) *.proto
