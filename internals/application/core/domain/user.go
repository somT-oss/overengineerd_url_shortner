package domain

import "time"

type User struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

func NewUser(name string, email string, password string) User {
	return User{
		Name: name,
		Email: email,
		Password: password,
		CreatedAt: int(time.Now().Unix()),
		UpdatedAt: int(time.Now().Unix()),
	}
} 