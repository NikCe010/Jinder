package grpc

import (
	jobs_service "Jinder/jinder-api/jobs/pkg/handler/grpc/protos"
	"Jinder/jinder-api/jobs/pkg/service"
	"Jinder/jinder-api/jobs/pkg/service/dto/resume"
	"context"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ResumeService struct {
	Services *service.Service
	jobs_service.UnimplementedResumeServiceServer
}

func NewResumeService(services *service.Service) *ResumeService {
	return &ResumeService{Services: services}
}

func (r ResumeService) GetResume(ctx context.Context, id *jobs_service.ResumeId) (*jobs_service.ResumeDto, error) {
	uuid, err := uuid.Parse(id.Content)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	res, err := r.Services.GetResume(uuid)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	dto, err := resume.ToGrpc(res)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return dto, nil
}

func (r ResumeService) GetResumes(ctx context.Context, request *jobs_service.GetResumesRequest) (*jobs_service.GetResumesResponse, error) {
	uuid, err := uuid.Parse(request.UserId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resumes, err := r.Services.GetResumes(uuid, request.Count, request.Offset)

	var response jobs_service.GetResumesResponse
	for _, r := range resumes {
		dto, err := resume.ToGrpc(r)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		response.Resumes = append(response.Resumes, dto)
	}
	return &response, nil
}

func (r ResumeService) CreateResume(ctx context.Context, dto *jobs_service.ResumeDto) (*jobs_service.ResumeId, error) {
	res, err := resume.GrpcToDto(dto)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	id, err := r.Services.CreateResume(res)
	if err != nil {
		return nil, err
	}
	return &jobs_service.ResumeId{
		Content: id.String(),
	}, nil
}

func (r ResumeService) UpdateResume(ctx context.Context, dto *jobs_service.ResumeDto) (*jobs_service.ResumeId, error) {
	res, err := resume.GrpcToDto(dto)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	id, err := r.Services.UpdateResume(res)
	if err != nil {
		return nil, err
	}
	return &jobs_service.ResumeId{
		Content: id.String(),
	}, nil
}

func (r ResumeService) DeleteResume(ctx context.Context, id *jobs_service.ResumeId) (*emptypb.Empty, error) {
	resumeId, err := uuid.Parse(id.Content)
	if err != nil {
		return nil, err
	}

	err = r.Services.DeleteResume(resumeId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (r ResumeService) ViewResume(ctx context.Context, id *jobs_service.ResumeId) (*emptypb.Empty, error) {
	resumeId, err := uuid.Parse(id.Content)
	if err != nil {
		return nil, err
	}

	err = r.Services.ViewResume(resumeId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
