package main

import (
	"go-hexagonal/internal/domain/language"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	cors "github.com/AdhityaRamadhanus/fasthttpcors"
)

func initRouter(s *fasthttp.Server, h *Handler, config *ConfigRouter,
	configCORS *ConfigCORS, configLanguage *ConfigLanguage) *router.Router {
	log.Println("Init router")

	r := router.New()

	r.RedirectTrailingSlash = config.RedirectTrailingSlash
	r.RedirectFixedPath = config.RedirectFixedPath
	r.HandleMethodNotAllowed = config.HandleMethodNotAllowed
	r.HandleOPTIONS = config.HandleOPTIONS

	CORS := initCors(configCORS)

	// Set servers router handler to the newly made router
	codes := make([]language.Code, len(configLanguage.AvailableLanguages))
	for k, v := range configLanguage.AvailableLanguages {
		codes[k] = language.Code(v)
	}

	s.Handler = CORS.CorsMiddleware(h.middleware.SetLanguage(r.Handler, codes, language.Code(configLanguage.DefaultLanguage)))

	initRoutes(r, h)
	log.Println("Init router done")

	return r
}

func initRoutes(r *router.Router, h *Handler) {
	r.POST("/register", h.user.Register)
	r.GET("/test", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	})
}

func initCors(config *ConfigCORS) (corsHandler *cors.CorsHandler) {
	log.Println("Init cors")
	// WithCors build cors
	corsHandler = cors.NewCorsHandler(cors.Options{
		// if you leave allowedOrigins empty then fasthttpcors will treat it as "*"
		AllowedOrigins: config.AllowedOrigins, // Only allow example.com to access the resource

		AllowCredentials: config.AllowCredentials,

		// if you leave allowedHeaders empty then fasthttpcors will accept any non-simple headers
		AllowedHeaders: config.AllowedHeaders,

		// if you leave this empty, only simple method will be accepted
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}, // only allow get or post to resource
		AllowMaxAge:    config.AllowMaxAge,                                  // cache the preflight result
		Debug:          config.Debug,                                        // turn on when strange cors behavior
	})
	log.Println("Init cors done")
	return
}
