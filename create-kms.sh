#!/bin/sh

apigeecli products create -f auto -n orders_product -m "Orders Product" -e prod1 -d "A product for orders app" -g ./orders/orders-product.json 
apigeecli products create -f auto -n inventory_product -m "Inventory Product" -e prod1 -d "a product for inventory app" -g ./inventory/inventory-product.json
apigeecli products create -f auto -n tracking_product -m "Tracking Product" -e prod1 -d "A product for tracking app" -g ./tracking/tracking-product.json
apigeecli products create -f auto -n helloworld_product -m "helloworld product" -e prod2 -p helloworld -d "A helloworld product" --legacy=true

apigeecli developers create -n apps@sample.com -u faziodev -f fazio -s user

apigeecli apps create -n tracking_app -p tracking_product -e apps@sample.com
apigeecli apps create -n inventory_app -p inventory_product -e apps@sample.com
apigeecli apps create -n orders_app -p orders_product -e apps@sample.com
