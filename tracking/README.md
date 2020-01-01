# sample-apps-tracking

## Usage

This sample contains two applications:

* A tracking microservice is a gRPC server on port 50051
* A tracking client microservice which is a REST/HTTP server (on port 8080) and a gRPC client

```bash

curl localhost:8080/tracking-client/200000

{
    "trackingId": "200000",
    "status": "delivered",
    "created": "2014-11-18T10:51:54Z",
    "updated": "2014-11-19T10:51:54Z",
    "signed": "John Tester",
    "weight": "17.6",
    "est_delivery_date": "2014-11-27T00:00:00Z",
    "carrier": "UPS",
    "trackingLocation": {
        "city": "Beverly Hills",
        "state": "CA",
        "country": "USA",
        "zip": 90210
    }
}
```