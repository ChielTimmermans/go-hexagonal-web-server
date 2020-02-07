package json

import (
	"go-hexagonal/internal/domain/language"
	"go-hexagonal/internal/domain/middleware"

	"github.com/valyala/fasthttp"
)

type middlewareHandler struct {
	answer *Answer
	// availableLanguages []string
	// defaultLanguage    string
}

func NewMiddleware(a *Answer) (middleware.Handler, error) {
	return &middlewareHandler{
		answer: a,
	}, nil
}

func (h *middlewareHandler) SetLanguage(next fasthttp.RequestHandler,
	availableLang []language.Code, defaultLang language.Code) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		templanguageCode := ctx.Request.Header.Peek("Accept-Language")
		foundLanguage := false
		for _, v := range availableLang {
			if string(templanguageCode) == string(v) {
				foundLanguage = true
				ctx.SetUserValueBytes([]byte("language"), v)
				ctx.Response.Header.Set("Content-Language", string(v))
			}
		}

		if !foundLanguage {
			ctx.SetUserValue("language", defaultLang)
			ctx.Response.Header.Set("Content-Language", string(defaultLang))
		}
		next(ctx)
	}
}
