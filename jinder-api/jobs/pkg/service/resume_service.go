package service

import (
	"Jinder/jinder-api/jobs/pkg/infrastructure"
	"Jinder/jinder-api/jobs/pkg/infrastructure/event"
	"Jinder/jinder-api/jobs/pkg/repository"
	"Jinder/jinder-api/jobs/pkg/service/dto/resume"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type ResumeService struct {
	repo repository.Resume
	infr infrastructure.Recommendation
}

func (s ResumeService) ViewResume(resumeId uuid.UUID) error {
	r, err := s.repo.Get(resumeId)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	err = s.infr.NotifyWhenResumeViewed(event.NewResumeViewed(r.UserId, r.Id))
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (s ResumeService) GetResume(resumeId uuid.UUID) (resume.Resume, error) {
	r, err := s.repo.Get(resumeId)
	if err != nil {
		log.Error(err.Error())
		return resume.Resume{}, err
	}
	return resume.ToDto(r), err
}

func (s ResumeService) GetResumes(userId uuid.UUID, count, offset string) ([]resume.Resume, error) {
	resumes, err := s.repo.GetWithPaging(userId, count, offset)
	if err != nil {
		log.Error(err.Error())
		return []resume.Resume{}, err
	}
	var dtos []resume.Resume
	for _, r := range resumes {
		dtos = append(dtos, resume.ToDto(r))
	}

	return dtos, err
}

func (s ResumeService) CreateResume(dto resume.Resume) (uuid.UUID, error) {
	model := resume.ToDomain(dto)

	id, err := s.repo.Create(model)
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	err = s.infr.NotifyWhenResumeAdded(event.NewResumeCreated(dto.UserId, id))
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}
	return id, nil
}

func (s ResumeService) UpdateResume(dto resume.Resume) (uuid.UUID, error) {
	model := resume.ToDomain(dto)

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

func NewResumeService(repo repository.Resume, infr infrastructure.Recommendation) *ResumeService {
	return &ResumeService{repo: repo, infr: infr}
}
