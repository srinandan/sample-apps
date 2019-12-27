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
