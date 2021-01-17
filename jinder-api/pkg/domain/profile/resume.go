package profile

import (
	"Jinder/jinder-api/pkg/domain/profile/shared"
	"github.com/google/uuid"
	"time"
)

type Resume struct {
	Id     uuid.UUID
	UserId uuid.UUID
	shared.ProgrammerLevel
	shared.ProgrammerType
	shared.ProgrammerLanguage
	ExtraSkills     []string
	WorkExperiences []WorkExperience //Last 3 places
}

func NewResume(userId uuid.UUID,
	programmerLevel shared.ProgrammerLevel,
	programmerType shared.ProgrammerType,
	programmerLanguage shared.ProgrammerLanguage,
	workExperiences []WorkExperience) *Resume {
	return &Resume{Id: uuid.New(), UserId: userId, ProgrammerLevel: programmerLevel, ProgrammerType: programmerType,
		ProgrammerLanguage: programmerLanguage, WorkExperiences: workExperiences}
}

type WorkExperience struct {
	Id       uuid.UUID
	ResumeId uuid.UUID

	CompanyName string
	From        time.Time
	To          time.Time
	Content     string // len = twitters
}
