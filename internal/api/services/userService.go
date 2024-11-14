package userServices

import (
	"context"
	"errors"
	"go-project/internal/api/model"
	"go-project/internal/api/repository"
	"go-project/util"
)

type UserService struct {
	repo userRepository.UserRepository
}

func NewUserService(repo userRepository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignupUser(ctx context.Context, user model.User) (model.User, error) {
	if _, exist := s.repo.GetUser(ctx, user.Email); exist {
		return model.User{}, errors.New("user already exist ")
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return model.User{}, errors.New("hashing password failed")
	}

	user.Password = hashedPassword
	success := s.repo.CreateNewUser(ctx, user)
	if !success {
		return model.User{}, errors.New("error in adding user")
	}

	user.Password = ""
	return user, nil
}

func (s *UserService) LoginUser(ctx context.Context, email string, password string) (model.User, error) {
	user, exist := s.repo.GetUser(ctx, email)
	if !exist {
		return model.User{}, errors.New("user does not exist")
	}

	passwordMatch := util.ComparePassword(password, user.Password)
	if !passwordMatch {
		return model.User{}, errors.New("password does not match")
	}

	user.Password = ""
	return user, nil
}
