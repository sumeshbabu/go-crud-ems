package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

//const dsn = "host=localhost user=postgres password=postgres dbname=employee_mgmnt port=5432 sslmode=disable TimeZone=Asia/Kolkata"

func InitializeMigration() {
	dbConnectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Kolkata",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))

	db, err = gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
	}
	db.AutoMigrate(&User{})
}

func initializeRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	http.ListenAndServe(":9000", router)
}
func main() {
	loadEnvVariable()
	InitializeMigration()
	initializeRouter()

	fmt.Println("Hello, playground")
}

func loadEnvVariable() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
