package userRepository

import (
	"go-project/internal/api/model"
)

var userDB []model.User

type UserRepository interface {
}

func CreateNewUser(newUser model.User) bool {
	//userDB = append(userDB, newUser)

	return true
}

func GetUser(email string) (model.User, bool) {
	//for _, user := range userDB {
	//	if user.Email == email {
	//		return user, true
	//	}
	//}
	return model.User{}, false
}
