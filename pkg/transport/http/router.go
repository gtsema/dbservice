package http

import (
	"dbservice/pkg/transport/http/handler"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/user", handler.UserGetHandler).Methods("GET")
	router.HandleFunc("/user", handler.UserPostHandler).Methods("POST")
	router.HandleFunc("/user", handler.UserDeleteHandler).Methods("DELETE")

	return router
}
