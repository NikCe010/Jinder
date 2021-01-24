package resume

import (
	"Jinder/jinder-api/pkg/domain/profile"
	domain "Jinder/jinder-api/pkg/domain/profile/shared"
	"Jinder/jinder-api/pkg/repository"
	"Jinder/jinder-api/pkg/service/dto"
	"Jinder/jinder-api/pkg/service/dto/shared"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type ResumeService struct {
	repo repository.Resume
}

func (s ResumeService) GetResume(resumeId uuid.UUID) (dto.Resume, error) {
	resume, err := s.repo.Get(resumeId)
	if err != nil {
		log.Error(err.Error())
		return dto.Resume{}, err
	}
	return MappingToDto(resume), err
}

func (s ResumeService) GetResumes(userId uuid.UUID, count int, offset int) ([]dto.Resume, error) {
	resumes, err := s.repo.GetWithPaging(userId, count, offset)
	if err != nil {
		log.Error(err.Error())
		return []dto.Resume{}, err
	}
	var dtos []dto.Resume
	for _, v := range resumes {
		dtos = append(dtos, MappingToDto(v))
	}

	return dtos, err
}

func (s ResumeService) CreateResume(resume dto.Resume) (uuid.UUID, error) {
	model := Mapping(resume)

	id, err := s.repo.Create(model)
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	return id, nil
}

func (s ResumeService) UpdateResume(resume dto.Resume) (uuid.UUID, error) {
	model := Mapping(resume)

	id, err := s.repo.Update(model)
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	return id, nil
}

func (s ResumeService) DeleteResume(resumeId uuid.UUID) error {
	err := s.repo.Delete(resumeId)

	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func NewService(repo repository.Resume) *ResumeService {
	return &ResumeService{repo: repo}
}

func Mapping(resume dto.Resume) profile.Resume {
	return profile.Resume{
		Id:                 resume.Id,
		UserId:             resume.UserId,
		ProgrammerLanguage: domain.ProgrammerLanguage(resume.ProgrammerLanguage),
		ProgrammerLevel:    domain.ProgrammerLevel(resume.ProgrammerLevel),
		ProgrammerType:     domain.ProgrammerType(resume.ProgrammerType),
		WorkExperiences:    MappingWorkExperience(resume.WorkExperiences),
	}
}

func MappingToDto(resume profile.Resume) dto.Resume {
	return dto.Resume{
		Id:                 resume.Id,
		UserId:             resume.UserId,
		ProgrammerLanguage: shared.ProgrammerLanguage(resume.ProgrammerLanguage),
		ProgrammerLevel:    shared.ProgrammerLevel(resume.ProgrammerLevel),
		ProgrammerType:     shared.ProgrammerType(resume.ProgrammerType),
		WorkExperiences:    MappingWorkExperienceToDto(resume.WorkExperiences),
	}
}

func MappingWorkExperience(experience []dto.WorkExperience) (result []profile.WorkExperience) {
	if len(experience) == 0 {
		return result
	}
	for _, w := range experience {
		unit := profile.WorkExperience{
			Id:          w.Id,
			ResumeId:    w.ResumeId,
			CompanyName: w.CompanyName,
			From:        w.From,
			To:          w.To,
			Content:     w.Content,
		}
		result = append(result, unit)
	}
	return result
}

func MappingWorkExperienceToDto(experience []profile.WorkExperience) (result []dto.WorkExperience) {
	if len(experience) == 0 {
		return result
	}
	for _, w := range experience {
		unit := dto.WorkExperience{
			Id:          w.Id,
			ResumeId:    w.ResumeId,
			CompanyName: w.CompanyName,
			From:        w.From,
			To:          w.To,
			Content:     w.Content,
		}
		result = append(result, unit)
	}
	return result
}
