package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	dbendpoints "goginkit/apiendpoints"

	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Tripleh1"
	dbname   = "postgres"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	r := mux.NewRouter()
	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	r.HandleFunc("/users", dbendpoints.AllUsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", dbendpoints.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/users", dbendpoints.InsertUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", dbendpoints.DeleteUserHandler).Methods(http.MethodDelete)
	r.HandleFunc("/users", dbendpoints.UpdateUserHandler).Methods(http.MethodPut)
	r.Use(mux.CORSMethodMiddleware(r))

	http.ListenAndServe(":8080", r)
}
