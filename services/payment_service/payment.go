package payment_service

import (
	"apigateway/domain/payment_model"
	"apigateway/errors"
	"apigateway/exrernalRepository/order_repo"
	"apigateway/exrernalRepository/payment_repo"
)

type PaymentServiceInterface interface {
	CreateOnlinePaymentOrder(item *payment_model.PayItemReq) (*payment_model.CreatePaymentResp, *errors.RestErr)
}

type paymentServiceRepo struct {
	orderPayRepo order_repo.OrderPayInterface
	orderRepo    order_repo.OrderInterface
	paymentRepo  payment_repo.Payment
}

func NewPaymentService(orderPayRepository order_repo.OrderPayInterface, orderRepository order_repo.OrderInterface, paymentRepository payment_repo.Payment) PaymentServiceInterface {
	return &paymentServiceRepo{
		orderPayRepo: orderPayRepository,
		orderRepo:    orderRepository,
		paymentRepo:  paymentRepository,
	}
}

func (s *paymentServiceRepo) CreateOnlinePaymentOrder(req *payment_model.PayItemReq) (*payment_model.CreatePaymentResp, *errors.RestErr) {

	orderResp, orderErr := s.orderRepo.GetOrder(req.OrderId)
	if orderErr != nil {
		return nil, orderErr
	}
	req.Amount = orderResp.PaymentSessionData.AmountLeft

	paymentItemReq := &payment_model.PaymentItem{
		PaymentSessionId: orderResp.PaymentSessionData.PaymentSessionId,
		Amount:           req.Amount,
		PaymentType:      "ONLINE",
		PaymentMode:      req.PaymentType,
		Status:           "INITIATED",
	}
	if err := s.orderPayRepo.AddPaymentItem(paymentItemReq); err != nil {
		return nil, err
	}
	return s.paymentRepo.CreatePaymentOrder(req)
}
