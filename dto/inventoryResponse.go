package dto

//format tampilan untuk rest api
type InventoryResponse struct {
	ID           int    `json:"item_id" gorm:"column:item_id"`
	Name         string `json:"name" gorm:"column:name"`
	CategoryItem string `json:"category_item" gorm:"column:category_item"`
	Stock        int    `json:"stock" gorm:"column:stock"`
	Price        int    `json:"price" gorm:"column:price"`
	Status       string `json:"status" gorm:"column:status"`
}
