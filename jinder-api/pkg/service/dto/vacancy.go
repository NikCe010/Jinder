package dto

import (
	"Jinder/jinder-api/pkg/service/dto/shared"
	"github.com/google/uuid"
)

type Vacancy struct {
	Id     uuid.UUID
	UserId uuid.UUID
	shared.ProgrammerLevel
	shared.ProgrammerType
	shared.ProgrammerLanguage
	CompanyName   string
	SalaryFrom    string
	SalaryTo      string
	OtherBenefits string
}

type FieldOfActivity int

const (
	Medicine FieldOfActivity = iota
	Banking
	Gaming
	Marketing
	Streaming
	Blockchain
)
