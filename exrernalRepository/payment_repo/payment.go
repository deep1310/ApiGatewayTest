package payment_repo

import (
	"apigateway/domain/payment_model"
	"apigateway/errors"
	"apigateway/exrernalRepository"
	"fmt"
	"net/http"
)

var (
	paymentClient = exrernalRepository.MakeDefaultRestyClient("http://localhost:5552/", 30)
)

type Payment interface {
	CreatePaymentOrder(req *payment_model.PayItemReq) (*payment_model.CreatePaymentResp, *errors.RestErr)
}

type paymentRepo struct{}

func New() Payment {
	return &paymentRepo{}
}

func (s *paymentRepo) CreatePaymentOrder(payReq *payment_model.PayItemReq) (*payment_model.CreatePaymentResp, *errors.RestErr) {

	var paymentResp payment_model.CreatePaymentResp
	var restErr *errors.RestErr
	req := paymentClient.R()
	paymentUrl := fmt.Sprintf("/CreatePayment")
	resp, err := req.SetBody(&payReq).SetResult(&paymentResp).SetError(&restErr).Post(paymentUrl)

	if err != nil {
		return nil, errors.InternalServerError("unable to create payment")
	} else if restErr != nil {
		return nil, restErr
	} else if resp.StatusCode() != http.StatusOK {
		return nil, errors.InternalServerError("unable to create payment")
	}

	return &paymentResp, nil
}
