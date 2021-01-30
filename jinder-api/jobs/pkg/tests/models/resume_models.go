package models

import (
	"Jinder/jinder-api/jobs/pkg/domain/profile"
	domain "Jinder/jinder-api/jobs/pkg/domain/profile/shared"
	"Jinder/jinder-api/jobs/pkg/service/dto/resume"
	"Jinder/jinder-api/jobs/pkg/service/dto/resume/work_experience"
	"Jinder/jinder-api/jobs/pkg/service/dto/shared"
	"github.com/google/uuid"
	"time"
)

var (
	resumeId = uuid.New()
)

var ResumeDto = resume.Resume{
	Id:                 resumeId,
	UserId:             uuid.New(),
	ProgrammerLevel:    shared.Middle,
	ProgrammerType:     shared.Backend,
	ProgrammerLanguage: shared.Golang,
	WorkExperiences: []work_experience.WorkExperience{
		{Id: uuid.New(),
			ResumeId:    resumeId,
			CompanyName: "TestCompany",
			From:        time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			To:          time.Now(),
			Content:     "Created a monitoring system, mentored two juniors, migrated the project to gRPS"},
	},
}

var Resume = profile.Resume{
	Id:                 resumeId,
	UserId:             uuid.New(),
	ProgrammerLevel:    domain.Middle,
	ProgrammerType:     domain.Backend,
	ProgrammerLanguage: domain.Golang,
	WorkExperiences: []profile.WorkExperience{
		{Id: uuid.New(),
			ResumeId:    resumeId,
			CompanyName: "TestCompany",
			From:        time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
			To:          time.Now(),
			Content:     "Created a monitoring system, mentored two juniors, migrated the project to gRPS"},
	},
}
