package event

import "github.com/google/uuid"

type VacancyCreated struct {
	UserId    string
	VacancyId string
}

func NewVacancyCreated(userId uuid.UUID, vacancyId uuid.UUID) VacancyCreated {
	return VacancyCreated{UserId: userId.String(), VacancyId: vacancyId.String()}
}
