package event

import "github.com/google/uuid"

type ResumeCreated struct {
	UserId   string
	ResumeId string
}

func NewResumeCreated(userId uuid.UUID, resumeId uuid.UUID) ResumeCreated {
	return ResumeCreated{UserId: userId.String(), ResumeId: resumeId.String()}
}
