package dto

import (
	"Jinder/jinder-api/pkg/service/dto/shared"
	"github.com/google/uuid"
	"time"
)

type Resume struct {
	Id                        uuid.UUID `json:"id"`
	UserId                    uuid.UUID `json:"user_id"`
	shared.ProgrammerLevel    `json:"programmer_level"`
	shared.ProgrammerType     `json:"programmer_type"`
	shared.ProgrammerLanguage `json:"programmer_language"`
	WorkExperiences           []WorkExperience `json:"work_experiences"` //Last 3 places
}

func NewResume(userId uuid.UUID, programmerLevel shared.ProgrammerLevel, programmerType shared.ProgrammerType,
	programmerLanguage shared.ProgrammerLanguage, workExperiences []WorkExperience) *Resume {
	return &Resume{Id: uuid.New(), UserId: userId, ProgrammerLevel: programmerLevel, ProgrammerType: programmerType,
		ProgrammerLanguage: programmerLanguage, WorkExperiences: workExperiences}
}

type WorkExperience struct {
	Id       uuid.UUID `json:"id"`
	ResumeId uuid.UUID `json:"resume_id"`

	CompanyName string    `json:"company_name"`
	From        time.Time `json:"from"`
	To          time.Time `json:"to"`
	Content     string    `json:"content"` // len = twitters
}
