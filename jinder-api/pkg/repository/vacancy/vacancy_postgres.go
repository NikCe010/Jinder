package vacancy

import (
	"Jinder/jinder-api/pkg/domain/profile"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sqlx.DB
}

func (p Postgres) Get(vacancyId uuid.UUID) (profile.Vacancy, error) {
	panic("implement me")
}

func (p Postgres) GetWithPaging(userId uuid.UUID, count int, page int) ([]profile.Vacancy, error) {
	panic("implement me")
}

func (p Postgres) Create(vacancy profile.Vacancy) error {
	panic("implement me")
}

func (p Postgres) Update(vacancy profile.Vacancy) error {
	panic("implement me")
}

func (p Postgres) Delete(vacancyId uuid.UUID) error {
	panic("implement me")
}

func NewPostgres(db *sqlx.DB) *Postgres {
	return &Postgres{db: db}
}
