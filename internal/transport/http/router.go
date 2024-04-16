package http

import (
	"dbservice/internal/transport/http/handler"
	"github.com/gorilla/mux"
)

func NewRouter(handler *handler.Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/user", handler.UserGetHandler).Methods("GET")
	router.HandleFunc("/user", handler.UserPostHandler).Methods("POST")
	router.HandleFunc("/user", handler.UserDeleteHandler).Methods("DELETE")

	return router
}
