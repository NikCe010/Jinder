package infrastructure

type Infrastructure struct {
	Jobs
}

func NewInfrastructure(url string) *Infrastructure {
	return &Infrastructure{Jobs: NewJobsClient(url)}
}

type Jobs interface {
	GetVacancies()
	GetResumes()
}
