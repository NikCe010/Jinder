package registration

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
	Name     string    `db:"name"`
	Surname  string    `db:"surname"`
	Birthday time.Time `db:"birthday"`
}

type Credentials struct {
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
}

type Role int

const (
	Recruiter Role = iota
	Programmer
)
