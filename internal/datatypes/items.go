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

type Price struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}

type Inventory struct {
	SalePrice            Price  `json:"salePrice,omitempty"`
	SellOnGoogleQuantity string `json:"sellOnGoogleQuantity,omitempty"`
	Availability         int    `json:"availability,omitempty"`
	Kind                 string `json:"kind,omitempty"`
	Price                Price  `json:"price,omitempty"`
}

type Item struct {
	Inventory Inventory `json:"inventory,omitempty"`
	ProductId string    `json:"productId,omitempty"`
	StoreCode string    `json:"storeCode,omitempty"`
}
