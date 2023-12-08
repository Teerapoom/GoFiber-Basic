package model

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var User_der = User{
	Email:    "userTest@mail.com",
	Password: "0872746355",
}
