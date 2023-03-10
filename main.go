package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	dbendpoints "goginkit/handlers"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Listening on localhost:8080")
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env")
	}
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
	})
	r := mux.NewRouter()
	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	r.HandleFunc("/users", dbendpoints.AllUsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", dbendpoints.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/users", dbendpoints.InsertUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", dbendpoints.DeleteUserHandler).Methods(http.MethodDelete)
	r.HandleFunc("/users", dbendpoints.UpdateUserHandler).Methods(http.MethodPut)
	r.Use(mux.CORSMethodMiddleware(r))

	http.ListenAndServe(":8080", c.Handler(r))
}
