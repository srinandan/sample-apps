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


echo $1 
echo $2
echo $3

if [ -z "$1" ]
  then
    echo "org name is a mandatory parameter. Usage: 'check-preqs.sh {org} {env} {path-to-service-account.json}'"
    exit 1
fi

if [ -z "$2" ]
  then
    echo "env name is a mandatory parameter. Usage: 'check-preqs.sh {org} {env} {path-to-service-account.json}'"
    exit 1
fi

if [ -z "$3" ]
  then
    echo "service account is a mandatory parameter. Usage: 'check-preqs.sh {org} {env} {path-to-service-account.json}'"
    exit 1
fi

apigeecli 2>&1 >/dev/null
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "this script depends on apigeecli (https://github.com/srinandan/apigeecli)"
  exit 1
fi

zip --help 2>&1 >/dev/null
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "this script depends on the zip utility. Please install zip and re-run the command"
  exit 1
fi

exit 0
