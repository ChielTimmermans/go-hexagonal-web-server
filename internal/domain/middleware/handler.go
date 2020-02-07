package middleware

import (
	"go-hexagonal/internal/domain/language"

	"github.com/valyala/fasthttp"
)

type Handler interface {
	SetLanguage(next fasthttp.RequestHandler, availableLang []language.Code, defaultLang language.Code) fasthttp.RequestHandler
}
