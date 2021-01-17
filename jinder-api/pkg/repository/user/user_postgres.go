package user

import (
	"Jinder/jinder-api/pkg/domain/registration"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sqlx.DB
}

func (p Postgres) Register(user registration.User) error {
	panic("implement me")
}

func (p Postgres) Update(user registration.User) error {
	panic("implement me")
}

func (p Postgres) Get(userId uuid.UUID) (registration.User, error) {
	panic("implement me")
}

func NewPostgres(db *sqlx.DB) *Postgres {
	return &Postgres{db: db}
}
