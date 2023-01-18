package main

import (
	"net/http"
	//"github.com/gin-gonic/gin"
	"encoding/json"
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

type employee struct{
	ID string `json:"id"`
	NAME string `json:"name"`
}

var db *gorm.DB

func initDB() {
    var err error
	dataSourceNAME := "root:Gd#2m@2001@tcp(localhost:3306)/"
	db, err = gorm.Open("mysql", dataSourceNAME)

	if err!= nil {
		fmt.Println(err)
        panic(err)
    }
	//db.Exec("CREATE DATABASE employees_db")
	db.Exec("USE employees_db")

	db.AutoMigrate(&employee{})
}

// var employees =[]employee{
// 	{ID:"1",NAME: "John"},
// 	{ID:"2",NAME: "Jane"},
// 	{ID:"3",NAME: "Joe"},
// 	{ID:"4",NAME: "Jack"},
// 	{ID:"5",NAME: "Jill"},
// 	{ID:"6",NAME: "Jacky"},
// }

func postEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee employee
	json.NewDecoder(r.Body).Decode(&newEmployee)
	db.Create(&newEmployee)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newEmployee)
}

func getEmployee(w http.ResponseWriter, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []employee
	db.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}

func main(){
	router:= mux.NewRouter()
	router.HandleFunc("/employees",postEmployee).Methods("POST")
	router.HandleFunc("/employees",getEmployee).Methods("GET")
	initDB()
    log.Fatal(http.ListenAndServe(":6000", router))
}