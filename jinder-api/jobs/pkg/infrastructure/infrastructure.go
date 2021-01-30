package infrastructure

import (
	"Jinder/jinder-api/jobs/pkg/infrastructure/event"
	"github.com/nats-io/nats.go"
)

type Config struct {
	NatsUrl string
}

type Infrastructure struct {
	Recommendation
}

func NewInfrastructure(conn *nats.Conn) *Infrastructure {
	return &Infrastructure{Recommendation: NewRecommendationClient(conn)}
}

type Recommendation interface {
	NotifyWhenVacancyAdded(vacancyCreated event.VacancyCreated) error
	NotifyWhenResumeAdded(resumeCreated event.ResumeCreated) error
	NotifyWhenVacancyViewed(vacancyCreated event.VacancyViewed) error
	NotifyWhenResumeViewed(resumeCreated event.ResumeViewed) error
}
