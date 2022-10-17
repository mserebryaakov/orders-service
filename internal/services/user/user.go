package user

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"orders-service/internal/domain/user"
	repository "orders-service/internal/domain/user"

	"github.com/dgrijalva/jwt-go"

	"time"
)

// Соль для хэширования пароля
// Время жизни токена
// Подпись токена
const (
	salt       = "foiuwm549sd432h"
	tokenTTL   = 12 * time.Hour
	signingKey = "tboobtwebpdfigk"
)

// объявим структуру для настройки генерации токена с UserId
type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

// Сервис взаимодействия с пользователями
type UserUseCase struct {
	repo repository.UserRepository
}

// Создание сервиса взаимодействия с пользователями
func NewUserService(ur repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		repo: ur,
	}
}

// Usecase регистрации пользователя
func (s *UserUseCase) CreateUser(ctx context.Context, user user.User) (string, error) {
	user.Password = generatePasswordHash(user.Password)
	id, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	return id, nil
}

// Usecase генерация токена (вход)
func (s *UserUseCase) GenerateToken(ctx context.Context, user user.UserSignUpDTO) (string, error) {
	user.Password = generatePasswordHash(user.Password)
	u, err := s.repo.GetUser(ctx, user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		u.ID,
	})

	return token.SignedString([]byte(signingKey))
}

// Usecase регистрации пользователя
func (s *UserUseCase) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

// Генерация хэша пароля
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
