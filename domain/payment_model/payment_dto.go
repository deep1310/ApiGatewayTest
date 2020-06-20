package payment_model

type CreatePaymentResp struct {
	RedirectUrl string `json:"redirectUrl"`
}

type PaymentItemUpdateReq struct {
	PaymentItemId    int64  `json:"paymentItemId"`
	PaymentSessionId int64  `json:"paymentSessionId"`
	Status           string `json:"status"`
}

type PaymentItem struct {
	PaymentItemId    int64   `json:"paymentItemId"`
	PaymentSessionId int64   `json:"paymentSessionId"`
	Amount           float64 `json:"amount"`
	PaymentType      string  `json:"paymentType"`
	PaymentMode      string  `json:"paymentMode"`
	Status           string  `json:"status"`
}

type PayItemReq struct {
	Amount      float64 `json:"amount"`
	PaymentType string  `json:"paymentType"`
	OrderId     int64   `json:"orderId"`
}
