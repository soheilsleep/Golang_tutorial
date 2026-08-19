package examples

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type TestHandler struct {
}

func CreateServerWithMux() {
	mux := http.NewServeMux()
	mux.Handle("/google", http.RedirectHandler("https://google.com", 307))
	mux.Handle("/yahoo", http.RedirectHandler("https://yahoo.com", 307))
	mux.HandleFunc("/get-user", GetUser)

	server1 := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      mux,
	}
	err := server1.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
func CreateServerWithCustomHandler() {
	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      TestHandler{},
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (t TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{'id': '%d', 'name': '%s'}", 1, "soheil")
}
