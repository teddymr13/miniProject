package service

import (
	"invertory/domain"
	"invertory/errs"
	"reflect"
	"testing"
)

func TestDefaultInventoryService_GetInventoryAll(t *testing.T) {
	type args struct {
		pagination domain.Pagination
	}
	tests := []struct {
		name  string
		s     DefaultInventoryService
		args  args
		want  domain.Pagination
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetInventoryAll(tt.args.pagination)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultInventoryService.GetInventoryAll() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DefaultInventoryService.GetInventoryAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDefaultInventoryService_GetInventoryByID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name  string
		s     DefaultInventoryService
		args  args
		want  *domain.Inventory
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetInventoryByID(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultInventoryService.GetInventoryByID() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DefaultInventoryService.GetInventoryByID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDefaultInventoryService_DeleteInventory(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name  string
		s     DefaultInventoryService
		args  args
		want  *domain.Inventory
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.DeleteInventory(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultInventoryService.DeleteInventory() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DefaultInventoryService.DeleteInventory() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDefaultInventoryService_CreateInventory(t *testing.T) {
	type args struct {
		input domain.CreateInventoryInput
	}
	tests := []struct {
		name  string
		s     DefaultInventoryService
		args  args
		want  *domain.Inventory
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateInventory(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultInventoryService.CreateInventory() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DefaultInventoryService.CreateInventory() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDefaultInventoryService_UpdateInventory(t *testing.T) {
	type args struct {
		id    int
		input domain.CreateInventoryInput
	}
	tests := []struct {
		name  string
		s     DefaultInventoryService
		args  args
		want  *domain.Inventory
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateInventory(tt.args.id, tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultInventoryService.UpdateInventory() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DefaultInventoryService.UpdateInventory() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewInventoryService(t *testing.T) {
	type args struct {
		repository domain.InventoryRepository
	}
	tests := []struct {
		name string
		args args
		want DefaultInventoryService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInventoryService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInventoryService() = %v, want %v", got, tt.want)
			}
		})
	}
}
