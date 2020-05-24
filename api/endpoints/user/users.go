package user

import (
	"github.com/gorilla/mux"
)

func SetUserEndpoint(router *mux.Router) {
	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/users/{name}", getUser).Methods("GET")
	router.HandleFunc("/api/users", createUser).Methods("POST")
	router.HandleFunc("/api/users/{name}", editUser).Methods("PUT")
	router.HandleFunc("/api/users/{name}", delUser).Methods("DELETE")
}
