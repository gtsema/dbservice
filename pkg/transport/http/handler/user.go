package handler

import (
	"dbservice/pkg/service"
	"fmt"
	"net/http"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get")
}

func UserPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")
}

func UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("del")
}
