package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"

	"encoding/json"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Tripleh1"
	dbname   = "postgres"
)

func main() {
	r := mux.NewRouter()

	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	r.HandleFunc("/foo", fooHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	r.HandleFunc("/users", AllUsersHandler).Methods(http.MethodGet)
	r.Use(mux.CORSMethodMiddleware(r))

	http.ListenAndServe(":8080", r)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte(`{"alive": true}`))
}

func AllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()
	rows, err := db.Query(`SELECT * FROM public.Users`)
	CheckError(err)
	defer rows.Close()
	var emps Employees
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		CheckError(err)
		emp := Employee{
			Id:   id,
			Name: name,
		}
		emps.AddEmployee(emp)
	}
	jsonStr, err := json.Marshal(emps)
	CheckError(err)
	w.Write([]byte(jsonStr))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
func (emps *Employees) AddEmployee(employee Employee) []Employee {
	emps.EmployeeList = append(emps.EmployeeList, employee)
	return emps.EmployeeList
}

type Employee struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Employees struct {
	EmployeeList []Employee `json:"EmployeeList"`
}
