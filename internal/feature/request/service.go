package request

import (
	"context"
	"errors"
	"github.com/marki-eriker/kim-worker-go/internal/feature/user"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"math"
)

type IService interface {
	ListRequests(input *model.RequestListInput, user *user.User) ([]*Request, *model.PaginationOutput, error)
	GetRequest(requestID uint, user user.User) (*Request, error)
	UpdateRequestStatus(input model.RequestUpdateStatusInput) (*Request, error)
}

type Service struct {
	repository IRepository
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) ListRequests(input *model.RequestListInput, user *user.User) ([]*Request, *model.PaginationOutput, error) {
	var allowServiceTypes []uint

	switch input.ServiceID {
	case nil:
		allowServiceTypes = user.ServiceTypes
	default:
		for _, v := range user.ServiceTypes {
			if v == *input.ServiceID {
				allowServiceTypes = append(allowServiceTypes, *input.ServiceID)
			}
		}
	}

	if len(allowServiceTypes) == 0 {
		return nil, nil, errors.New("forbidden")
	}

	requests, count, err := s.repository.GetAllRequest(context.Background(), input, allowServiceTypes)
	if err != nil {
		return nil, nil, err
	}

	totalPages := math.Ceil(float64(count) / float64(*input.Filter.PageSize))

	pagination := model.PaginationOutput{
		TotalItems:      count,
		TotalPages:      int(totalPages),
		Page:            *input.Filter.Page,
		ItemsPerPage:    *input.Filter.PageSize,
		HasNextPage:     int(totalPages) > *input.Filter.Page,
		HasPreviousPage: *input.Filter.Page > 1,
	}

	return requests, &pagination, nil
}

func (s *Service) GetRequest(requestID uint, user user.User) (*Request, error) {
	request, err := s.repository.GetRequest(context.Background(), requestID)
	if err != nil {
		return nil, err
	}

	allow := true

	for _, v := range user.ServiceTypes {
		if v == request.ServiceTypeID {
			allow = true
			break
		}
	}

	if !allow {
		return nil, errors.New("unauthorized")
	}

	return request, nil
}

func (s *Service) UpdateRequestStatus(input model.RequestUpdateStatusInput) (*Request, error) {
	return s.repository.UpdateRequestStatus(context.Background(), input.RequestID, input.NewStatus)
}
