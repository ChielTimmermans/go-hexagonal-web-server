package main

import (
	"go-hexagonal/internal/domain/middleware"
	"go-hexagonal/internal/domain/user"
	"go-hexagonal/internal/handler/json"

	jsoniter "github.com/json-iterator/go"
)

type Handler struct {
	user       user.Handler
	middleware middleware.Handler
}

func initHandler(s *Service) (h *Handler, err error) {
	jsonIteratorPool, jsonStreamPool := initJSON()
	errors := json.InitErrors(jsonStreamPool)
	h = &Handler{}

	if h.user, err = json.NewUser(s.user, jsonIteratorPool, jsonStreamPool, errors); err != nil {
		return nil, err
	}
	if h.middleware, err = json.NewMiddleware(errors); err != nil {
		return nil, err
	}
	return
}

func initJSON() (jsonIteratorPool jsoniter.IteratorPool, jsonStreamPool jsoniter.StreamPool) {
	jsonIteratorPool = jsoniter.NewIterator(jsoniter.ConfigFastest).Pool()

	// TODO not sure about the 100 buffer size, will lookup in the future
	jsonStreamPool = jsoniter.NewStream(jsoniter.ConfigFastest, nil, 100).Pool()
	return jsonIteratorPool, jsonStreamPool
}
