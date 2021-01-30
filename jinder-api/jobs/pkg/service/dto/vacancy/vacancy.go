package vacancy

import (
	"Jinder/jinder-api/jobs/pkg/domain/profile"
	domain "Jinder/jinder-api/jobs/pkg/domain/profile/shared"
	jobs_service "Jinder/jinder-api/jobs/pkg/handler/grpc/protos"
	"Jinder/jinder-api/jobs/pkg/service/dto/shared"
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

func ToDomain(vacancy Vacancy) profile.Vacancy {
	return profile.Vacancy{
		Id:                 vacancy.Id,
		UserId:             vacancy.UserId,
		ProgrammerLanguage: domain.ProgrammerLanguage(vacancy.ProgrammerLanguage),
		ProgrammerLevel:    domain.ProgrammerLevel(vacancy.ProgrammerLevel),
		ProgrammerType:     domain.ProgrammerType(vacancy.ProgrammerType),
		CompanyName:        vacancy.CompanyName,
		SalaryFrom:         vacancy.SalaryFrom,
		SalaryTo:           vacancy.SalaryTo,
		ExtraBenefits:      vacancy.ExtraBenefits,
	}
}

func ToDto(vacancy profile.Vacancy) Vacancy {
	return Vacancy{
		Id:                 vacancy.Id,
		UserId:             vacancy.UserId,
		ProgrammerLanguage: shared.ProgrammerLanguage(vacancy.ProgrammerLanguage),
		ProgrammerLevel:    shared.ProgrammerLevel(vacancy.ProgrammerLevel),
		ProgrammerType:     shared.ProgrammerType(vacancy.ProgrammerType),
		CompanyName:        vacancy.CompanyName,
		SalaryFrom:         vacancy.SalaryFrom,
		SalaryTo:           vacancy.SalaryTo,
		ExtraBenefits:      vacancy.ExtraBenefits,
	}
}

func ToGrpc(vacancy Vacancy) *jobs_service.VacancyDto {
	return &jobs_service.VacancyDto{
		Id:            vacancy.Id.String(),
		UserId:        vacancy.UserId.String(),
		Level:         jobs_service.ProgrammerLevel(vacancy.ProgrammerLevel),
		Type:          jobs_service.ProgrammerType(vacancy.ProgrammerType),
		Language:      jobs_service.ProgrammerLanguage(vacancy.ProgrammerLanguage),
		SalaryFrom:    vacancy.SalaryFrom,
		SalaryTo:      vacancy.SalaryTo,
		ExtraBenefits: vacancy.ExtraBenefits,
	}
}

func GrpcToDto(vacancy *jobs_service.VacancyDto) (Vacancy, error) {
	vacancyId, err := uuid.Parse(vacancy.Id)
	if err != nil {
		return Vacancy{}, err
	}
	userId, err := uuid.Parse(vacancy.Id)
	if err != nil {
		return Vacancy{}, err
	}

	return Vacancy{
		Id:                 vacancyId,
		UserId:             userId,
		ProgrammerLanguage: shared.ProgrammerLanguage(vacancy.Language),
		ProgrammerLevel:    shared.ProgrammerLevel(vacancy.Level),
		ProgrammerType:     shared.ProgrammerType(vacancy.Type),
		CompanyName:        vacancy.CompanyName,
		SalaryFrom:         vacancy.SalaryFrom,
		SalaryTo:           vacancy.SalaryTo,
		ExtraBenefits:      vacancy.ExtraBenefits,
	}, nil
}
