package management

import (
	"http-request/Handlers"
	"log"
	"net/http"
	"time"
)

func Run() {
	mux := http.NewServeMux()
	mux.Handle("/users/", &Handlers.UsersHandler{})

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
