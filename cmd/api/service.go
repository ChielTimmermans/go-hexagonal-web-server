package main

import (
	"go-hexagonal/internal/domain/user"
	"go-hexagonal/internal/service"
)

type Service struct {
	user user.Servicer
}

func initService(storage *Storage) (s *Service, err error) {
	s = &Service{}

	if s.user, err = service.NewUser(storage.user); err != nil {
		return nil, err
	}
	return
}
