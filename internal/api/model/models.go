package model

type User struct {
	Id       string `json:"_id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
