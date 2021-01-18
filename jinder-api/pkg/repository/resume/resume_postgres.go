package repository

import (
	"Jinder/jinder-api/pkg/domain/profile"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type Postgres struct {
	db *sqlx.DB
}

func (p Postgres) Get(resumeId uuid.UUID) (profile.Resume, error) {
	resume := new(profile.Resume)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	getQuery := fmt.Sprintf("SELECT id, user_id, programmer_level, programmer_type, programmer_language "+
		"FROM %s WHERE id = ?", "resumes")
	err := p.db.QueryRowContext(ctx, getQuery, resumeId).
		Scan(&resume.Id, &resume.UserId, &resume.ProgrammerLevel, &resume.ProgrammerType, &resume.ProgrammerLanguage)

	if err != nil {
		return *resume, err
	}
	return *resume, nil
}

func (p Postgres) GetAll(userId uuid.UUID) ([]profile.Resume, error) {
	users := make([]profile.Resume, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	getAllQuery := fmt.Sprintf("SELECT id, user_id, programmer_level, programmer_type, programmer_language "+
		"FROM %s WHERE user_id = ?", "resumes")
	rows, err := p.db.QueryContext(ctx, getAllQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		resume := new(profile.Resume)
		err = rows.Scan(
			&resume.Id,
			&resume.UserId,
			&resume.ProgrammerLevel,
			&resume.ProgrammerType,
			&resume.ProgrammerLanguage,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, *resume)
	}

	return users, nil
}

func (p Postgres) Create(resume profile.Resume) (uuid.UUID, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return uuid.UUID{}, err
	}
	defer tx.Rollback()

	var itemId uuid.UUID
	createItemQuery := fmt.Sprintf("INSERT INTO %s (id, user_id, programmer_level, programmer_type, "+
		"programmer_language) VALUES (?,?,?,?,?)", "resumes")
	stmt, err := tx.Prepare(createItemQuery)
	if err != nil {
		return uuid.UUID{}, err
	}

	defer stmt.Close()
	_, err = stmt.Exec(resume.Id, resume.UserId, resume.ProgrammerLevel, resume.ProgrammerType, resume.ProgrammerLanguage)
	if err != nil {
		tx.Rollback()
		return uuid.UUID{}, err
	}

	return itemId, tx.Commit()
}

func (p Postgres) Update(resume profile.Resume) (uuid.UUID, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return uuid.UUID{}, err
	}
	defer tx.Rollback()

	var itemId uuid.UUID
	createItemQuery := fmt.Sprintf("UPDATE %s SET user_id=?, programmer_level=?, programmer_type=?, "+
		"programmer_language=? WHERE id=?", "resumes")
	stmt, err := tx.Prepare(createItemQuery)
	if err != nil {
		return uuid.UUID{}, err
	}

	defer stmt.Close()
	_, err = stmt.Exec(resume.UserId, resume.ProgrammerLevel, resume.ProgrammerType, resume.ProgrammerLanguage, resume.Id)
	if err != nil {
		tx.Rollback()
		return uuid.UUID{}, err
	}

	return itemId, tx.Commit()
}

func (p Postgres) Delete(resumeId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	deleteItemQuery := fmt.Sprintf("DELETE FROM %s WHERE id=?", "resumes")
	stmt, err := p.db.PrepareContext(ctx, deleteItemQuery)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, resumeId)
	return err
}

func NewResumePostgres(db *sqlx.DB) *Postgres {
	return &Postgres{db: db}
}
