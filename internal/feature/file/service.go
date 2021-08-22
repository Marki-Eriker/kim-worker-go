package file

import (
	"context"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"time"
)

type IService interface {
	Create(input *model.FileCreateInput) (*File, error)
}

type Service struct {
	repository IRepository
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(input *model.FileCreateInput) (*File, error) {
	file := File{
		OriginalFilename: input.FileName,
		Extension:        input.Extension,
		MimeType:         input.MimeType,
		Size:             input.Size,
		Checksum:         input.Checksum,
		CreatedAt:        time.Now(),
	}

	if err := s.repository.Save(context.Background(), &file); err != nil {
		return nil, err
	}

	return &file, nil
}
