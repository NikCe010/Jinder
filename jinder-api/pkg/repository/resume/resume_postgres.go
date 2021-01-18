package repository

import (
	"Jinder/jinder-api/pkg/domain/profile"
	"fmt"
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

func (p Postgres) Create(resume profile.Resume) (uuid.UUID, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return uuid.UUID{}, err
	}
	defer tx.Rollback()

	var itemId uuid.UUID
	createItemQuery := fmt.Sprintf("INSERT INTO %s (id, user_id, programmer_level, programmer_type, "+
		"programmer_language, extra_skills) VALUES (?,?,?,?,?,?)", "resumes")
	stmt, err := tx.Prepare(createItemQuery)
	if err != nil {
		return uuid.UUID{}, err
	}

	defer stmt.Close()
	_, err = stmt.Exec(resume.Id, resume.UserId, resume.ProgrammerLevel, resume.ProgrammerType, resume.ProgrammerLanguage, resume.ExtraSkillsToText())
	if err != nil {
		tx.Rollback()
		return uuid.UUID{}, err
	}

	return itemId, tx.Commit()
}

func (p Postgres) Update(resume profile.Resume) error {
	panic("implement me")
}

func (p Postgres) Delete(resumeId uuid.UUID) error {
	panic("implement me")
}

func NewResumePostgres(db *sqlx.DB) *Postgres {
	return &Postgres{db: db}
}
