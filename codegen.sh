#!/usr/bin/env /bin/bash

PROTOC_PATH=$(whereis protoc | awk '{print $2}')

cd defs && $PROTOC_PATH *.proto --go_out=. --go-grpc_out=. && cd ..