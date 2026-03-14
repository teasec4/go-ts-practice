package repository

import (
	"justforfun/internal/model"
)

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	FindUserByID(id string) (*model.User, error)
	CreateUser(user model.User) (error)
}
