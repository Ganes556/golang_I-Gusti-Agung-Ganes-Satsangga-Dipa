package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	router := &Router{mux, map[string]map[string]http.HandlerFunc{}}
	
	router.Get("/users", UserHandlers.Gets)
	router.Get("/users/:id", UserHandlers.Get)
	router.Post("/users", UserHandlers.Create)
	router.Put("/users/:id", UserHandlers.Update)
	router.Delete("/users/:id", UserHandlers.Delete)

	fmt.Println("server started at localhost:8000")
	router.Run(":8000")
}