package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AbdulWasay1207/notes-sharing-app/controllers"
	middleware "github.com/AbdulWasay1207/notes-sharing-app/middlewares"
)

type config struct {
	addr string
}

type application struct {
	config config
}

func (api *application) getRouter() http.Handler {
	router := http.NewServeMux()

	router.Handle("POST /v1/notes", middleware.JwtAuth(http.HandlerFunc(controllers.CreateNewNote)))
	router.HandleFunc("POST /v1/signup", controllers.RegisterUser)
	router.HandleFunc("POST /v1/login", controllers.LoginUser)
	router.Handle("GET /v1/notes/{id}/{pass}", middleware.JwtAuth(http.HandlerFunc(controllers.GetNote)))
	// router.HandleFunc("GET /v1/notes", controllers.GetAllNote)
	router.Handle("DELETE /v1/notes/{id}", middleware.JwtAuth(http.HandlerFunc(controllers.Delete)))

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
