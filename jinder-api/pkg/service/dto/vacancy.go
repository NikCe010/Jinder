package dto

import (
	"Jinder/jinder-api/pkg/service/dto/shared"
	"github.com/google/uuid"
)

type Vacancy struct {
	Id                        uuid.UUID `json:"id"`
	UserId                    uuid.UUID `json:"user_id"`
	shared.ProgrammerLevel    `json:"programmer_level"`
	shared.ProgrammerType     `json:"programmer_type"`
	shared.ProgrammerLanguage `json:"programmer_language"`
	CompanyName               string `json:"company_name"`
	SalaryFrom                string `json:"salary_from"`
	SalaryTo                  string `json:"salary_to"`
	ExtraBenefits             string `json:"extra_benefits"`
}
