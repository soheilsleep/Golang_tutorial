package Handlers

import (
	"encoding/json"
	"fmt"
	"http-request/models"
	"net/http"
)

type UsersHandler struct{}

func (u *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && len(r.URL.Query().Get("id")) > 0:
		GetUser(w, r)
		return
	case r.Method == http.MethodGet:
		GetUsers(w, r)
		return
	case r.Method == http.MethodPost:
		CreateUser(w, r)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-Api-Key")
	for k, v := range r.Header {
		fmt.Println(k, v)
	}
	if apiKey != "123456789" {
		w.WriteHeader(401)
		fmt.Fprintln(w, "Invalid API Key")
		return
	}
	var user *models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Decoding Error %v", err)
		return
	}
	//add to db
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created")

}
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	//get form db
	fmt.Fprintf(w, "got user by id: %s\n", id)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	//get list form db
	fmt.Fprintf(w, "got users")
}
