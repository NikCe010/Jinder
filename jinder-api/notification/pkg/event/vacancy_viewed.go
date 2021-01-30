package event

import "github.com/google/uuid"

type VacancyViewed struct {
	UserId    string
	VacancyId string
}

func NewVacancyViewed(userId uuid.UUID, vacancyId uuid.UUID) VacancyViewed {
	return VacancyViewed{UserId: userId.String(), VacancyId: vacancyId.String()}
}
