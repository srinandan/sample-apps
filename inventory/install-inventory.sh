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

proxy_name=inventory

# check if proxy exists
apigeecli apis get -o $1 -n $proxy_name
RESULT=$?
if [ $RESULT -ne 0 ]; then
    # zip the bundle
    zip -r $proxy_name.zip apiproxy
    RESULT=$?
    if [ $RESULT -ne 0 ]; then
      echo "failed to zip bundle"
      exit 1
    fi
    # import api proxy
    apigeecli apis create -o $1 -p ${proxy_name}.zip -n $proxy_name
    RESULT=$?
    if [ $RESULT -ne 0 ]; then
        echo "failed to import proxy"
        exit 1
    fi
    # deploy api proxy
    apigeecli apis deploy -o $1 -e $2 -n $proxy_name v1
    RESULT=$?
    if [ $RESULT -ne 0 ]; then
        echo "failed to deploy proxy"
        exit 1
    fi    
fi

# create inventory product
apigeecli products create -o $1 -f auto -n ${proxy_name}_product -m "Inventory Product" -e $2 -p $proxy_name -d "A product for inventory app"
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to create product"
  exit 1
fi

# create inventory app
apigeecli apps create -o $1 -n ${proxy_name}_app -p ${proxy_name}_product -e apps@sample.com
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo "failed to create app"
  exit 1
fi