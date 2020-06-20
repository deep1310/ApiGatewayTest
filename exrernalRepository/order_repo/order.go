package order_repo

import (
	"apigateway/domain/order_model"
	"apigateway/errors"
	"apigateway/exrernalRepository"
	"fmt"
	"net/http"
)

var (
	orderClient = exrernalRepository.MakeDefaultRestyClient("http://localhost:5551/", 30)
)

type OrderInterface interface {
	CreateOrder(*order_model.CreateOrderReq) (*order_model.CreateOrderResp, *errors.RestErr)
	GetOrder(int64) (*order_model.OrderResp, *errors.RestErr)
	OrderComplete(int64) (*order_model.OrderResp, *errors.RestErr)
}

type orderRepo struct{}

func OrderService() OrderInterface {
	return &orderRepo{}
}

func (oService *orderRepo) CreateOrder(orderReq *order_model.CreateOrderReq) (*order_model.CreateOrderResp, *errors.RestErr) {

	var orderDetails order_model.CreateOrderResp
	var restErr *errors.RestErr
	req := orderClient.R()
	orderCreateUrl := fmt.Sprintf("/Order/CreateOrder")
	resp, err := req.SetBody(&orderReq).SetResult(&orderDetails).SetError(&restErr).Post(orderCreateUrl)

	if err != nil {
		return nil, errors.InternalServerError("unable to create order")
	} else if restErr != nil {
		return nil, restErr
	} else if resp.StatusCode() != http.StatusOK {
		return nil, errors.InternalServerError("unable to create order")
	}

	return &orderDetails, nil
}

func (oService *orderRepo) GetOrder(orderId int64) (*order_model.OrderResp, *errors.RestErr) {

	var orderDetails order_model.OrderResp
	var restErr *errors.RestErr
	req := orderClient.R()
	orderCreateUrl := fmt.Sprintf("/Order/GetOrder/:%d", orderId)
	resp, err := req.SetResult(&orderDetails).SetError(&restErr).Get(orderCreateUrl)

	if err != nil {
		return nil, errors.InternalServerError("unable to get order")
	} else if restErr != nil {
		return nil, restErr
	} else if resp.StatusCode() != http.StatusCreated {
		return nil, errors.InternalServerError("unable to get order")
	}

	return &orderDetails, nil
}

func (oService *orderRepo) OrderComplete(orderId int64) (*order_model.OrderResp, *errors.RestErr) {

	var orderDetails order_model.OrderResp
	var restErr *errors.RestErr
	req := orderClient.R()
	orderCreateUrl := fmt.Sprintf("/Order/CompleteOrder/:%d", orderId)
	resp, err := req.SetResult(&orderDetails).SetError(&restErr).Post(orderCreateUrl)

	if err != nil {
		return nil, errors.InternalServerError("unable to complete order")
	} else if restErr != nil {
		return nil, restErr
	} else if resp.StatusCode() != http.StatusOK {
		return nil, errors.InternalServerError("unable to complete order")
	}
	return &orderDetails, nil
}
