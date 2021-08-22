package refreshtoken

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type IService interface {
	CreateRefreshToken(agent string, userID uint, w *http.ResponseWriter) error
	Logout(token string, w *http.ResponseWriter) error
	GetTokenUserID(token string) (uint, error)
}

type Service struct {
	repository IRepository
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateRefreshToken(agent string, userID uint, w *http.ResponseWriter) error {
	ctx := context.Background()
	expires := time.Now().Add(time.Hour * 24 * 30)

	token := RefreshToken{
		Token:     uuid.New().String(),
		Agent:     agent,
		ExpiresIn: expires,
		UserId:    userID,
	}

	err := s.repository.Save(ctx, token)
	if err != nil {
		return err
	}

	tokens, _ := s.repository.GetByUserId(ctx, userID)

	if len(tokens) > 5 {
		tokensToDelete := tokens[5:]
		ids := make([]uint, 0, len(tokensToDelete))
		for _, v := range tokensToDelete {
			ids = append(ids, v.ID)
		}

		_ = s.repository.DeleteByIds(ctx, ids)
	}

	http.SetCookie(*w, &http.Cookie{
		Name:     "refresh-token",
		Value:    token.Token,
		Expires:  expires,
		HttpOnly: true,
	})

	return nil
}

func (s *Service) Logout(token string, w *http.ResponseWriter) error {
	http.SetCookie(*w, &http.Cookie{
		Name:    "refresh-token",
		Value:   "",
		Expires: time.Now(),
	})

	err := s.repository.DeleteByValue(context.Background(), token)
	return err
}

func (s *Service) GetTokenUserID(token string) (uint, error) {
	t, err := s.repository.GetByValue(context.Background(), token)
	if err != nil {
		return 0, err
	}
	return t.UserId, nil
}
