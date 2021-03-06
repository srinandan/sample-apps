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

// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type Inventory struct {
	SalePrice            *Price  `json:"salePrice"`
	SellOnGoogleQuantity *string `json:"sellOnGoogleQuantity"`
	Availability         *int    `json:"availability"`
	Kind                 *string `json:"kind"`
	Price                *Price  `json:"price"`
}

type Item struct {
	Inventory *Inventory `json:"inventory"`
	ProductID *string    `json:"productId"`
	StoreCode *string    `json:"storeCode"`
}

type LineItem struct {
	Item     *Item `json:"item"`
	Quantity *int  `json:"quantity"`
}

type Order struct {
	OperationID string      `json:"operationId"`
	ShipmentID  *string     `json:"shipmentId"`
	LineItems   []*LineItem `json:"lineItems"`
	Carrier     *string     `json:"carrier"`
	TrackingID  *string     `json:"trackingId"`
}

type Price struct {
	Currency *string `json:"currency"`
	Value    *string `json:"value"`
}
