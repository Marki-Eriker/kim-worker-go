package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/marki-eriker/kim-worker-go/internal/feature/user"
)

type ctxKey string

const (
	CurrentUserKey ctxKey = "currentUser"
	RefreshToken   ctxKey = "refreshToken"
	Cookie         ctxKey = "cookie"
	UserAgent      ctxKey = "userAgent"
)

func AuthMiddleware(repo user.IRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			rt, err := r.Cookie("refresh-token")
			if err == nil {
				ctx = context.WithValue(ctx, RefreshToken, rt)
			}

			ctx = context.WithValue(ctx, Cookie, &w)

			userAgent := r.Header.Get("User-Agent")
			ctx = context.WithValue(ctx, UserAgent, userAgent)

			t := r.Header.Get("X-JWT")

			if t == "" {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			token, err := parseToken(t)
			if err != nil {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			userID, err := strconv.Atoi(claims["id"].(string))
			if err != nil {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			currentUser, err := repo.GetById(ctx, uint(userID))
			if err != nil {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			ctx = context.WithValue(ctx, CurrentUserKey, currentUser)

			ctx = context.WithValue(ctx, RefreshToken, rt)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func parseToken(t string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		t := []byte(os.Getenv("JWT_SECRET"))
		return t, nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse token error: %v", err)
	}

	return jwtToken, nil
}

func GetCurrentUserFromCTX(ctx context.Context) (*user.User, error) {
	if ctx.Value(CurrentUserKey) == nil {
		return nil, errors.New("no user in context")
	}

	currentUser, ok := ctx.Value(CurrentUserKey).(user.User)

	if !ok {
		return nil, errors.New("no user in context")
	}

	return &currentUser, nil
}

func GetRefreshTokenFromCTX(ctx context.Context) (string, error) {
	if ctx.Value(RefreshToken) == nil {
		return "", errors.New("no refresh token in cookie")
	}

	token, ok := ctx.Value(RefreshToken).(*http.Cookie)
	if !ok {
		return "", errors.New("failed to cast cookie")
	}

	return token.Value, nil
}

func GetResponseWriterFromCTX(ctx context.Context) (*http.ResponseWriter, error) {
	if ctx.Value(Cookie) == nil {
		return nil, errors.New("no writer in ctx")
	}

	w, ok := ctx.Value(Cookie).(*http.ResponseWriter)
	if !ok {
		return nil, errors.New("failed to cast ResponseWriter")
	}

	return w, nil
}
