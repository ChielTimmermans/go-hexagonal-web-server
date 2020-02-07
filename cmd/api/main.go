package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.Println("STARTING hexagonal design")
	log.SetFlags(log.Llongfile)  
	var err error
	config := initConfig()
	server := initServer(config.Server.Name)

	var storage *Storage
	if storage, err = initStorage(config.DBPostgres, "postgres"); err != nil {
		log.Fatal(err)
	}
	var service *Service
	if service, err = initService(storage); err != nil {
		log.Fatal(err)
	}
	var handler *Handler
	if handler, err = initHandler(service); err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())

	initRouter(server, handler, config.Router, config.CORS, config.Language)

	log.Printf("Starting up go-hexagonal back-end, listening on port: %d\n", config.Server.Port)
	log.Fatal(server.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port)))
}
