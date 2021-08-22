package user

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	tableName struct{} `pg:"users"`

	ID           uint
	Email        string
	FullName     string
	BaseRole     model.BaseRole
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ServiceTypes []uint `pg:",array"`
}

type CustomClaim struct {
	ID    string         `json:"id"`
	Roles model.BaseRole `json:"roles"`
	jwt.StandardClaims
}

func (u *User) HashPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(passwordHash)
	return nil
}

func (u *User) GenToken() (string, error) {
	expiredAt := time.Now().Add(time.Minute * 30)

	claims := CustomClaim{
		ID:    strconv.FormatInt(int64(u.ID), 10),
		Roles: u.BaseRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (u *User) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (u *User) GenRefreshToken() string {
	return uuid.New().String()
}
