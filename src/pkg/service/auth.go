package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mauiwaui024/todo-golang"
	"github.com/mauiwaui024/todo-golang/pkg/repository"
)

const (
	salt       = "sjlkdhfnkjsdfhsjkd1233j4k;234k"
	tokenTTL   = 12 * time.Hour
	signingKey = "sjdnfjksdnfkj23494nfJJFQe"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

// 4.cоздали auth.go в котором реализовываем интерфейс, структруа которая в конструкторе принимает репозиторий с базой
type AuthService struct {
	//8.исправили на интерфейс авторизации из репозитория
	// repo repository.Repository
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

// 5. имплементируем метод CreateUser в котором будем передавать нашу структуру юзера еще на слой ниже - в репозиторий
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	//хэшанули пассворд
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

// /аутентификация
func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	//для генерации токена нам надо получить юзера из базы

	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accesToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		//анонимная функция возвращаяю ключ подпист или ошибку
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type")
	}
	return claims.UserId, nil
}
