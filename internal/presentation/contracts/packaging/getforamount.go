package contracts

type GetForAmountResponse struct {
	Packs        []*Packs `json:"packs"`
	PackQuantity int      `json:"packQuantity"`
	TotalAmount  int      `json:"totalAmount"`
	LeftAmount   int      `json:"leftAmount"`
}

type Packs struct {
	Size     int `json:"size"`
	Quantity int `json:"quantity"`
}
