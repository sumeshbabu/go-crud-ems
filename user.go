package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, from getUsers")
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}
func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, from createUser")
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)

	json.NewEncoder(w).Encode(&user)
}
func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, from updateUser")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	json.NewEncoder(w).Encode(user)
}
func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, from getUser")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, from deleteUser")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	db.Delete(&user, params["id"])
	//json.NewEncoder(w).Encode(""")
}
