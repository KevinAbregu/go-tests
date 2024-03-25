package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// From WSL to local MySQL server on Windows Host
const DNS = "root:root@tcp(192.168.0.22:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	// Open MySQL connection
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	// Migrate Table of User if it doesnt exist.
	DB.AutoMigrate(&User{})

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Header response
	w.Header().Set("Content-Type", "application/json")
	var users []User
	// Find all users
	DB.Find(&users)
	// Response all users (json)
	json.NewEncoder(w).Encode(users)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get all request parameters
	params := mux.Vars(r)
	var user User
	// Find first coincidence of a user with the ID indicated
	DB.First(&user, params["id"])
	// Response User (json)
	json.NewEncoder(w).Encode(user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	// Read Request Body and save it on user variable
	json.NewDecoder(r.Body).Decode(&user)
	// Create user variable
	DB.Create(&user)
	// Response User Created
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	// Find first coincidence of a user with the ID indicated and save it on user variable
	DB.First(&user, params["id"])
	// Read Request Body and save it on user variable
	json.NewDecoder(r.Body).Decode(&user)
	// Save user on DB
	DB.Save(&user)
	// Response User Updated
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	// Delete the user indicated
	DB.Delete(&user, params["id"])
	// Response with a string
	json.NewEncoder(w).Encode("The user is deleted Successfully")

}
