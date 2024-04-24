package handler

import (
	"dbservice/internal/service"
	"net/http"
)

type UserHandler interface {
	UserGetHandler(w http.ResponseWriter, r *http.Request)
	UserPostHandler(w http.ResponseWriter, r *http.Request)
	UserPutHandler(w http.ResponseWriter, r *http.Request)
	UserDeleteHandler(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	UserHandler
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		UserHandler: NewUserHandlerHttp(s),
	}
}
