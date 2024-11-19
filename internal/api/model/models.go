package model

type User struct {
	Id       string `json:"_id,omitempty" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Phone    string `json:"phone" bson:"phone"`
	Password string `json:"password" bson:"password"`
}

type Token struct {
	Id     string `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId string `json:"user_id" bson:"userid"`
	Token  string `json:"token" bson:"token"`
}
