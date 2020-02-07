package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

func initServer(name string) (server *fasthttp.Server) {
	log.Println("Init server")
	server = &fasthttp.Server{
		Name: name,
	}
	log.Println("Init server done")
	return
}
