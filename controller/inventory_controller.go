package controller

import (
	"apigateway/errors"
	"apigateway/exrernalRepository/inventory_repo"
	"apigateway/services/inventory_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func GetProductDetails(c *gin.Context) {

	productId := c.Param("product_id")
	productId = strings.TrimSpace(productId)
	if productId == "" {
		apiReqErr := errors.BadRequestError("product id is empty")
		c.JSON(apiReqErr.Code, apiReqErr)
		return
	}

	productIdInt, userErr := strconv.ParseInt(productId, 10, 64)
	if userErr != nil {
		apiReqErr := errors.BadRequestError("product id is not int")
		c.JSON(apiReqErr.Code, apiReqErr)
		return
	}

	result, err := inventory_service.NewInventoryService(inventory_repo.New()).GetProductDetails(productIdInt)
	if err != nil {
		c.JSON(err.Code, result)
		return
	}
	c.JSON(http.StatusOK, result)
}
