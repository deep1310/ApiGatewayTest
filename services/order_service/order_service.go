package order_service

import (
	"apigateway/domain/inventory_model"
	"apigateway/domain/order_model"
	"apigateway/domain/payment_model"
	"apigateway/errors"
	"apigateway/exrernalRepository/inventory_repo"
	"apigateway/exrernalRepository/order_repo"
)

type OrderServiceInterface interface {
	CreateOrder(*order_model.CreateOrderReq) (*order_model.CreateOrderResp, *errors.RestErr)
	CompleteOrder(int64) (*order_model.OrderResp, *errors.RestErr)
}

type orderServiceRepo struct {
	inventoryRepo inventory_repo.Inventory
	orderRepo     order_repo.OrderInterface
	orderPayRepo  order_repo.OrderPayInterface
}

func NewOrderService(inventoryRepository inventory_repo.Inventory, orderRepository order_repo.OrderInterface, orderPayRepository order_repo.OrderPayInterface) OrderServiceInterface {
	return &orderServiceRepo{
		inventoryRepo: inventoryRepository,
		orderRepo:     orderRepository,
		orderPayRepo:  orderPayRepository,
	}
}

func (o *orderServiceRepo) CreateOrder(orderRequest *order_model.CreateOrderReq) (*order_model.CreateOrderResp, *errors.RestErr) {

	if err := orderRequest.ValidateOrderRequest(); err != nil {
		return nil, err
	}

	prodDetails, invErr := o.inventoryRepo.GetProductDetails(orderRequest.ProductId)
	if invErr != nil {
		return nil, invErr
	}

	if prodDetails.AvailableQuantity < orderRequest.Quantity {
		return nil, errors.InternalServerError("Not enough quantity available")
	}

	orderRequest.Fare = prodDetails.Fare * float64(orderRequest.Quantity)
	orderRequest.DiscountAmt = prodDetails.DiscountAmt * float64(orderRequest.Quantity)
	return o.orderRepo.CreateOrder(orderRequest)
}

func (o *orderServiceRepo) CompleteOrder(orderId int64) (*order_model.OrderResp, *errors.RestErr) {

	orderResp, orderErr := o.orderRepo.GetOrder(orderId)
	if orderErr != nil {
		/*
			call payment refund method
		*/
		return nil, orderErr
	}

	updateQuantityReq := &inventory_model.ProductQtyUpdate{
		ProductId: orderResp.OrderItem.ProductId,
		Quantity:  orderResp.OrderItem.Quantity,
	}
	if err := o.inventoryRepo.UpdateProductQuantity(updateQuantityReq); err != nil {
		return nil, err
	}
	updateOrderResp, err := o.completeOrderWorkflow(orderResp)
	if err != nil {
		o.rollbackSeatsAndInitiateRefund(orderResp)
	}
	return updateOrderResp, nil
}

func (o *orderServiceRepo) completeOrderWorkflow(orderResp *order_model.OrderResp) (*order_model.OrderResp, *errors.RestErr) {

	var onlinePaymentId int64
	for _, paymentItem := range orderResp.PaymentItemData {
		if paymentItem.PaymentType == "ONLINE" {
			onlinePaymentId = paymentItem.PaymentItemId
		}
	}
	payUpdateReq := &payment_model.PaymentItemUpdateReq{
		PaymentSessionId: orderResp.PaymentSessionData.PaymentSessionId,
		Status:           "COMPLETED",
		PaymentItemId:    onlinePaymentId,
	}
	if err := o.orderPayRepo.UpdatePaymentItem(payUpdateReq); err != nil {
		return nil, err
	}

	updateOrderResp, orderErr := o.orderRepo.OrderComplete(orderResp.OrderItem.OrderId)
	if orderErr != nil {
		return nil, orderErr
	}
	return updateOrderResp, nil
}

func (o *orderServiceRepo) rollbackSeatsAndInitiateRefund(orderResp *order_model.OrderResp) {

	/*
		call payment refund method. Not writing code for this as fo now
	*/

	/*
		call quantity rollback method
	*/

	updateQuantityReq := &inventory_model.ProductQtyUpdate{
		ProductId:  orderResp.OrderItem.ProductId,
		Quantity:   orderResp.OrderItem.Quantity,
		IsRollback: true,
	}
	o.inventoryRepo.UpdateProductQuantity(updateQuantityReq)
}
