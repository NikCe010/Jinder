package registration

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id uuid.UUID
	Person
	Credentials

	Role
}

func NewUser(person Person, credentials Credentials, role Role) *User {
	return &User{Id: uuid.New(), Person: person, Credentials: credentials, Role: role}
}

type Person struct {
	Name     string
	Surname  string
	Birthday time.Time
}

type Credentials struct {
	Email                string
	Password             string
	PasswordConfirmation string
}

type Role string

const (
	Recruiter  = "Recruiter"
	Programmer = "Programmer"
)
