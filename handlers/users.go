package dbendpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	idReq := mux.Vars(r)["id"]
	psqlconn := os.Getenv("credentials")
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		w.Write([]byte("Something happened while accessing database"))
		return
	}
	//Convert to int
	id, err := strconv.Atoi(idReq)
	if err != nil {
		w.Write([]byte("Failed to convert ID to integer"))
		return
	}
	user := Users{Id: id}
	result := db.Find(&user)
	if result.Error != nil {
		w.Write([]byte("Failed to look for user"))
		return
	}
	jsonStr, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte("Error while converting to Json"))
		return
	}
	w.Write([]byte(jsonStr))
}
func InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	var usuario Users
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		w.Write([]byte("Error decoding user"))
		return
	}
	psqlconn := os.Getenv("credentials")
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		w.Write([]byte("Error connecting to DB"))
		return
	}
	res := db.Create(&usuario)
	if res.Error != nil {
		w.Write([]byte("Error creating user"))
		return
	}
	w.Write([]byte(fmt.Sprintf("{\"Success\":true, \"NewID\":%d, \"Message\": \"User %s Added\"}", usuario.Id, usuario.Name)))
}
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var usuario Users
	var userFound Users
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	psqlconn := os.Getenv("credentials")
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		w.Write([]byte("Error connecting to DB"))
		return
	}
	db.First(&userFound, usuario.Id)
	userFound.Name = usuario.Name
	userFound.Email = usuario.Email
	userFound.Password = usuario.Password
	userFound.UpdatedAt = time.Now()
	db.Save(&userFound)
	w.Write([]byte(fmt.Sprintf("{\"Success\":true, \"Message\": \"User updated %d\"}", userFound.Id)))
}

func AllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	psqlconn := os.Getenv("credentials")
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		w.Write([]byte("Error connecting to DB"))
		return
	}
	db.AutoMigrate(&Users{})
	allUsers := []Users{}
	result := db.Find(&allUsers)
	if result.Error != nil {
		w.Write([]byte("Error getting users"))
		return
	}
	jsonStr, err := json.Marshal(allUsers)
	if err != nil {
		w.Write([]byte("Error while converting to Json"))
		return
	}
	w.Write([]byte(jsonStr))
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	idReq := mux.Vars(r)["id"]
	psqlconn := os.Getenv("credentials")
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		w.Write([]byte("Something happened while accessing database"))
		return
	}
	//Convert to int
	id, err := strconv.Atoi(idReq)
	if err != nil {
		w.Write([]byte("Failed to convert ID to integer"))
		return
	}
	user := Users{Id: id}
	result := db.Find(&user)
	fmt.Println(result.Error, result.RowsAffected)
	if result.Error != nil {
		w.Write([]byte("Failed to look for user"))
		return
	}
	if result.RowsAffected == 0 {
		w.Write([]byte(fmt.Sprintf("{\"Success\":false, \"Message\": \"No user with id of %d\"}", id)))
		return
	}
	//db.Delete(&user)
	w.Write([]byte(fmt.Sprintf("{\"Success\":true, \"Message\": \"Deleted user with id of %d\"}", id)))
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
	Id        int            `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"createdat"`
	UpdatedAt time.Time      `json:"updatedat"`
	DeletedAt gorm.DeletedAt `json:"deletedat"`
}
type Employees struct {
	EmployeeList []Employee `json:"EmployeeList"`
}
