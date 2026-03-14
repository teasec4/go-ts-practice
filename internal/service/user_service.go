package service

import (
	"justforfun/internal/model"
	"justforfun/internal/repository"

	"github.com/google/uuid"
)

type UserService struct{
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService{
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers()([]model.User, error){
	return s.Repo.GetAllUsers()
}

func (s *UserService) FindUserByID(id string) (*model.User, error){
	return s.Repo.FindUserByID(id)
}

func (s *UserService) CreateUser(name string)(*model.User, error){
	user := model.User{
		Id: uuid.New(),
		Name: name,
	}
	if err := s.Repo.CreateUser(user); err != nil{
		return nil, err
	}
	
	return &user, nil
}

