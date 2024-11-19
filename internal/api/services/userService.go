package userServices

import (
	"context"
	"errors"
	"fmt"
	"go-project/internal/api/model"
	userRepository "go-project/internal/api/repository"
	"go-project/util"
	jwtauth "go-project/util/auth"
)

type UserService struct {
	repo userRepository.UserRepository
	auth jwtauth.Auth
}

func NewUserService(repo userRepository.UserRepository, auth jwtauth.Auth) *UserService {
	return &UserService{repo: repo, auth: auth}
}

func (s *UserService) SignupUser(ctx context.Context, user model.User) (model.User, string, string, error) {
	if _, exist := s.repo.GetUser(ctx, user.Email); exist {
		return model.User{}, "", "", errors.New("user already exist ")
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return model.User{}, "", "", errors.New("hashing password failed")
	}

	user.Password = hashedPassword
	newUser, err := s.repo.CreateNewUser(ctx, user)
	if err != nil {
		return model.User{}, "", "", errors.New("error in adding user")
	}

	accessToken, _ := s.auth.GenerateAccessToken(user.Id)
	refreshToken, _ := s.auth.GenerateRefreshToken(user.Id)

	newToken := model.Token{
		Token:  refreshToken,
		UserId: newUser.Id,
	}
	_, err = s.repo.AddToken(ctx, newToken)
	if err != nil {
		return model.User{}, "", "", errors.New("error in adding token")
	}
	user.Password = ""
	return user, accessToken, refreshToken, nil
}

func (s *UserService) LoginUser(ctx context.Context, email string, password string) (model.User, string, string, error) {
	user, exist := s.repo.GetUser(ctx, email)
	if !exist {
		return model.User{}, "", "", errors.New("user does not exist")
	}

	passwordMatch := util.ComparePassword(password, user.Password)
	if !passwordMatch {
		return model.User{}, "", "", errors.New("password does not match")
	}

	accessToken, _ := s.auth.GenerateAccessToken(user.Id)
	refreshToken, _ := s.auth.GenerateRefreshToken(user.Id)

	newToken := model.Token{
		Token:  refreshToken,
		UserId: user.Id,
	}
	_, err := s.repo.AddToken(ctx, newToken)
	if err != nil {
		return model.User{}, "", "", errors.New("error in adding token")
	}

	user.Password = ""
	return user, accessToken, refreshToken, nil
}

func (s *UserService) TokenRefresh(ctx context.Context, token string) (string, error) {
	tokenFound, err := s.repo.GetToken(ctx, token)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	accessToken, err := s.auth.GenerateAccessToken(tokenFound.UserId)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return accessToken, err

}
