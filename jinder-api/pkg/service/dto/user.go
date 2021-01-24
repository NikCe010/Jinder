package dto

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	Person      `json:"person"`
	Credentials `json:"credentials"`

	Role `json:"role"`
}

func NewUser(person Person, credentials Credentials, role Role) *User {
	return &User{Id: uuid.New(), Person: person, Credentials: credentials, Role: role}
}

type Person struct {
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Birthday time.Time `json:"birthday"`
}

type Credentials struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type Role int

const (
	Recruiter Role = iota
	Programmer
)
