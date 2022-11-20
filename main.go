package main

import (
	"db-connection/handlers"
	"net/http"
)

func main() {

	userHandler := handlers.NewUserHandler()
	sm := http.NewServeMux()
	sm.Handle("/users", userHandler)
	http.ListenAndServe(":8080", sm)
}
