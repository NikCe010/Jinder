package work_experience

import (
	"Jinder/jinder-api/jobs/pkg/domain/profile"
	jobs_service "Jinder/jinder-api/jobs/pkg/handler/grpc/protos"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"time"
)

type WorkExperience struct {
	Id       uuid.UUID `json:"id"`
	ResumeId uuid.UUID `json:"resume_id"`

	CompanyName string    `json:"company_name"`
	From        time.Time `json:"from"`
	To          time.Time `json:"to"`
	Content     string    `json:"content"`
}

func ToDomain(experience []WorkExperience) (result []profile.WorkExperience) {
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

func ToDto(experience []profile.WorkExperience) (result []WorkExperience) {
	if len(experience) == 0 {
		return result
	}
	for _, w := range experience {
		unit := WorkExperience{
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

func ToGrpc(exp []WorkExperience) (result []*jobs_service.WorkExp, err error) {
	if len(exp) == 0 {
		return result, nil
	}
	for _, e := range exp {
		from, err := ptypes.TimestampProto(e.From)
		if err != nil {
			return nil, err
		}
		to, err := ptypes.TimestampProto(e.To)
		if err != nil {
			return nil, err
		}

		unit := &jobs_service.WorkExp{
			Id:          e.Id.String(),
			ResumeId:    e.ResumeId.String(),
			CompanyName: e.CompanyName,
			From:        from,
			To:          to,
			Content:     e.Content,
		}
		result = append(result, unit)
	}
	return result, nil
}

func GrpcToDto(exp []*jobs_service.WorkExp) (result []WorkExperience, err error) {
	if len(exp) == 0 {
		return result, nil
	}
	for _, e := range exp {
		id, err := uuid.Parse(e.Id)
		if err != nil {
			return nil, err
		}
		resumeId, err := uuid.Parse(e.ResumeId)
		if err != nil {
			return nil, err
		}

		from, err := ptypes.Timestamp(e.From)
		if err != nil {
			return nil, err
		}

		to, err := ptypes.Timestamp(e.To)
		if err != nil {
			return nil, err
		}

		unit := WorkExperience{
			Id:          id,
			ResumeId:    resumeId,
			CompanyName: e.CompanyName,
			From:        from,
			To:          to,
			Content:     e.Content,
		}
		result = append(result, unit)
	}
	return result, nil
}
