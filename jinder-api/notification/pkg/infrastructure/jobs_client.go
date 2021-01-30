package infrastructure

import "net/http"

type JobsClient struct {
	http.Client
}

func NewJobsClient(url string) *JobsClient {
	return &JobsClient{}
}

func (j JobsClient) GetVacancies() {
	panic("implement me")
}

func (j JobsClient) GetResumes() {
	panic("implement me")
}
