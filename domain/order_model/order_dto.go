package order_model

import (
	"apigateway/errors"
)

type CreateOrderReq struct {
	Quantity    int     `json:"quantity"`
	ProductId   int64   `json:"productId"`
	UserId      int64   `json:"userId"`
	Fare        float64 `json:"fare"`
	DiscountAmt float64 `json:"discountAmt"`
}

type CreateOrderResp struct {
	OrderId   int64   `json:"orderId"`
	UserId    int64   `json:"userId"`
	Quantity  int     `json:"quantity"`
	ProductId int64   `json:"productId"`
	Fare      float64 `json:"fare"`
	Discount  float64 `json:"discount"`
	Status    string  `json:"status"`
}

type OrderResp struct {
	OrderItem          CreateOrderResp `json:"order_item"`
	PaymentSessionData PaymentSession  `json:"paymentSession"`
	PaymentItemData    []PaymentItem   `json:"paymentItems"`
}

type PaymentSession struct {
	PaymentSessionId int64   `json:"PaymentSessionId"`
	OrderId          int64   `json:"orderId"`
	OrderAmount      float64 `json:"orderAmount"`
	AmountPaid       float64 `json:"amountPaid"`
	AmountLeft       float64 `json:"amountLeft"`
	Status           string  `json:"status"`
}

type PaymentItem struct {
	PaymentItemId    int64   `json:"paymentItemId"`
	PaymentSessionId int64   `json:"paymentSessionId"`
	Amount           float64 `json:"amount"`
	PaymentType      string  `json:"paymentType"`
	Status           string  `json:"status"`
}

func (order *CreateOrderReq) ValidateOrderRequest() *errors.RestErr {

	if (order.ProductId) <= 0 {
		if err := errors.BadRequestError("product id is empty"); err != nil {
			return err
		}
	}
	if order.Quantity == 0 {
		if err := errors.BadRequestError("quantity cannot be 0"); err != nil {
			return err
		}
	}

	return nil
}
