//Repository for Resume functionality using postgresSQL
package resume

import (
	"Jinder/jinder-api/pkg/domain/profile"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sqlx.DB
}

func (p Postgres) Get(resumeId uuid.UUID) (profile.Resume, error) {
	panic("implement me")
}

func (p Postgres) GetAll(userId uuid.UUID) ([]profile.Resume, error) {
	panic("implement me")
}

func (p Postgres) Create(resume profile.Resume) error {
	panic("implement me")
}

func (p Postgres) Update(resume profile.Resume) error {
	panic("implement me")
}

func (p Postgres) Delete(resumeId uuid.UUID) error {
	panic("implement me")
}

func NewPostgres(db *sqlx.DB) *Postgres {
	return &Postgres{db: db}
}
