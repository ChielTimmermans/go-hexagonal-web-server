package user

import "github.com/valyala/fasthttp"

type Handler interface {
	Register(ctx *fasthttp.RequestCtx)
}
