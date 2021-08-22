package user

import (
	"context"
	"errors"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"math"
	"time"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrTokenGenerate      = errors.New("error generating token")
	ErrNoFieldToUpdate    = errors.New("no field to update")
)

type IService interface {
	ListUsers(filter *model.PaginationInput) ([]*User, *model.PaginationOutput, error)
	CreateUser(input model.UserCreateInput) (*User, error)
	DeleteUser(id uint) error
	Login(input model.UserLoginInput) (string, uint, error)
	GetUser(id uint) (User, error)
	UpdateUserMe(input model.UserUpdateMeInput, user *User) (*User, error)
	UpdatePassword(input model.UserUpdatePasswordInput) error
	UpdateUser(input model.UserUpdateMainInput) (*User, error)
	UpdateUserServiceTypes(input model.UserGrantRequestAccessInput) (*User, error)
}

type Service struct {
	repository IRepository
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) ListUsers(filter *model.PaginationInput) ([]*User, *model.PaginationOutput, error) {
	users, count, err := s.repository.GetAll(context.Background(), filter)
	if err != nil {
		return nil, nil, err
	}

	totalPages := math.Ceil(float64(count) / float64(*filter.PageSize))

	pagination := model.PaginationOutput{
		TotalItems:      count,
		TotalPages:      int(totalPages),
		Page:            *filter.Page,
		ItemsPerPage:    *filter.PageSize,
		HasNextPage:     int(totalPages) > *filter.Page,
		HasPreviousPage: *filter.Page > 1,
	}

	return users, &pagination, nil
}

func (s *Service) GetUser(id uint) (User, error) {
	user, err := s.repository.GetById(context.Background(), id)
	return user, err
}

func (s *Service) CreateUser(input model.UserCreateInput) (*User, error) {
	var user User
	if err := user.HashPassword(input.Password); err != nil {
		return nil, err
	}

	user.BaseRole = input.BaseRole
	user.FullName = input.FullName
	user.Email = input.Email

	if _, err := s.repository.Save(context.Background(), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) DeleteUser(id uint) error {
	err := s.repository.DeleteByID(context.Background(), id)
	return err
}

func (s *Service) Login(input model.UserLoginInput) (string, uint, error) {
	user, err := s.repository.GetByEmail(context.Background(), input.Email)
	if err != nil {
		return "", 0, ErrInvalidCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return "", 0, ErrInvalidCredentials
	}

	token, err := user.GenToken()
	if err != nil {
		return "", 0, ErrTokenGenerate
	}

	return token, user.ID, nil
}

func (s *Service) UpdateUserMe(input model.UserUpdateMeInput, user *User) (*User, error) {

	didUpdate := false

	if input.Email != nil {
		if len(*input.Email) > 0 {
			user.Email = *input.Email
			didUpdate = true
		}
	}

	if input.FullName != nil {
		if len(*input.FullName) > 0 {
			user.FullName = *input.FullName
			didUpdate = true
		}
	}

	if !didUpdate {
		return nil, ErrNoFieldToUpdate
	}

	user.UpdatedAt = time.Now()

	err := s.repository.Update(context.Background(), user)
	return user, err
}

func (s *Service) UpdatePassword(input model.UserUpdatePasswordInput) error {
	user, err := s.repository.GetById(context.Background(), input.ID)
	if err != nil {
		return err
	}

	if err := user.HashPassword(input.Password); err != nil {
		return err
	}

	user.UpdatedAt = time.Now()

	err = s.repository.Update(context.Background(), &user)
	return err
}

func (s *Service) UpdateUser(input model.UserUpdateMainInput) (*User, error) {
	ctx := context.Background()
	user, err := s.repository.GetById(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	didUpdate := false

	if input.Email != nil {
		if len(*input.Email) > 0 {
			user.Email = *input.Email
			didUpdate = true
		}
	}

	if input.FullName != nil {
		if len(*input.FullName) > 0 {
			user.FullName = *input.FullName
			didUpdate = true
		}
	}

	if input.BaseRole != nil {
		user.BaseRole = *input.BaseRole
		didUpdate = true
	}

	if !didUpdate {
		return nil, ErrNoFieldToUpdate
	}

	user.UpdatedAt = time.Now()

	err = s.repository.Update(ctx, &user)
	return &user, err
}

func (s *Service) UpdateUserServiceTypes(input model.UserGrantRequestAccessInput) (*User, error) {
	ctx := context.Background()
	user, err := s.repository.GetById(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	user.ServiceTypes = input.ServiceTypes

	err = s.repository.Update(ctx, &user)
	return &user, err
}
