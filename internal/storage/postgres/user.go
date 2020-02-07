package postgres

import (
	"database/sql"
	"errors"
	"go-hexagonal/internal/domain/user"

	"github.com/lib/pq"
	"github.com/valyala/fasthttp"
)

type userStorage struct {
	masterConn *sql.DB
}

func NewUserStorage(connections ...*sql.DB) (user.Servicer, error) {
	if err := DatabaseConnectionCheck(connections...); err != nil {
		return nil, err
	}
	return &userStorage{
		masterConn: connections[0],
	}, nil
}

func (s *userStorage) Register(u *user.Service) (statusCode int, err error) {
	uS := u.ToPostgres()
	_, err = s.masterConn.Exec("INSERT INTO users (name, password, email, customer_id, role) VALUES ($1,$2,$3,$4,$5)",
		uS.Name, uS.Password, uS.Email, nil, uS.Role)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "25P02" || pqErr.Code == "23505" {
				return fasthttp.StatusUnprocessableEntity, errors.New("name:duplicated")
			}
		}
		return fasthttp.StatusInternalServerError, err
	}
	return fasthttp.StatusOK, nil
}
