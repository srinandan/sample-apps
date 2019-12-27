# sample-apps-inventory

## Usage

This sample application listens on port `8080` and exposes a basePath `/items`. The application supports GET, LIST and CREATE items.

```bash

curl localhost:8080/items/999999

{
  "inventory": {
    "salePrice": {
      "currency": "USD",
      "value": "225.72"
    },
    "sellOnGoogleQuantity": "1",
    "kind": "content#inventory",
    "price": {
      "currency": "USD",
      "value": "238.13"
    }
  },
  "productId": "999999",
  "storeCode": "online"
}
```

