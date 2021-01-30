package mocks

import (
	"Jinder/jinder-api/jobs/pkg/domain/registration"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"time"
)

type MockUserRepository struct {
	mock.Mock
}

func (m MockUserRepository) Register(user registration.User) (uuid.UUID, error) {
	return user.Id, nil
}

func (m MockUserRepository) Update(user registration.User) (uuid.UUID, error) {
	return user.Id, nil
}

func (m MockUserRepository) Get(userId uuid.UUID) (registration.User, error) {
	return registration.User{
		Id: userId,
		Person: registration.Person{
			Name:     "Vasya",
			Surname:  "Vasyliev",
			Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		Credentials: registration.Credentials{
			Email:        "testemail@mail.ru",
			PasswordHash: "qwerty123",
		},
		Role: registration.Programmer,
	}, nil
}

func (m MockUserRepository) GetByEmail(email string) (registration.User, error) {
	if email == "test@test.ru" {
		id, _ := uuid.Parse("25790180-5c2d-11eb-ae93-0242ac131111")
		return registration.User{
			Id: id,
			Person: registration.Person{
				Name:     "Vasya",
				Surname:  "Vasyliev",
				Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			Credentials: registration.Credentials{
				Email:        email,
				PasswordHash: "$2a$04$w8/XJrPLfa9K6uDdxynzc.eCL/kA2Yuonu8Y.1lZAZNr0N.OXdGTO",
			},
			Role: registration.Programmer,
		}, nil
	} else {
		return registration.User{
			Id: uuid.UUID{},
			Person: registration.Person{
				Name:     "Vasya",
				Surname:  "Vasyliev",
				Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			Credentials: registration.Credentials{
				Email:        email,
				PasswordHash: "$2a$04$v89aJOj.BDmu8e/ict5IOXDHrigEZHsFq0/yjw5gXk6KUe7vVBP.",
			},
			Role: registration.Programmer,
		}, nil
	}

}
