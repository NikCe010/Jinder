package profile

import (
	"Jinder/jinder-api/jobs/pkg/domain/profile/shared"
	"github.com/google/uuid"
)

type Vacancy struct {
	Id                        uuid.UUID `db:"id"`
	UserId                    uuid.UUID `db:"user_id"`
	shared.ProgrammerType     `db:"programmer_type"`
	shared.ProgrammerLevel    `db:"programmer_level"`
	shared.ProgrammerLanguage `db:"programmer_language"`
	CompanyName               string `db:"company_name"`
	SalaryFrom                string `db:"salary_from"`
	SalaryTo                  string `db:"salary_to"`
	ExtraBenefits             string `db:"extra_benefits"`
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
