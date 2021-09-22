package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

const dsn = "host=localhost user=postgres password=postgres dbname=employee_mgmnt port=5432 sslmode=disable TimeZone=Asia/Kolkata"

func InitializeMigration() {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("db errror" + err.Error())
	}
	db.AutoMigrate(&User{})
}

func initializeRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	http.ListenAndServe(":9000", router)
}
func main() {
	InitializeMigration()
	initializeRouter()

	fmt.Println("Hello, playground")
}
