package controller

import (
	"apigateway/domain/payment_model"
	"apigateway/errors"
	"apigateway/exrernalRepository/order_repo"
	"apigateway/exrernalRepository/payment_repo"
	"apigateway/services/payment_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MakePayment(c *gin.Context) {

	var payReq payment_model.PayItemReq

	if err := c.ShouldBindJSON(&payReq); err != nil {
		apiReqErr := errors.BadRequestError("invalid request")
		c.JSON(apiReqErr.Code, apiReqErr)
		return
	}

	payClient := payment_service.NewPaymentService(order_repo.OrderPayService(),order_repo.OrderService(), payment_repo.New())
	result, err := payClient.CreatePaymentOrder(&PayItemReq)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
