package inventory_repo

import (
	"apigateway/domain/inventory_model"
	"apigateway/errors"
	"apigateway/exrernalRepository"
	"fmt"
	"net/http"
)

var (
	inventoryClient = exrernalRepository.MakeDefaultRestyClient("http://localhost:5550/", 30)
)

type Inventory interface {
	GetProductDetails(int64) (*inventory_model.ProductDetails, *errors.RestErr)
	UpdateProductQuantity(*inventory_model.ProductQtyUpdate) *errors.RestErr
}

type inventoryRepo struct{}

func New() Inventory {
	return &inventoryRepo{}
}

func (s *inventoryRepo) GetProductDetails(productId int64) (*inventory_model.ProductDetails, *errors.RestErr) {

	var productDetails inventory_model.ProductDetails
	var restErr *errors.RestErr
	req := inventoryClient.R()
	inventoryUrl := fmt.Sprintf("/ProductDetails/%d", productId)
	resp, err := req.SetResult(&productDetails).SetError(&restErr).Get(inventoryUrl)

	if err != nil {
		return nil, errors.InternalServerError("unable to get product details")
	} else if restErr != nil {
		return nil, restErr
	} else if resp.StatusCode() != http.StatusOK {
		return nil, errors.InternalServerError("unable to get product details")
	}

	return &productDetails, nil
}

func (s *inventoryRepo) UpdateProductQuantity(updateReq *inventory_model.ProductQtyUpdate) *errors.RestErr {

	var restErr *errors.RestErr
	req := inventoryClient.R()
	inventoryUrl := fmt.Sprintf("/ProductDetails/UpdateQuantity/")
	resp, err := req.SetBody(&updateReq).SetError(&restErr).Post(inventoryUrl)

	if err != nil {
		return errors.InternalServerError("unable to update product quantity")
	} else if restErr != nil {
		return restErr
	} else if resp.StatusCode() != http.StatusOK {
		return errors.InternalServerError("unable to update product quantity")
	}

	return nil
}
