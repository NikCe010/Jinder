package event

import "github.com/google/uuid"

type ResumeViewed struct {
	UserId   string
	ResumeId string
}

func NewResumeViewed(userId uuid.UUID, resumeId uuid.UUID) ResumeViewed {
	return ResumeViewed{UserId: userId.String(), ResumeId: resumeId.String()}
}
