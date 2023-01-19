package main

import (
	"fmt"
	. "github.com/Vectormike/go-with-mongo/staffrestapi/config"
	. "github.com/Vectormike/go-with-mongo/staffrestapi/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func HealthStatus(w http.ResponseWriter, r *http.Request) {
	serverName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Server Name: %s ", serverName)
	retResponse(w, http.StatusOK, map[string]string{"server": name, "result": "success"})
}

func init() {
	// Set up the database connection
	config := Config{}
	dao := StaffDAO{}

	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database

	dao.Connect()
}

func main() {
	http.HandleFunc("/health-status", HealthStatus)
	r := mux.NewRouter()
	// r.HandleFunc("/staff", AllStaff).Methods("GET")
	// r.HandleFunc("/staff", CreateStaff).Methods("POST")
	// r.HandleFunc("/staff/{id}", FindStaff).Methods("GET")
	// r.HandleFunc("/staff/{id}", UpdateStaff).Methods("PUT")
	// r.HandleFunc("/staff/{id}", DeleteStaff).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
