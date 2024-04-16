package handler

import (
	"dbservice/internal/models/api"
	"dbservice/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandlerHttp struct {
	userService service.UserService
}

func NewUserHandlerHttp(userService service.UserService) *UserHandlerHttp {
	return &UserHandlerHttp{userService: userService}
}

func (h UserHandlerHttp) UserGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method: %s, Path: %s", r.Method, r.URL.Path)

	user, err := h.userService.GetUser(r.RequestURI)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h UserHandlerHttp) UserPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method: %s, Path: %s", r.Method, r.URL.Path)

	var u api.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := h.userService.CreateUser(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h UserHandlerHttp) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method: %s, Path: %s", r.Method, r.URL.Path)
}
