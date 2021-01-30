package resume

import (
	"Jinder/jinder-api/jobs/pkg/domain/profile"
	domain "Jinder/jinder-api/jobs/pkg/domain/profile/shared"
	jobs_service "Jinder/jinder-api/jobs/pkg/handler/grpc/protos"
	"Jinder/jinder-api/jobs/pkg/service/dto/resume/work_experience"
	"Jinder/jinder-api/jobs/pkg/service/dto/shared"
	"github.com/google/uuid"
)

type Resume struct {
	Id                        uuid.UUID `json:"id"`
	UserId                    uuid.UUID `json:"user_id"`
	shared.ProgrammerLevel    `json:"programmer_level"`
	shared.ProgrammerType     `json:"programmer_type"`
	shared.ProgrammerLanguage `json:"programmer_language"`
	WorkExperiences           []work_experience.WorkExperience `json:"work_experiences"` //Last 3 places
}

func NewResume(userId uuid.UUID, programmerLevel shared.ProgrammerLevel, programmerType shared.ProgrammerType,
	programmerLanguage shared.ProgrammerLanguage, workExperiences []work_experience.WorkExperience) *Resume {
	return &Resume{Id: uuid.New(), UserId: userId, ProgrammerLevel: programmerLevel, ProgrammerType: programmerType,
		ProgrammerLanguage: programmerLanguage, WorkExperiences: workExperiences}
}

func ToDomain(resume Resume) profile.Resume {
	return profile.Resume{
		Id:                 resume.Id,
		UserId:             resume.UserId,
		ProgrammerLanguage: domain.ProgrammerLanguage(resume.ProgrammerLanguage),
		ProgrammerLevel:    domain.ProgrammerLevel(resume.ProgrammerLevel),
		ProgrammerType:     domain.ProgrammerType(resume.ProgrammerType),
		WorkExperiences:    work_experience.ToDomain(resume.WorkExperiences),
	}
}

func ToDto(resume profile.Resume) Resume {
	return Resume{
		Id:                 resume.Id,
		UserId:             resume.UserId,
		ProgrammerLanguage: shared.ProgrammerLanguage(resume.ProgrammerLanguage),
		ProgrammerLevel:    shared.ProgrammerLevel(resume.ProgrammerLevel),
		ProgrammerType:     shared.ProgrammerType(resume.ProgrammerType),
		WorkExperiences:    work_experience.ToDto(resume.WorkExperiences),
	}
}

func ToGrpc(resume Resume) (*jobs_service.ResumeDto, error) {
	workExp, err := work_experience.ToGrpc(resume.WorkExperiences)
	if err != nil {
		return nil, err
	}
	return &jobs_service.ResumeDto{
		Id:              resume.Id.String(),
		UserId:          resume.UserId.String(),
		Language:        jobs_service.ProgrammerLanguage(resume.ProgrammerLanguage),
		Level:           jobs_service.ProgrammerLevel(resume.ProgrammerLevel),
		Type:            jobs_service.ProgrammerType(resume.ProgrammerType),
		WorkExperiences: workExp,
	}, nil
}

func GrpcToDto(resume *jobs_service.ResumeDto) (Resume, error) {
	resumeId, err := uuid.Parse(resume.Id)
	if err != nil {
		return Resume{}, err
	}
	userId, err := uuid.Parse(resume.Id)
	if err != nil {
		return Resume{}, err
	}
	workExp, err := work_experience.GrpcToDto(resume.WorkExperiences)
	if err != nil {
		return Resume{}, err
	}

	return Resume{
		Id:                 resumeId,
		UserId:             userId,
		ProgrammerLanguage: shared.ProgrammerLanguage(resume.Language),
		ProgrammerLevel:    shared.ProgrammerLevel(resume.Level),
		ProgrammerType:     shared.ProgrammerType(resume.Type),
		WorkExperiences:    workExp,
	}, nil
}
