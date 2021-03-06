# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#./third_party/protoc-gen.sh
#protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 service.proto
#protoc -I/Users/srinandans/local_workspace/protoc-3.11.2-osx-x86_64/include -I/Users/srinandans/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.4/third_party/googleapis --proto_path=api/proto/v1 --proto_path=third_party  --descriptor_set_out=./third_party/service.pb --go_out=plugins=grpc:pkg/api/v1 service.proto

#git clone https://github.com/googleapis/googleapis
#export GOOGLEAPIS_DIR=<path_to_googleapis>
protoc -I${GOOGLEAPIS_DIR} --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 service.proto
protoc -I${GOOGLEAPIS_DIR} --include_imports --include_source_info --descriptor_set_out=service.pb --proto_path=api/proto/v1 --proto_path=third_party service.proto