package domain

import (
	"invertory/errs"
)

type Inventory struct {
	ID           int    `json:"itemID" gorm:"column:item_id"`
	Name         string `json:"name" gorm:"column:name"`
	CategoryItem string `json:"categoryItem" gorm:"column:category_item"`
	Stock        int    `json:"stock" gorm:"column:stock"`
	Price        int    `json:"price" gorm:"column:price"`
	Status       string `json:"status" gorm:"column:status"`
}

type CreateInventoryInput struct {
	Name         string `json:"name" binding:"required"`
	CategoryItem string `json:"categoryItem" binding:"required"`
	Stock        int    `json:"stock" binding:"required"`
	Price        int    `json:"price" binding:"required"`
	Status       string `json:"status" binding:"required"`
}

type InventoryRepository interface {
	FindInventoryAll(Pagination) (Pagination, *errs.AppErr)
	FindInventoryByID(int) (*Inventory, *errs.AppErr)
	DeleteInventories(int) (*Inventory, *errs.AppErr)
	CreateInventories(Inventory) (*Inventory, *errs.AppErr)
	UpdateInventories(int, Inventory) (*Inventory, *errs.AppErr)
}

type Pagination struct {
	Limit      int         `json:"limit"`
	Page       int         `json:"page"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

// func (c Inventory) convertStatusName() string {
// 	statusName := "new"
// 	if c.Status == "0" {
// 		statusName = "used"
// 	}
// 	return statusName
// }

// func (c Inventory) ToDTO() dto.InventoryResponse {
// 	return dto.InventoryResponse{
// 		ID:           c.ID,
// 		Name:         c.Name,
// 		CategoryItem: c.CategoryItem,
// 		Price:        c.Price,
// 		Status:       c.convertStatusName(),
// 	}
// }
