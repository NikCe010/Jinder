package profile

import (
	"Jinder/jinder-api/pkg/domain/profile/shared"
	"fmt"
	"github.com/google/uuid"
	"strings"
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

func NewResume(userId uuid.UUID, programmerLevel shared.ProgrammerLevel, programmerType shared.ProgrammerType,
	programmerLanguage shared.ProgrammerLanguage, workExperiences []WorkExperience) *Resume {
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

func (r *Resume) ExtraSkillsToText() string {
	if len(r.ExtraSkills) <= 0 {
		return ""
	}

	var sb strings.Builder
	if _, err := sb.WriteString("Extra skills:\n"); err != nil {
		return ""
	}

	for _, e := range r.ExtraSkills {
		sb.WriteString(fmt.Sprintf("%s;\n", e))
	}
	return sb.String()
}
