package repository

import (
	"Jinder/jinder-api/jobs/pkg/domain/registration"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
)

type UserPostgres struct {
	db *sqlx.DB
}

func (p UserPostgres) Register(user registration.User) (uuid.UUID, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	defer tx.Rollback()

	createItemQuery := fmt.Sprintf("INSERT INTO %s (id, name, surname, birthday, email, password_hash, role) "+
		"VALUES ($1,$2,$3,$4,$5,$6,$7)", UsersTable)
	stmt, err := tx.Prepare(createItemQuery)

	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Name, user.Surname, user.Birthday, user.Email, user.PasswordHash, user.Role)
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		return uuid.UUID{}, err
	}
	return user.Id, tx.Commit()
}

func (p UserPostgres) Update(user registration.User) (uuid.UUID, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	defer tx.Rollback()

	updateQuery := fmt.Sprintf("UPDATE %s SET name=$1, surname=$2, birthday=$3, email=$4, password_hash=$5, "+
		"role=$6 WHERE id=$7", UsersTable)
	stmt, err := tx.Prepare(updateQuery)
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Surname, user.Birthday, user.Email, user.PasswordHash, user.Role, user.Id)
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		return uuid.UUID{}, err
	}
	return user.Id, tx.Commit()
}

func (p UserPostgres) Get(userId uuid.UUID) (registration.User, error) {
	user := new(registration.User)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	getQuery := fmt.Sprintf("SELECT id, name, surname, birthday, email, password_hash, role "+
		"FROM %s WHERE id = $1", UsersTable)
	err := p.db.QueryRowContext(ctx, getQuery, userId).
		Scan(&user.Id, &user.Name, &user.Surname, &user.Birthday, &user.Email, &user.PasswordHash, &user.Role)

	if err != nil {
		log.Error(err.Error())
		return *user, err
	}
	return *user, nil
}

func (p UserPostgres) GetByEmail(email string) (registration.User, error) {
	user := new(registration.User)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	getQuery := fmt.Sprintf("SELECT id, name, surname, birthday, email, password_hash, role "+
		"FROM %s WHERE email = $1", UsersTable)
	err := p.db.QueryRowContext(ctx, getQuery, email).
		Scan(&user.Id, &user.Name, &user.Surname, &user.Birthday, &user.Email, &user.PasswordHash, &user.Role)

	if err != nil {
		log.Error(err.Error())
		return *user, err
	}
	return *user, nil
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}
