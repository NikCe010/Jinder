package token_manager

import (
	"Jinder/jinder-api/pkg/repository"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type TokenManagerService struct {
	repo repository.User
}

func NewTokenManagerService(repo repository.User) *TokenManagerService {
	return &TokenManagerService{repo: repo}
}

const (
	signingKey = "superSecretPassword"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId uuid.UUID `json:"user_id"`
}

func (s TokenManagerService) Generate(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s TokenManagerService) Validate(token string) (uuid.UUID, error) {
	result, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	claims, ok := result.Claims.(*tokenClaims)
	if !ok {
		return uuid.UUID{}, errors.New("result claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
