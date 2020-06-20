package order_repo

import (
	"apigateway/domain/payment_model"
	"apigateway/errors"
	"apigateway/exrernalRepository"
	"fmt"
	"net/http"
)

var (
	payOrderClient = exrernalRepository.MakeDefaultRestyClient("http://localhost:5552/", 30)
)

type OrderPayInterface interface {
	AddPaymentItem(*payment_model.PaymentItem) *errors.RestErr
	UpdatePaymentItem(*payment_model.PaymentItemUpdateReq) *errors.RestErr
}

type orderPayRepo struct{}

func OrderPayService() OrderPayInterface {
	return &orderPayRepo{}
}

func (s *orderPayRepo) AddPaymentItem(itemReq *payment_model.PaymentItem) *errors.RestErr {

	var restErr *errors.RestErr
	req := orderClient.R()
	payCreateUrl := fmt.Sprintf("/Payment/AddPaymentItem")
	resp, err := req.SetBody(&itemReq).SetError(&restErr).Post(payCreateUrl)
	if err != nil {
		return errors.InternalServerError("unable to add pay item")
	} else if restErr != nil {
		return restErr
	} else if resp.StatusCode() != http.StatusCreated {
		return errors.InternalServerError("unable to create order")
	}

	return nil
}

func (s *orderPayRepo) UpdatePaymentItem(payUpReq *payment_model.PaymentItemUpdateReq) *errors.RestErr {

	var restErr *errors.RestErr
	req := orderClient.R()
	payUpdateUrl := fmt.Sprintf("/Payment/UpdatePaymentItem")
	resp, err := req.SetBody(&payUpReq).SetError(&restErr).Post(payUpdateUrl)

	if err != nil {
		return errors.InternalServerError("unable to update pay order")
	} else if restErr != nil {
		return restErr
	} else if resp.StatusCode() != http.StatusOK {
		return errors.InternalServerError("unable to create order")
	}

	return nil
}
