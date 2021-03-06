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


apigeecli developers get -o $1 -n apps@sample.com
RESULT=$?
if [ $RESULT -ne 0 ]; then
  # create a sample developer
  apigeecli developers create -o $1 -n apps@sample.com -u faziodev -f fazio -s user -a $3
  RESULT=$?
  if [ $RESULT -ne 0 ]; then
    echo "failed to create developer"
    exit 1
  fi
fi

# install orders
cd orders && ./install-orders.sh $1 $2 $3 && cd ..

# install inventory
cd inventory &&  ./install-inventory.sh  $1 $2 $3  && cd ..

# install tracking
cd tracking && ./install-tracking.sh  $1 $2 $3 && cd ..

# install websockets
cd  websockets && ./install-websockets.sh $1 $2 $3  && cd ..