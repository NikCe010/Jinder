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

type VacancyPostgres struct {
	db *sqlx.DB
}

func (p VacancyPostgres) Get(vacancyId uuid.UUID) (profile.Vacancy, error) {
	vacancy := new(profile.Vacancy)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	getItemQuery := fmt.Sprintf("SELECT id, user_id, programmer_level, programmer_language, "+
		"programmer_type, company_name, salary_from, salary_to, extra_benefits FROM %s WHERE id=$1", "vacancies")

	err := p.db.QueryRowContext(ctx, getItemQuery, vacancyId).
		Scan(&vacancy.Id, &vacancy.UserId, &vacancy.ProgrammerLevel, &vacancy.ProgrammerLanguage, &vacancy.ProgrammerType,
			&vacancy.CompanyName, &vacancy.SalaryFrom, &vacancy.SalaryTo, &vacancy.ExtraBenefits)

	if err != nil {
		log.Error(err.Error())
		return *vacancy, err
	}
	return *vacancy, nil
}

func (p VacancyPostgres) GetWithPaging(userId uuid.UUID, count int, page int) ([]profile.Vacancy, error) {
	vacancies := make([]profile.Vacancy, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	getItemsQuery := fmt.Sprintf("SELECT id, user_id, programmer_type, programmer_level, "+
		"programmer_language, company_name, salary_from, salary_to, extra_benefits FROM %s WHERE user_id=$1 "+
		"OFFSET %d LIMIT %d", "vacancies", page*count, count)
	log.Print(getItemsQuery)

	rows, err := p.db.QueryContext(ctx, getItemsQuery, userId)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	for rows.Next() {
		var vacancy profile.Vacancy
		err = rows.Scan(
			&vacancy.Id,
			&vacancy.UserId,
			&vacancy.ProgrammerLevel,
			&vacancy.ProgrammerLanguage,
			&vacancy.ProgrammerType,
			&vacancy.CompanyName,
			&vacancy.SalaryFrom,
			&vacancy.SalaryTo,
			&vacancy.ExtraBenefits,
		)

		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		vacancies = append(vacancies, vacancy)
	}

	rows.Close()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return vacancies, nil
}

func (p VacancyPostgres) Create(vacancy profile.Vacancy) (uuid.UUID, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}
	defer tx.Rollback()

	createItemQuery := fmt.Sprintf("INSERT INTO %s (id, user_id, programmer_level, programmer_language, "+
		"programmer_type, company_name, salary_from, salary_to, extra_benefits) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)", "vacancies")

	log.Debug(createItemQuery)
	stmt, err := tx.Prepare(createItemQuery)
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}
	_, err = stmt.Exec(vacancy.Id, vacancy.UserId, vacancy.ProgrammerLevel, vacancy.ProgrammerLanguage, vacancy.ProgrammerType,
		vacancy.CompanyName, vacancy.SalaryFrom, vacancy.SalaryTo, vacancy.ExtraBenefits)
	defer stmt.Close()

	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	return vacancy.Id, tx.Commit()
}

func (p VacancyPostgres) Update(vacancy profile.Vacancy) (uuid.UUID, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}
	defer tx.Rollback()

	updateItemQuery := fmt.Sprintf("UPDATE %s SET user_id=$1, programmer_level=$2, programmer_language=$3, "+
		"programmer_type=$4, company_name=$5, salary_from=$6, salary_to=$7, extra_benefits=$8 WHERE id=$9", "vacancies")

	stmt, err := tx.Prepare(updateItemQuery)
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}
	_, err = stmt.Exec(vacancy.UserId, vacancy.ProgrammerLevel, vacancy.ProgrammerLanguage, vacancy.ProgrammerType,
		vacancy.CompanyName, vacancy.SalaryFrom, vacancy.SalaryTo, vacancy.ExtraBenefits, vacancy.Id)
	defer stmt.Close()

	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	return vacancy.Id, tx.Commit()
}

func (p VacancyPostgres) Delete(vacancyId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	deleteItemQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1", "vacancies")
	stmt, err := p.db.PrepareContext(ctx, deleteItemQuery)

	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, vacancyId)
	return err
}

func NewVacancyPostgres(db *sqlx.DB) *VacancyPostgres {
	return &VacancyPostgres{db: db}
}
