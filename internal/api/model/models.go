package model

type User struct {
	Id       string `json:"_id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Token struct {
	Id     string `json:"_id,omitempty"`
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}
