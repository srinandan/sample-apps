# sample-apps-orders-gql

## Usage

This sample is a graphql implementation of the orders microservice. It requires the [orders](../orders) and [inventory](../inventory) services to be available.

```bash

curl localhost:8080/query -H "Content-Type: application/json" -d '{"query":"query { getOrder(operationId: 100000) { operationId shipmentId lineItems { item { inventory { salePrice { currency value } sellOnGoogleQuantity kind availability } } quantity } carrier trackingId } }"}'

{
  "data": {
    "getOrder": {
      "operationId": "100000",
      "shipmentId": "shipment-1",
      "lineItems": [
        {
          "item": {
            "inventory": {
              "salePrice": {
                "currency": "USD",
                "value": "225.72"
              },
              "sellOnGoogleQuantity": "1",
              "kind": "content#inventory",
              "availability": 0
            }
          },
          "quantity": 1
        }
      ],
      "carrier": "FedEx",
      "trackingId": "200000"
    }
  }
}
```