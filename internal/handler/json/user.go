package json

import (
	"errors"
	"go-hexagonal/internal/domain/user"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

type userHandler struct {
	service          user.Servicer
	jsonIteratorPool jsoniter.IteratorPool
	jsonStreamPool   jsoniter.StreamPool
	answer           *Answer
}

func NewUser(service user.Servicer, jip jsoniter.IteratorPool, jsp jsoniter.StreamPool, a *Answer) (user.Handler, error) {
	if service == nil {
		return nil, errors.New("userservice_nil")
	}
	return &userHandler{
		service:          service,
		jsonIteratorPool: jip,
		jsonStreamPool:   jsp,
		answer:           a,
	}, nil
}

func (h *userHandler) Register(ctx *fasthttp.RequestCtx) {
	var u user.JSON
	jsonIterator := h.jsonIteratorPool.BorrowIterator(ctx.PostBody())
	defer h.jsonIteratorPool.ReturnIterator(jsonIterator)
	jsonIterator.ReadVal(&u)

	if statusCode, err := h.service.Register(u.ToService()); err != nil {
		h.answer.SetAnswer(ctx, statusCode, err)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusNoContent)
}
