package inventory_model

type ProductDetails struct {
	ProductId         int64   `json:"productId"`
	Quantity          int     `json:"quantity"`
	AvailableQuantity int     `json:"availableQuantity"`
	Fare              float64 `json:"fare"`
	DiscountAmt       float64 `json:"discountAmt"`
}

type ProductQtyUpdate struct {
	ProductId  int64 `json:"productId"`
	Quantity   int   `json:"quantity"`
	IsRollback bool  `json:"isRollback"`
}
