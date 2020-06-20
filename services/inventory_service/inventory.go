package inventory_service

import (
	"apigateway/domain/inventory_model"
	"apigateway/errors"
	"apigateway/exrernalRepository/inventory_repo"
)

type InventoryServiceInterface interface {
	GetProductDetails(int64) (*inventory_model.ProductDetails, *errors.RestErr)
}

type inventoryServiceRepo struct {
	inventoryRepo inventory_repo.Inventory
}

func NewInventoryService(inventoryRepository inventory_repo.Inventory) InventoryServiceInterface {
	return &inventoryServiceRepo{
		inventoryRepo: inventoryRepository,
	}
}

func (s *inventoryServiceRepo) GetProductDetails(productId int64) (*inventory_model.ProductDetails, *errors.RestErr) {
	return s.GetProductDetails(productId)
}
