package domain

import (
	"fmt"
	"invertory/errs"
	"invertory/logger"

	"gorm.io/gorm"
)

type UsersRepositoryDB struct {
	db *gorm.DB
}

func NewUsersRepositoryDB(client *gorm.DB) UsersRepositoryDB {
	return UsersRepositoryDB{client}
}

func (u UsersRepositoryDB) RegisterUsersInput(users Users) (Users, *errs.AppErr) {
	err := u.db.Create(&users)

	if err != nil {
		logger.Error("error to register user")
		return users, errs.NewUnexpectedError("Unexpected error")
	}
	return users, nil
}

func (u UsersRepositoryDB) LoginUsersInput(email string) (Users, *errs.AppErr) {

	var users Users
	err := u.db.Where("email = ?", email).Find(&users).Error
	if err != nil {
		logger.Error("error to login user")
		return users, errs.NewUnexpectedError("unexpected error")
	}
	fmt.Println("users", users)
	return users, nil
}

func (s *UsersRepositoryDB) FindUsersByID(id int) (*Users, *errs.AppErr) {
	var err error
	var users Users
	err = s.db.Where("user_id = ?", id).Find(&users).Error
	if err != nil {
		logger.Error("error fetch data to inventory table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	} else {
		return &users, nil
	}
}
