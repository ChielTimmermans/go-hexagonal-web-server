package service

import (
	"errors"
	"go-hexagonal/internal/domain/user"

	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	storage user.Storager
}

func NewUser(storage user.Storager) (user.Servicer, error) {
	if storage == nil {
		return nil, errors.New("userstorage_nil")
	}
	return &userService{
		storage: storage,
	}, nil
}

func (s *userService) Register(u *user.Service) (int, error) {
	if err := u.Validate(); err != nil {
		return fasthttp.StatusUnprocessableEntity, err
	}
	var err error
	if u.Password, err = hashAndSalt(u.Password); err != nil {
		return fasthttp.StatusInternalServerError, err
	}
	if statusCode, err := s.storage.Register(u); err != nil {
		return statusCode, err
	}
	return fasthttp.StatusOK, nil
}

func hashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
