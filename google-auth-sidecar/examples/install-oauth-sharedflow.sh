#!/bin/bash
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


./check-prereqs.sh $1 $2 $3
RESULT=$?
if [ $RESULT -ne 0 ]; then
  exit 1
fi

echo "passed prerequisite checks..."

# zip the bundle
cd oauth-sharedflow && zip -r ../oauth-sharedflow.zip sharedflowbundle && cd ..

# import the shared flow
apigeecli sharedflows create -o $1 -n oauth-sharedflow -p oauth-sharedflow.zip -a $3

# deploy the shared flow
apigeecli sharedflows deploy -o $1 -e $2 -n oauth-sharedflow -v 1

# delete the zipped bundle
rm -rf oauth-sharedflow.zip