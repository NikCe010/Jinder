package grpc

import (
	jobs_service "Jinder/jinder-api/jobs/pkg/handler/grpc/protos"
	"Jinder/jinder-api/jobs/pkg/service"
	"Jinder/jinder-api/jobs/pkg/service/dto/vacancy"
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type VacancyService struct {
	Services *service.Service
	jobs_service.UnimplementedVacancyServiceServer
}

func (v VacancyService) GetVacancy(ctx context.Context, id *jobs_service.VacancyId) (*jobs_service.VacancyDto, error) {
	vacancyId, err := uuid.Parse(id.Content)
	if err != nil {
		return nil, err
	}

	vac, err := v.Services.GetVacancy(vacancyId)
	return vacancy.ToGrpc(vac), nil
}

func (v VacancyService) GetVacancies(ctx context.Context, request *jobs_service.GetVacanciesRequest) (*jobs_service.GetVacanciesResponse, error) {
	userId, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}

	vacancies, err := v.Services.GetVacancies(userId, request.Count, request.Offset)
	var response *jobs_service.GetVacanciesResponse
	for _, vac := range vacancies {
		response.Vacancies = append(response.Vacancies, vacancy.ToGrpc(vac))
	}
	return response, nil
}

func (v VacancyService) CreateVacancy(ctx context.Context, dto *jobs_service.VacancyDto) (*jobs_service.VacancyId, error) {
	vac, err := vacancy.GrpcToDto(dto)
	if err != nil {
		return nil, err
	}

	id, err := v.Services.CreateVacancy(vac)
	if err != nil {
		return nil, err
	}

	return &jobs_service.VacancyId{Content: id.String()}, nil
}

func (v VacancyService) UpdateVacancy(ctx context.Context, dto *jobs_service.VacancyDto) (*jobs_service.VacancyId, error) {
	vac, err := vacancy.GrpcToDto(dto)
	if err != nil {
		return nil, err
	}

	id, err := v.Services.UpdateVacancy(vac)
	if err != nil {
		return nil, err
	}

	return &jobs_service.VacancyId{Content: id.String()}, nil
}

func (v VacancyService) DeleteVacancy(ctx context.Context, id *jobs_service.VacancyId) (*emptypb.Empty, error) {
	vacancyId, err := uuid.Parse(id.Content)
	if err != nil {
		return nil, err
	}

	err = v.Services.DeleteVacancy(vacancyId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (v VacancyService) ViewVacancy(ctx context.Context, id *jobs_service.VacancyId) (*emptypb.Empty, error) {
	vacancyId, err := uuid.Parse(id.Content)
	if err != nil {
		return nil, err
	}

	err = v.Services.ViewVacancy(vacancyId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func NewVacancyService(services *service.Service) *VacancyService {
	return &VacancyService{Services: services}
}
