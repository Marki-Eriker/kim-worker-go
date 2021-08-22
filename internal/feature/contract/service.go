package contract

import (
	"context"
	"errors"
	"fmt"
	"github.com/marki-eriker/kim-worker-go/internal/feature/user"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"math"
	"time"
)

type IService interface {
	ListContracts(input *model.ContractListInput, user *user.User) ([]*Contract, *model.PaginationOutput, error)
	CreateOrUpdate(input *model.ContractCreateInput) (*Contract, error)
	Find(input *model.ContractFindInput) (*Contract, error)
}

type Service struct {
	repository IRepository
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) ListContracts(input *model.ContractListInput, user *user.User) ([]*Contract, *model.PaginationOutput, error) {
	var allowServiceTypes []uint

	switch input.ServiceTypeID {
	case nil:
		allowServiceTypes = user.ServiceTypes
	default:
		for _, v := range user.ServiceTypes {
			if v == *input.ServiceTypeID {
				allowServiceTypes = append(allowServiceTypes, *input.ServiceTypeID)
			}
		}
	}

	if len(allowServiceTypes) == 0 {
		return nil, nil, errors.New("forbidden")
	}

	contracts, count, err := s.repository.GetAllContract(context.Background(), input, allowServiceTypes)
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

	return contracts, &pagination, nil
}

func (s *Service) CreateOrUpdate(input *model.ContractCreateInput) (*Contract, error) {
	if contract, err := s.repository.FindByRequestID(context.Background(), input.RequestID); err == nil {
		contract.FileStorageItemID = input.FileID
		contract.Number = input.ContractNumber
		err = s.repository.UpdateFileID(context.Background(), contract)
		return contract, err
	}

	newContract := Contract{
		ServiceRequestID:  input.RequestID,
		Number:            input.ContractNumber,
		ContractorID:      input.ContractorID,
		FileStorageItemID: input.FileID,
		CreatedAt:         time.Now(),
	}

	err := s.repository.Save(context.Background(), &newContract)
	fmt.Printf("%+v", newContract)
	return &newContract, err
}

func (s *Service) Find(input *model.ContractFindInput) (*Contract, error) {
	return s.repository.FindById(context.Background(), input.ContractID)
}
