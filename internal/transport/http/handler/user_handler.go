package handler

import (
	"dbservice/internal/models/api"
	"dbservice/internal/service"
	"encoding/json"
	"net/http"
)

type UserHandlerHttp struct {
	userService service.UserService
}

func NewUserHandlerHttp(userService service.UserService) *UserHandlerHttp {
	return &UserHandlerHttp{userService: userService}
}

func (h UserHandlerHttp) UserGetHandler(w http.ResponseWriter, r *http.Request) {
	user, err := h.userService.GetUser(r.URL.Query().Get("chatId"))
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h UserHandlerHttp) UserPostHandler(w http.ResponseWriter, r *http.Request) {
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

func (h UserHandlerHttp) UserPutHandler(w http.ResponseWriter, r *http.Request) {
	var u api.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := h.userService.UpdateUser(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h UserHandlerHttp) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.userService.DeleteUser(r.URL.Query().Get("chatId")); err != nil {
		if err.Error() == "user not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
