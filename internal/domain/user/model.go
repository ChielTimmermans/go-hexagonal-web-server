package user

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gopkg.in/guregu/null.v3"
)

type JSON struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	CustomerID int       `json:"customer_id"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
}

type Service struct {
	ID         int
	Name       string
	Password   string
	Email      string
	CustomerID int
	Role       string
	CreatedAt  time.Time
}

type Postgres struct {
	ID         int
	Name       string
	Password   string
	Email      string
	CustomerID null.Int
	Role       string
	CreatedAt  time.Time
}

func (j *JSON) ToService() *Service {
	return &Service{
		ID:         j.ID,
		Name:       j.Name,
		Password:   j.Password,
		Email:      j.Email,
		CustomerID: j.CustomerID,
		Role:       j.Role,
		CreatedAt:  j.CreatedAt,
	}
}

func (s *Service) ToJSON() *JSON {
	return &JSON{
		ID:         s.ID,
		Name:       s.Name,
		Password:   s.Password,
		Email:      s.Email,
		CustomerID: s.CustomerID,
		Role:       s.Role,
		CreatedAt:  s.CreatedAt,
	}
}

func (p *Postgres) ToService() *Service {
	return &Service{
		ID:         p.ID,
		Name:       p.Name,
		Password:   p.Password,
		Email:      p.Email,
		CustomerID: int(p.CustomerID.ValueOrZero()),
		Role:       p.Role,
		CreatedAt:  p.CreatedAt,
	}
}

func (s *Service) ToPostgres() *Postgres {
	return &Postgres{
		ID:         s.ID,
		Name:       s.Name,
		Password:   s.Password,
		Email:      s.Email,
		CustomerID: null.NewInt(int64(s.CustomerID), true),
		Role:       s.Role,
		CreatedAt:  s.CreatedAt,
	}
}

const (
	ADMIN         = "admin"
	CUSTOMERADMIN = "customeradmin"
	USER          = "user"
)

func (s *Service) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Name,
			validation.Required.Error("name_required"),
			validation.Length(3, 255).Error("name_length")),

		validation.Field(&s.Email,
			validation.Required.Error("email_required"),
			validation.Length(3, 255).Error("email_length"),
			is.Email.Error("email_wrong_format")),

		validation.Field(&s.Role,
			validation.Required.Error("role_required"),
			validation.In(ADMIN, CUSTOMERADMIN, USER).Error("role_not_found")),

		validation.Field(&s.Password,
			validation.Required.Error("password_required"),
			validation.Length(3, 255).Error("password_length")),
	)
}
