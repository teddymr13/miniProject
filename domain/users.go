package domain

import (
	"invertory/errs"
	"time"
)

type Users struct {
	ID        int       `json:"member_id" gorm:"column:user_id"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"password" gorm:"column:password"`
	Role      string    `json:"role" gorm:"column:role"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
}

type UsersRepository interface {
	RegisterUsersInput(Users) (Users, *errs.AppErr)
	LoginUsersInput(string) (Users, *errs.AppErr)
	FindUsersByID(id int) (*Users, *errs.AppErr)
}

type UsersInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
