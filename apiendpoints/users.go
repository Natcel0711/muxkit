package dbendpoints

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Tripleh1"
	dbname   = "postgres"
)

func test() {
	fmt.Println("Dickhead")
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte(`{"alive": true}`))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	idReq := mux.Vars(r)["id"]
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		panic("Something happened while accessing database")
	}
	//Convert to int
	id, err := strconv.Atoi(idReq)
	if err != nil {
		panic("Failed to convert ID to integer")
	}
	user := Users{Id: id}
	result := db.Find(&user)
	if result.Error != nil {
		panic("Failed to look for user")
	}
	jsonStr, err := json.Marshal(user)
	if err != nil {
		panic("Error while converting to Json")
	}
	w.Write([]byte(jsonStr))
}
func InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	var usuario Users
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		panic("Something happened while accessing database")
	}
	res := db.Create(&usuario)
	if res.Error != nil {
		panic("Error while creating user")
	}
	fmt.Println(res.RowsAffected, usuario.Id)
}
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var usuario Users
	var userFound Users
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		panic("Something happened while accessing database")
	}
	db.First(&userFound, usuario.Id)
	userFound.Name = usuario.Name
	db.Save(&userFound)
	w.Write([]byte(fmt.Sprintf("{\"Success\":true, \"Message\": \"User updated %d\"}", userFound.Id)))
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
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	idReq := mux.Vars(r)["id"]
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		panic("Something happened while accessing database")
	}
	//Convert to int
	id, err := strconv.Atoi(idReq)
	if err != nil {
		panic("Failed to convert ID to integer")
	}
	user := Users{Id: id}
	result := db.Find(&user)
	fmt.Println(result.Error, result.RowsAffected)
	if result.Error != nil {
		panic("Failed to look for user")
	}
	if result.RowsAffected == 0 {
		w.Write([]byte(fmt.Sprintf("{\"Success\":false, \"Message\": \"No user with id of %d\"}", id)))
		return
	}
	//db.Delete(&user)
	w.Write([]byte(fmt.Sprintf("{\"Success\":true, \"Message\": \"Deleted user with id of %d\"}", id)))
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
type Users struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Employees struct {
	EmployeeList []Employee `json:"EmployeeList"`
}
