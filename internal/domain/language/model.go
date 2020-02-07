package language

import (
	"fmt"
	"go-hexagonal/internal"
)

type JSON struct {
	ENG string `json:"eng"`
	NLD string `json:"nld"`
}

type Service struct {
	ENG string
	NLD string
}

type Code string

// List of Language
const (
	NLD Code = "NLD"
	ENG Code = "ENG"
)

func (l *Service) ToJSON(lang Code) string {
	switch lang {
	case NLD:
		return l.NLD
	case ENG:
		return l.ENG
	default:
		return l.ENG
	}
}

func (j *JSON) ToService() *Service {
	return &Service{
		NLD: j.NLD,
		ENG: j.ENG,
	}
}

func (l *Service) ToPostgres() string {
	return fmt.Sprintf("{\"ENG\":\"%s\", \"NLD\": \"%s\"}", l.ENG, l.NLD)
}

func (l *Service) ToPostgresConcat() string {
	return internal.Concat("{\"ENG\":\"", l.ENG, "\", \"NLD\": \"", l.NLD, "\"}")
}

func (c *Code) ToPostgres() string {
	return internal.Concat("$.", string(*c))
}
