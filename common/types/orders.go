package types

type LineItem struct {
	LineItemId string `json:"lineItemId,omitempty"`
	Quantity   int    `json:"quantity,omitempty"`
}

type Order struct {
	OperationId string     `json:"operationId,omitempty"`
	ShipmentId  string     `json:"shipmentId,omitempty"`
	LineItems   []LineItem `json:"lineItems,omitempty"`
	Carrier     string     `json:"carrier,omitempty"`
	TrackingID  string     `json:"trackingId,omitempty"`
}
