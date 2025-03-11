package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AbdulWasay1207/notes-sharing-app/controllers"
)

type config struct {
	addr string
}

type application struct {
	config config
}

func (api *application) getRouter() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /v1/notes", controllers.CreateNewNote)
	router.HandleFunc("GET /v1/notes/{id}/{pass}", controllers.GetNote)
	router.HandleFunc("GET /v1/notes", controllers.GetAllNote)
	router.HandleFunc("DELETE /v1/notes/{id}/{pass}", controllers.Delete)

	return router
}

func (api *application) run(router http.Handler) error {
	srv := &http.Server{
		Addr:         api.config.addr,
		Handler:      router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Printf("Server has started at %s", api.config.addr)

	return srv.ListenAndServe()
}
