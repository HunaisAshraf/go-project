package model

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Posts struct {
	Id     string `json:"id,omitempty"`
	Title  string `json:"title"`
	Body   string `json:"Body"`
	UserId string `json:"userId"`
}
