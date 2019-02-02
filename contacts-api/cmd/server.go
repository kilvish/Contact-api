package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/anil/contacts-api/internal/users"
	"github.com/anil/contacts-api/resources"
	"github.com/gorilla/mux"
)

func main() {
	store, err := resources.Init("mysql")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	userService := users.NewUserService(store)
	router := mux.NewRouter()
	users.HTTPHandler(userService, router)
	fmt.Println("Starting service on port 8080")
	if err := http.ListenAndServe(":8082", router); err != nil {
		fmt.Println("Error while running service")
	}
}
