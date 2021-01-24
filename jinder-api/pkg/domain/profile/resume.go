package profile

import (
	"Jinder/jinder-api/pkg/domain/profile/shared"
	"github.com/google/uuid"
	"time"
)

type Resume struct {
	Id                        uuid.UUID `db:"id"`
	UserId                    uuid.UUID `db:"user_id"`
	shared.ProgrammerLevel    `db:"programmer_level"`
	shared.ProgrammerType     `db:"programmer_type"`
	shared.ProgrammerLanguage `db:"programmer_language"`
	WorkExperiences           []WorkExperience
}

func NewResume(userId uuid.UUID, programmerLevel shared.ProgrammerLevel, programmerType shared.ProgrammerType,
	programmerLanguage shared.ProgrammerLanguage, workExperiences []WorkExperience) *Resume {
	return &Resume{Id: uuid.New(), UserId: userId, ProgrammerLevel: programmerLevel, ProgrammerType: programmerType,
		ProgrammerLanguage: programmerLanguage, WorkExperiences: workExperiences}
}

type WorkExperience struct {
	Id       uuid.UUID `db:"id"`
	ResumeId uuid.UUID `db:"resume_id"`

	CompanyName string    `db:"company_name"`
	From        time.Time `db:"experience_from"`
	To          time.Time `db:"experience_to"`
	Content     string    `db:"content"`
}
