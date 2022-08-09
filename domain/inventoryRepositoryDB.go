package domain

import (
	"invertory/errs"
	"invertory/logger"
	"math"

	"gorm.io/gorm"
)

type InventoryRepositoryDB struct {
	db *gorm.DB
}

func NewInventoryRepositoryDB(client *gorm.DB) InventoryRepositoryDB {

	return InventoryRepositoryDB{client}
}

func (s *InventoryRepositoryDB) FindInventoryAll(pagination Pagination) (Pagination, *errs.AppErr) {
	var p Pagination
	// var err error
	var inventorys []Inventory

	tr, totalPages := 0, 0 // pagination attribute
	offset := pagination.Page * pagination.Limit

	errFind := s.db.Limit(pagination.Limit).Offset(offset).Find(&inventorys).Error
	if errFind != nil {
		return p, errs.NewUnexpectedError("pagination error")
	}
	pagination.Rows = inventorys
	// count all data
	totalRows := int64(tr)
	errCount := s.db.Model(inventorys).Count(&totalRows).Error
	if errCount != nil {
		return p, errs.NewUnexpectedError("pagination error")
	}

	pagination.TotalRows = totalRows
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1
	pagination.TotalPages = totalPages
	return pagination, nil
}

func (s *InventoryRepositoryDB) FindInventoryByID(id int) (*Inventory, *errs.AppErr) {
	var err error
	var inventory Inventory
	err = s.db.Where("item_id = ?", id).Find(&inventory).Error
	if err != nil {
		logger.Error("error fetch data to inventory table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	} else {
		return &inventory, nil
	}
}

func (s *InventoryRepositoryDB) DeleteInventories(id int) (*Inventory, *errs.AppErr) {
	var err error
	var inventory Inventory
	err = s.db.Where("item_id = ?", id).Delete(&inventory).Error
	if err != nil {
		logger.Error("error fetch data to inventory table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	} else {
		return &inventory, nil
	}
}

func (s *InventoryRepositoryDB) CreateInventories(i Inventory) (*Inventory, *errs.AppErr) {
	err := s.db.Create(&i).Error
	if err != nil {
		logger.Error("error fetch data to inventory table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	} else {
		return &i, nil
	}
}

func (s *InventoryRepositoryDB) UpdateInventories(id int, i Inventory) (*Inventory, *errs.AppErr) {
	err := s.db.Model(&i).Where("item_id = ?", id).Updates(i).Error
	if err != nil {
		logger.Error("error fetch data to inventory table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	} else {
		return &i, nil
	}
}
