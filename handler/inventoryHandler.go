package handler

import (
	"fmt"
	"invertory/domain"
	"invertory/dto"
	"invertory/errs"
	"invertory/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryHandlers struct {
	service service.InventoryService
}

func NewInventoryHandler(service service.InventoryService) *InventoryHandlers {
	return &InventoryHandlers{service: service}

}

func (ch *InventoryHandlers) GetInventoryAll(c *gin.Context) {

	// /customers?status=xxxx
	// status := c.Request.URL.Query().Get("status")
	pagination := dto.GeneratePaginationRequest(c)
	inventorys, err := ch.service.GetInventoryAll(*pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	} else {
		// successMsg := fmt.Println("Success get all data inventory")
		response := errs.SuccessRequest("Success get all data inventory")
		c.JSON(http.StatusOK, response)
	}

	c.JSON(http.StatusOK, inventorys)
}

func (ch *InventoryHandlers) GetInventoryByID(c *gin.Context) {
	// /customers?status=xxxx
	// status := c.Request.URL.Query().Get("status")
	itemid := c.Param("itemid")
	itemId, _ := strconv.Atoi(itemid)
	inventorys, err := ch.service.GetInventoryByID(itemId)
	if err != nil {
		errsMessage := fmt.Sprintf("Failed get data inventory by item id %v invalid!", itemid)
		res := errs.NewBadRequestError(errsMessage)
		c.JSON(http.StatusBadRequest, res)
		return
	} else {
		successMsg := fmt.Sprintf("Success get data inventory by item id %v", itemid)
		response := errs.SuccessRequest(successMsg)
		c.JSON(http.StatusOK, response)
	}
	c.JSON(http.StatusOK, inventorys)
}

func (ch *InventoryHandlers) DeleteInventory(c *gin.Context) {
	// /customers?status=xxxx
	// status := r.URL.Query().Get("status")
	itemid := c.Param("itemid")
	itemId, _ := strconv.Atoi(itemid)
	inventorys, err := ch.service.DeleteInventory(itemId)
	if err != nil {
		errsMessage := fmt.Sprintf("Failed delete data inventory %v invalid!", itemid)
		res := errs.NewBadRequestError(errsMessage)
		c.JSON(http.StatusBadRequest, res)
		return
	} else {
		successMsg := fmt.Sprintf("Success delete data inventory %v", itemid)
		response := errs.SuccessRequest(successMsg)
		c.JSON(http.StatusOK, response)
	}
	c.JSON(http.StatusOK, inventorys)
}

func (ch *InventoryHandlers) CreateInventory(c *gin.Context) {
	var input domain.CreateInventoryInput
	// fmt.Println("tedy", input)
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// errorMessage := gin.H{"errors": err.Error()}
		response := errs.NewBadRequestError("Failed create data inventory")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	inventorys, errSer := ch.service.CreateInventory(input)
	if errSer != nil {
		response := errs.SuccessRequest("Success create data inventory")
		c.JSON(http.StatusOK, response)
		return
	}
	// response := errs.NewValidationError("Your Food has been created")
	c.JSON(http.StatusOK, inventorys)
}

func (ch *InventoryHandlers) UpdateInventory(c *gin.Context) {

	itemid := c.Param("itemid")
	itemId, _ := strconv.Atoi(itemid)
	var input domain.CreateInventoryInput
	err := c.ShouldBindJSON(&input)
	inventorys, _ := ch.service.UpdateInventory(itemId, input)
	if err != nil {
		errsMessage := fmt.Sprintf("Failed update data inventory %v invalid!", itemid)
		res := errs.NewBadRequestError(errsMessage)
		c.JSON(http.StatusBadRequest, res)
		return
	} else {
		successMsg := fmt.Sprintf("Success update data inventory %v", itemid)
		response := errs.SuccessRequest(successMsg)
		c.JSON(http.StatusOK, response)
	}
	c.JSON(http.StatusOK, inventorys)
}

// func writeResponse(w http.ResponseWriter, code int, data interface{}) {
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(code)

// 	if err := json.NewEncoder(w).Encode(data); err != nil {
// 		panic(err)
// 	}
// }
