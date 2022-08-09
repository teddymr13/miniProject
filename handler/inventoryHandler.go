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

// func getCurrentMemberJWT(c *gin.Context) int {
// 	currMember := c.MustGet("currentMember").(members.Member)
// 	fmt.Println("currMember", currMember)
// 	memberId := currMember.ID
// 	return memberId
// }

func (ch *InventoryHandlers) GetInventoryAll(c *gin.Context) {
	pagination := dto.GeneratePaginationRequest(c)
	inventorys, err := ch.service.GetInventoryAll(*pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	} else {
		response := errs.SuccessRequestUser("Success get all data inventory", inventorys)
		c.JSON(http.StatusOK, response)
	}

	c.JSON(http.StatusOK, inventorys)
}

func (ch *InventoryHandlers) GetInventoryByID(c *gin.Context) {
	itemid := c.Param("itemid")
	itemId, _ := strconv.Atoi(itemid)
	inventorys, err := ch.service.GetInventoryByID(itemId)
	if err != nil {
		errsMessage := fmt.Sprintf("Failed get data inventory by item id %v error!", itemid)
		res := errs.NewBadRequestError(errsMessage)
		c.JSON(http.StatusBadRequest, res)
		return
	} else {
		response := errs.SuccessRequestUser("Success get data inventory by item id", itemid)
		c.JSON(http.StatusOK, response)
	}
	c.JSON(http.StatusOK, inventorys)
}

func (ch *InventoryHandlers) DeleteInventory(c *gin.Context) {
	itemid := c.Param("itemid")
	itemId, _ := strconv.Atoi(itemid)
	_, err := ch.service.DeleteInventory(itemId)
	if err != nil {
		errsMessage := fmt.Sprintf("Failed delete data inventory %v error!", itemid)
		res := errs.NewBadRequestError(errsMessage)
		c.JSON(http.StatusBadRequest, res)
		return
	} else {
		successMsg := fmt.Sprintf("Success delete data inventory by item id %v", itemId)
		response := errs.SuccessRequestUser(successMsg, nil)
		c.JSON(http.StatusOK, response)
	}

}

func (ch *InventoryHandlers) CreateInventory(c *gin.Context) {
	var input domain.CreateInventoryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := errs.NewBadRequestError("Failed create data inventory")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	inventorys, errSer := ch.service.CreateInventory(input)
	if errSer != nil {
		response := errs.NewBadRequestError("Failed create data inventory")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	errsMessage := "Success create data inventory"
	response := errs.SuccessRequestUser(errsMessage, inventorys)
	c.JSON(http.StatusCreated, response)
}

func (ch *InventoryHandlers) UpdateInventory(c *gin.Context) {

	itemid := c.Param("itemid")
	itemId, _ := strconv.Atoi(itemid)
	var input domain.CreateInventoryInput
	err := c.ShouldBindJSON(&input)
	inventorys, _ := ch.service.UpdateInventory(itemId, input)
	if err != nil {
		errMessage := fmt.Sprintf("Failed update data inventory %v error!", itemid)
		res := errs.NewBadRequestError(errMessage)
		c.JSON(http.StatusBadRequest, res)
		return
	} else {
		inventorys.ID = itemId
		successMsg := fmt.Sprintf("Success update data inventory by item id %v", itemId)
		response := errs.SuccessRequestUser(successMsg, inventorys)
		c.JSON(http.StatusOK, response)
	}
}
