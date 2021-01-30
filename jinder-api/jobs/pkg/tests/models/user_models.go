package models

import (
	"Jinder/jinder-api/jobs/pkg/domain/registration"
	"Jinder/jinder-api/jobs/pkg/service/dto/user"
	"github.com/google/uuid"
	"time"
)

var UserDto = user.User{
	Id: uuid.New(),
	Person: user.Person{
		Name:     "Vasya",
		Surname:  "Vasyliev",
		Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	Credentials: user.Credentials{
		Email:                "testemail@mail.ru",
		Password:             "qwerty123",
		PasswordConfirmation: "qwerty123",
	},
	Role: user.Programmer,
}

var UserDomain = registration.User{
	Id: uuid.New(),
	Person: registration.Person{
		Name:     "Vasya",
		Surname:  "Vasyliev",
		Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	Credentials: registration.Credentials{
		Email:        "test@mail.ru",
		PasswordHash: "$5$MnfsQ4iN$ZMTppKN16y/tIsUYs/obHlhdP.Os80yXhTurpBMUbA5",
	},
	Role: registration.Programmer,
}

var UserErrorDto = user.User{
	Id: uuid.New(),
	Person: user.Person{
		Name:     "Vasya",
		Surname:  "Vasyliev",
		Birthday: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	Credentials: user.Credentials{
		Email:                "testemail@mail.ru",
		Password:             "qwerty123",
		PasswordConfirmation: "qwe123rty",
	},
	Role: user.Programmer,
}
