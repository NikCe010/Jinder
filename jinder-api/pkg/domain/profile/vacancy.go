package profile

import (
	"Jinder/jinder-api/pkg/domain/profile/shared"
	"github.com/google/uuid"
)

type Vacancy struct {
	Id     uuid.UUID
	UserId uuid.UUID
	shared.ProgrammerLevel
	shared.ProgrammerType
	shared.ProgrammerLanguage
	CompanyName   string
	Team          string
	SalaryFrom    string
	SalaryTo      string
	OtherBenefits string
}

type FieldOfActivity string

const (
	Medicine   = "Medicine"
	Banking    = "Banking"
	Gaming     = "Gaming"
	Marketing  = "Marketing"
	Streaming  = "Streaming"
	Blockchain = "Blockchain"
)
