package service

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log/slog"
	"time"
)

type TokenClaims struct {
	jwt.StandardClaims
	Login string `json:"login"`
}

type _authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) service.AuthService {
	return _authService{repo: repo}
}

func (service _authService) Register(ctx context.Context, login,
	password string) (string, error) {

	hash := generatePassword(password)

	userName, err := service.repo.Register(ctx, login, hash)

	if err != nil {
		slog.Error(err.Error())
		return "", errors.New("не смогли создать пользователя")
	}

	return generateToken(userName)

}

func (service _authService) GenerateToken(ctx context.Context, login,
	password string) (string, error) {
	return "", nil
}

func generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(model.Salt)))
}

func generateToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(model.TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Login: login,
	})
	return token.SignedString([]byte(model.SignInKey))
}
