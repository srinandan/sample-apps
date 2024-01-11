// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
