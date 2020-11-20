package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"music/controllers"
	"net/http"

	//"net/http"
)
func main () {
	port := ":5000"
	router := mux.NewRouter()
	fmt.Println("Hello")

	router.HandleFunc("/welcome", controllers.Welcome).Methods("GET")
	router.HandleFunc("/Artist", controllers.Artist).Methods("POST", "GET")

	err:= http.ListenAndServe(port, router)
	if err !=nil{
		log.Fatal("unable to start api because ", err)
	}
}

