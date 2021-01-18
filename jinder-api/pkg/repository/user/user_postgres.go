package user

import (
	"Jinder/jinder-api/pkg/domain/registration"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type UserPostgres struct {
	db *sqlx.DB
}

func (p UserPostgres) Register(user registration.User) (uuid.UUID, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return uuid.UUID{}, err
	}

	defer tx.Rollback()

	createItemQery := fmt.Sprintf("INSERT INTO %s (id, name, surname, birthday, email, password_hash, role) VALUES (?,?,?,?,?,?,?)", "users")
	stmt, err := tx.Prepare(createItemQery)
	if err != nil {
		return uuid.UUID{}, err
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Name, user.Surname, user.Birthday, user.Email, user.PasswordHash, user.Role)
	if err != nil {
		tx.Rollback()
		return uuid.UUID{}, err
	}
	return user.Id, tx.Commit()
}

func (p UserPostgres) Update(user registration.User) (uuid.UUID, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return uuid.UUID{}, err
	}

	defer tx.Rollback()

	updateQuery := fmt.Sprintf("UPDATE %s SET name=?, surname=?, birthday=?, email=?, password_hash=?, role=? WHERE id=?", "users")
	stmt, err := tx.Prepare(updateQuery)
	if err != nil {
		return uuid.UUID{}, err
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Surname, user.Birthday, user.Email, user.PasswordHash, user.Role, user.Id)
	if err != nil {
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
		"FROM %s WHERE id = ?", "users")
	err := p.db.QueryRowContext(ctx, getQuery, userId).
		Scan(&user.Id, &user.Name, &user.Surname, &user.Birthday, &user.Email, &user.PasswordHash, &user.Role)

	if err != nil {
		return *user, err
	}
	return *user, nil
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}
