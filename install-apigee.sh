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

# create a sample developer
apigeecli developers create -o $1 -n apps@sample.com -u faziodev -f fazio -s user -a $3
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to import api proxy bundle"
  exit 1
fi

# create a inventory product
apigeecli products create -o $1 -f auto -n inventory_product -m "Inventory Product" -e $2 -p inventory -d "A product for inventory app" -a $3
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to create product"
  exit 1
fi

# create inventory app
apigeecli apps create -o $1 -n inventory_app -p inventory_product -e apps@sample.com -a $3
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to create app"
  exit 1
fi

# create orders product
apigeecli products create -o $1 -f auto -n orders_product -m "Orders Product" -e $2 -p orders -d "A product for orders app" -a $3
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to create product"
  exit 1
fi

# create orders app
apigeecli apps create -o $1 -n orders_app -p orders_product -e apps@sample.com -a $3
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to create app"
  exit 1
fi

# create tracking product
apigeecli products create -o $1 -f auto -n tracking_product -m "Tracking Product" -e $2 -p tracking -d "A product for tracking app" -a $3
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to create product"
  exit 1
fi

# create orders app
apigeecli apps create -o $1 -n tracking_app -p tracking_product -e apps@sample.com -a $3
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to create app"
  exit 1
fi

# websockets app
# zip the bundle
cd websockets && zip -r ../websockets.zip apiproxy && cd ..
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to zip bundle"
  exit 1
fi

# import the api proxy bundle
apigeecli apis create -o $1 -n websockets -p websockets.zip -a $3
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to import api proxy bundle"
  exit 1
fi

# create an api product
apigeecli products create -o $1 -f auto -n websocket_product -m "Websockets Product" -e $2 -p websockets -d "A sample products for websockets"
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to create an api product"
  exit 1
fi

# create a developer
apigeecli developers create -o $1 -n websocket-dev@sample.com -u websocket-dev -f sample -s user
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to create developer"
  exit 1
fi

# create a sample app
apigeecli apps create -o $1 -n websocket-app -p websocket_product -e websocket-dev@sample.com
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to create developer app"
  exit 1
fi