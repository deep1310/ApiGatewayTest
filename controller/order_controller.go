package controller

import (
	"apigateway/domain/order_model"
	"apigateway/errors"
	"apigateway/exrernalRepository/inventory_repo"
	"apigateway/exrernalRepository/order_repo"
	"apigateway/services/order_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func CreateOrder(c *gin.Context) {

	var orderReq order_model.CreateOrderReq
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		apiReqErr := errors.BadRequestError("invalid request")
		c.JSON(apiReqErr.Code, apiReqErr)
		return
	}

	orderClient := order_service.NewOrderService(inventory_repo.New(), order_repo.OrderService(), order_repo.OrderPayService())
	result, err := orderClient.CreateOrder(&orderReq)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func CompleteOrder(c *gin.Context) {
	orderId := c.Param("order_id")
	orderId = strings.TrimSpace(orderId)
	if orderId == "" {
		apiReqErr := errors.BadRequestError("order id is empty")
		c.JSON(apiReqErr.Code, apiReqErr)
		return
	}

	orderIdInt, userErr := strconv.ParseInt(orderId, 10, 64)
	if userErr != nil {
		apiReqErr := errors.BadRequestError("order id is not int")
		c.JSON(apiReqErr.Code, apiReqErr)
		return
	}

	orderClient := order_service.NewOrderService(inventory_repo.New(), order_repo.OrderService(), order_repo.OrderPayService())
	orderResp, err := orderClient.CompleteOrder(orderIdInt)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, orderResp)
}
