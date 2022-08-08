package service

import (
	"fmt"
	"invertory/domain"
	"invertory/errs"
)

type InventoryService interface {
	GetInventoryAll(domain.Pagination) (domain.Pagination, *errs.AppErr)
	GetInventoryByID(int) (*domain.Inventory, *errs.AppErr)
	DeleteInventory(int) (*domain.Inventory, *errs.AppErr)
	CreateInventory(domain.CreateInventoryInput) (*domain.Inventory, *errs.AppErr)
	UpdateInventory(int, domain.CreateInventoryInput) (*domain.Inventory, *errs.AppErr)
}

type DefaultInventoryService struct {
	repo domain.InventoryRepository
}

func (s DefaultInventoryService) GetInventoryAll(pagination domain.Pagination) (domain.Pagination, *errs.AppErr) {
	inventorys, err := s.repo.FindInventoryAll(pagination)
	if err != nil {
		return inventorys, err
	}
	return inventorys, nil
}

func (s DefaultInventoryService) GetInventoryByID(id int) (*domain.Inventory, *errs.AppErr) {

	inventorys, err := s.repo.FindInventoryByID(id)
	if err != nil {
		return nil, err
	}
	return inventorys, nil
}

func (s DefaultInventoryService) DeleteInventory(id int) (*domain.Inventory, *errs.AppErr) {

	inventorys, err := s.repo.DeleteInventories(id)
	if err != nil {
		return nil, err
	}
	return inventorys, nil
}

func (s DefaultInventoryService) CreateInventory(input domain.CreateInventoryInput) (*domain.Inventory, *errs.AppErr) {
	fmt.Println("input", input)
	inventory := domain.Inventory{}
	inventory.Name = input.Name
	inventory.CategoryItem = input.CategoryItem
	inventory.Stock = input.Stock
	inventory.Price = input.Price
	inventory.Status = input.Status

	inventorys, err := s.repo.CreateInventories(inventory)
	if err != nil {
		return nil, err
	}
	return inventorys, nil
}

func (s DefaultInventoryService) UpdateInventory(id int, input domain.CreateInventoryInput) (*domain.Inventory, *errs.AppErr) {
	inventory := domain.Inventory{}
	inventory.Name = input.Name
	inventory.CategoryItem = input.CategoryItem
	inventory.Stock = input.Stock
	inventory.Price = input.Price
	inventory.Status = input.Status

	inventorys, err := s.repo.UpdateInventories(id, inventory)
	if err != nil {
		return nil, err
	}
	return inventorys, nil
}

func NewInventoryService(repository domain.InventoryRepository) DefaultInventoryService {
	return DefaultInventoryService{repository}
}
