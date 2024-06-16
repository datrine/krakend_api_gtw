package sharedexports

import "time"

type CreateUser struct {
	Email        string `json:"email"`
	PasswashHash string `json:"password_hash"`
	LastName     string `json:"last_name"`
	FirstName    string `json:"first_name"`
}

type UpdateUser struct {
	PasswashHash string `json:"password_hash"`
	LastName     string `json:"last_name"`
	FirstName    string `json:"first_name"`
}

type User struct {
	Email        string    `json:"email"`
	PasswashHash string    `json:"password_hash"`
	LastName     string    `json:"last_name"`
	FirstName    string    `json:"first_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
