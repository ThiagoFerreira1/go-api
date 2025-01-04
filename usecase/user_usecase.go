package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type UserUseCase struct {
	repository *repository.UserRepository
}

func NewUserUseCase(repo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{repository: repo}
}

func (uu *UserUseCase) CreateUse(user model.User) (model.User, error) {
	userId, err := uu.repository.CreateUser(user)

	if err != nil {
		return model.User{}, err
	}

	user.ID = userId
	return user, nil
}
