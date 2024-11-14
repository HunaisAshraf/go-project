package userServices

import (
	"errors"
	"go-project/internal/api/model"
	"go-project/internal/api/repository"
	"go-project/util"
)

func SignupUser(user model.User) (model.User, error) {

	if _, exist := userRepository.GetUser(user.Email); exist {
		return model.User{}, errors.New("user already exist ")
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return model.User{}, errors.New("hashing password failed")
	}

	user.Password = hashedPassword
	success := userRepository.CreateNewUser(user)
	if !success {
		return model.User{}, errors.New("error in adding user")
	}

	user.Password = ""
	return user, nil
}

func LoginUser(email string, password string) (model.User, error) {
	user, exist := userRepository.GetUser(email)
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
