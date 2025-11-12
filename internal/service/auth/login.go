package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/avraam311/warehouse-control/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrWrongPassword = errors.New("wrong password")
)

func (s *Service) Login(ctx context.Context, usr *models.UserDTO) (string, error) {
	user, err := s.repo.GetUser(ctx, usr.Email)
	if err != nil {
		return "", fmt.Errorf("service/login.go - %w", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(usr.Password)); err != nil {
		return "", ErrWrongPassword
	}

	role := user.Role
	claims := &models.Claims{
		Role:   role,
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   usr.Email,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.cfg.GetString("JWT_SECRET"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
