package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HTTPHandler handles the HTTP request
func HTTPHandler(s Service, router *mux.Router) {
	handlers := getHandlers(s)
	router.Handle(`/users`, handlers.addUserHandler).Methods(http.MethodPost)
	router.Handle(`/users/{user_id}`, handlers.getUserHander).Methods(http.MethodGet)
	router.Handle(`/users`, handlers.getUserHander).Methods(http.MethodGet)
}
