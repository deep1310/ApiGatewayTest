package application

import (
	"apigateway/controller"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Start() {
	router.GET("/ProductDetails/:product_id", controller.GetProductDetails)
	router.POST("/Payment/MakePayment", controller.MakePayment)
	router.POST("/Order/CreateOrder", controller.CreateOrder)
	router.POST("/Order/CompleteOrder/:order_id", controller.CompleteOrder)

	router.Run(":5553")
}

/*
	Here this would have middle ware to get the user id from the
	auth platform and request would be signed by that for order_repo transactions

*/
