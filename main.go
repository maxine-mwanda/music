package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"music/controllers"
	"net/http"
)


func main(){
	var s controllers.Server

port := ":5000"
router := mux.NewRouter()
fmt.Println("Hello")

	router.HandleFunc("/welcome", controllers.Welcome).Methods("GET")
	router.HandleFunc("/Artist", s.Artist).Methods("POST", "GET")
	router.HandleFunc("/Album", s.Album).Methods("POST", "GET")
	router.HandleFunc("/Category", s.Category).Methods("POST", "GET")



err:= http.ListenAndServe(port, router)
	if err !=nil{
		log.Fatal("unable to start api because ", err)
	}
}

var server = controllers.Server{}

func Run(){
	_ = godotenv.Load()
	server.Init()
	server.Run()
}

