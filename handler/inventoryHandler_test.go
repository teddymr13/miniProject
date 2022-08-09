package handler

import (
	"invertory/service"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewInventoryHandler(t *testing.T) {
	type args struct {
		service service.InventoryService
	}
	tests := []struct {
		name string
		args args
		want *InventoryHandlers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInventoryHandler(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInventoryHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInventoryHandlers_GetInventoryAll(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ch   *InventoryHandlers
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.GetInventoryAll(tt.args.c)
		})
	}
}

func TestInventoryHandlers_GetInventoryByID(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ch   *InventoryHandlers
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.GetInventoryByID(tt.args.c)
		})
	}
}

func TestInventoryHandlers_DeleteInventory(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ch   *InventoryHandlers
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.DeleteInventory(tt.args.c)
		})
	}
}

func TestInventoryHandlers_CreateInventory(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ch   *InventoryHandlers
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.CreateInventory(tt.args.c)
		})
	}
}

func TestInventoryHandlers_UpdateInventory(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ch   *InventoryHandlers
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.UpdateInventory(tt.args.c)
		})
	}
}
