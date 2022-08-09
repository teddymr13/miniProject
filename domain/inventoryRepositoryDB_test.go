package domain

import (
	"errors"
	"fmt"
	"invertory/errs"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMock() (*gorm.DB, sqlmock.Sqlmock) {
	_, mock, err := sqlmock.New()
	// defer db.Close()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// sanityCheck()
	// db := getDBClient()

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "postgres", "admin", "localhost", "5432", "miniproject")
	gormDB, _ := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	return gormDB, mock
}

func TestInventoryRepositoryDB_FindInventoryAll(t *testing.T) {
	type args struct {
		pagination Pagination
	}
	tests := []struct {
		name    string
		args    args
		want    []Inventory
		wantErr *errs.AppErr
	}{
		// TODO: Add test cases.
		{
			"succcess get data all",
			args{},
			nil,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := NewMock()
			repo := NewInventoryRepositoryDB(db)
			mock.ExpectQuery(`select \* from inventories`).WillReturnError(errors.New(""))

			_, got1 := repo.FindInventoryAll(tt.args.pagination)
			if !reflect.DeepEqual(got1, tt.wantErr) {
				t.Errorf("InventoryRepositoryDB.FindInventoryAll() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}
