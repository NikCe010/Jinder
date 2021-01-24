package repository

import (
	"Jinder/jinder-api/pkg/domain/profile"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
)

type ResumePostgres struct {
	db *sqlx.DB
}

func (p ResumePostgres) Get(resumeId uuid.UUID) (profile.Resume, error) {
	resume := new(profile.Resume)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	getQuery := fmt.Sprintf("SELECT id, user_id, programmer_level, programmer_type, programmer_language "+
		"FROM %s WHERE id = $1", "resumes")
	err := p.db.QueryRowContext(ctx, getQuery, resumeId).
		Scan(&resume.Id, &resume.UserId, &resume.ProgrammerLevel, &resume.ProgrammerType, &resume.ProgrammerLanguage)

	if err != nil {
		log.Error(err.Error())
		return *resume, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	getAllQuery := fmt.Sprintf("SELECT id, resume_id, company_name, experience_from, experience_to, extra_content "+
		"FROM %s where resume_id = $1", "work_experience")
	rows, err := p.db.QueryContext(ctx, getAllQuery, resumeId)
	if err != nil {
		log.Error(err.Error())
		log.Error(getAllQuery)
		return *resume, err
	}
	defer rows.Close()

	for rows.Next() {
		var workExp profile.WorkExperience
		err = rows.Scan(
			&workExp.Id,
			&workExp.ResumeId,
			&workExp.CompanyName,
			&workExp.From,
			&workExp.To,
			&workExp.Content,
		)

		if err != nil {
			log.Error(err.Error())
			return *resume, err
		}
		resume.WorkExperiences = append(resume.WorkExperiences, workExp)
	}

	return *resume, nil
}

func (p ResumePostgres) GetWithPaging(userId uuid.UUID, count int, page int) ([]profile.Resume, error) {
	resumes := make([]profile.Resume, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	getAllQuery := fmt.Sprintf("SELECT id, user_id, programmer_level, programmer_type, programmer_language "+
		"FROM %s WHERE user_id = ?"+
		"OFFSET %d LIMIT %d", "resumes", page*count, count)
	rows, err := p.db.QueryContext(ctx, getAllQuery, userId)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var resume profile.Resume
		err = rows.Scan(
			&resume.Id,
			&resume.UserId,
			&resume.ProgrammerLevel,
			&resume.ProgrammerType,
			&resume.ProgrammerLanguage,
		)

		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		resumes = append(resumes, resume)
	}

	return resumes, nil
}

func (p ResumePostgres) Create(resume profile.Resume) (uuid.UUID, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}
	defer tx.Rollback()

	createItemCommand := fmt.Sprintf("INSERT INTO %s (id, user_id, programmer_level, programmer_type, "+
		"programmer_language) VALUES ($1,$2,$3,$4,$5)", "resumes")
	_, err = tx.Exec(createItemCommand, resume.Id, resume.UserId, resume.ProgrammerLevel, resume.ProgrammerType, resume.ProgrammerLanguage)
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	for _, w := range resume.WorkExperiences {
		createWorkExpCommand := fmt.Sprintf("INSERT INTO %s (id, resume_id, company_name, experience_from, "+
			"experience_to, extra_content) VALUES ($1,$2,$3,$4,$5, $6)", "work_experience")
		_, err = tx.Exec(createWorkExpCommand, w.Id, w.ResumeId, w.CompanyName, w.From, w.To, w.Content)
		if err != nil {
			log.Error(err.Error())
			return uuid.UUID{}, err
		}
	}

	return resume.Id, tx.Commit()
}

func (p ResumePostgres) Update(resume profile.Resume) (uuid.UUID, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}
	defer tx.Rollback()

	updateItemCommand := fmt.Sprintf("UPDATE %s SET user_id=$1, programmer_level=$2, programmer_type=$3, "+
		"programmer_language=$4 WHERE id=$5", "resumes")

	_, err = tx.Exec(updateItemCommand, resume.UserId, resume.ProgrammerLevel, resume.ProgrammerType, resume.ProgrammerLanguage, resume.Id)
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		return uuid.UUID{}, err
	}

	for _, w := range resume.WorkExperiences {
		updateWorkExpCommand := fmt.Sprintf("UPDATE %s SET resume_id=$1, company_name=$2, experience_from=$3, "+
			"experience_to=$4, extra_content=$5 where id=$6", "work_experience")
		_, err = tx.Exec(updateWorkExpCommand, w.ResumeId, w.CompanyName, w.From, w.To, w.Content, w.Id)
		if err != nil {
			log.Error(err.Error())
			return uuid.UUID{}, err
		}
	}

	return resume.Id, tx.Commit()
}

func (p ResumePostgres) Delete(resumeId uuid.UUID) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer tx.Rollback()

	deleteItemQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1", "resumes")
	_, err = tx.Exec(deleteItemQuery, resumeId)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	deleteWorkExpCommand := fmt.Sprintf("DELETE FROM %s WHERE resume_id=$1", "work_experience")
	_, err = tx.Exec(deleteWorkExpCommand, resumeId)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return tx.Commit()
}

func NewResumePostgres(db *sqlx.DB) *ResumePostgres {
	return &ResumePostgres{db: db}
}
