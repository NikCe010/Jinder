package models

import (
	"Jinder/jinder-api/jobs/pkg/domain/profile"
	domain "Jinder/jinder-api/jobs/pkg/domain/profile/shared"
	"Jinder/jinder-api/jobs/pkg/service/dto/shared"
	"Jinder/jinder-api/jobs/pkg/service/dto/vacancy"
	"github.com/google/uuid"
)

var vacancyId = uuid.New()

var VacancyDto = vacancy.Vacancy{
	Id:                 vacancyId,
	UserId:             uuid.New(),
	ProgrammerLevel:    shared.Middle,
	ProgrammerType:     shared.Backend,
	ProgrammerLanguage: shared.Golang,
	CompanyName:        "Test Company",
	SalaryFrom:         "150000",
	SalaryTo:           "200000",
	ExtraBenefits:      "Medical Insurance, paid vacation 31 days",
}

var Vacancy = profile.Vacancy{
	Id:                 vacancyId,
	UserId:             uuid.New(),
	ProgrammerLevel:    domain.Middle,
	ProgrammerType:     domain.Backend,
	ProgrammerLanguage: domain.Golang,
	CompanyName:        "Test Company",
	SalaryFrom:         "150000",
	SalaryTo:           "200000",
	ExtraBenefits:      "Medical Insurance, paid vacation 31 days",
}
