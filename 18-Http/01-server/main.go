package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	go func() {
		server := http.Server{
			Addr:         ":8080",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		}
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Fatal(http.ListenAndServe(":8090", nil))


}
