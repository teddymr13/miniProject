package service

import (
	"fmt"
	"invertory/domain"
	"invertory/errs"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UsersService interface {
	CreateUsers(domain.UsersInput) (domain.Users, *errs.AppErr)
	LoginUsers(domain.Login) (domain.Users, *errs.AppErr)
	GetUsersByID(id int) (*domain.Users, *errs.AppErr)
}

type DefaultUsersService struct {
	repo domain.UsersRepository
}

func NewUsersService(repository domain.UsersRepository) DefaultUsersService {
	return DefaultUsersService{repository}
}

func (u DefaultUsersService) CreateUsers(input domain.UsersInput) (domain.Users, *errs.AppErr) {
	user := domain.Users{}
	user.Email = input.Email
	user.Role = input.Role
	user.CreatedAt = time.Now()

	hashPassword, errBc := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if errBc != nil {
		return user, errs.NewUnexpectedError("unexpected error")
	}
	fmt.Println(hashPassword)

	user.Password = string(hashPassword)

	users, err := u.repo.RegisterUsersInput(user)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (u DefaultUsersService) LoginUsers(input domain.Login) (domain.Users, *errs.AppErr) {
	var users domain.Users
	users.Email = input.Email
	Email := input.Email
	Password := input.Password

	user, err := u.repo.LoginUsersInput(Email)
	if err != nil {
		return user, err
	}

	if Email == "" {
		return user, err
	}
	errByc := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
	if errByc != nil {
		return users, errs.NewBadRequestError("wrong password email")
	}

	return user, nil
}

func (u DefaultUsersService) GetUsersByID(id int) (*domain.Users, *errs.AppErr) {
	users, err := u.repo.FindUsersByID(id)
	if err != nil {
		return users, err
	}
	if users.Email == "" {
		return users, errs.NewBadRequestError("no found on that email")
	}
	return users, nil
}
