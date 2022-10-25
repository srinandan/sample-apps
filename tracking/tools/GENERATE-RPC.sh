#!/bin/bash
#
# Copyright 2021 Google LLC. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
set -e
. tools/PROTOS.sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
mkdir -p pkg #&& java

export GOOGLEAPIS_DIR=~/local_workspace/sample-apps/tracking

for proto in ${ALL_PROTOS[@]}; do
	echo "Generating Go types for $proto"
	protoc $proto --proto_path='.' \
		-I ${GOOGLEAPIS_DIR} \
        --include_source_info --include_imports \
		--go_out='./pkg' -o service.pb	        
done
for proto in ${ALL_PROTOS[@]}; do
	echo "Generating Go gRPC for $proto"
	protoc $proto --proto_path='.' \
		-I ${GOOGLEAPIS_DIR} \
		--go-grpc_out='./pkg'
#		--java_out='./java'        
done
for proto in ${ALL_PROTOS[@]}; do
	echo "Generating gRPC-Gateway for $proto"
	protoc $proto --proto_path='.' \
		-I ${GOOGLEAPIS_DIR} \
		--grpc-gateway_out='logtostderr=true:./pkg'
done