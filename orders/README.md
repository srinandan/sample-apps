# sample-apps-orders

## Usage

This sample application listens on port `8080` and exposes a basePath `/orders`. The application supports GET, LIST and CREATE items.

```bash

curl localhost:8080/items/100000

{
    "operationId": "100000",
    "shipmentId": "shipment-1",
    "lineItems": [
        {
        "lineItemId": "999999",
        "quantity": 1
        }
    ],
    "carrier": "FedEx",
    "trackingId": "200000"
}
```

