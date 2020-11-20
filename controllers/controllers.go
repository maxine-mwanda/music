package controllers

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"music/Entities"
	"music/Models"
	"music/middleware"
	"net/http"
)

func Welcome (writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Header.Get("user-agent"))
	log.Println("Hello")
	response := map[string]interface{}{
		"name": "Green",
		"dorm": "12",
	}
	middleware.JsonResponse(writer, 200, response)
}

func Artist (writer http.ResponseWriter, r*http.Request) {
	params := mux.Vars(r)
	UserId := params["UserId"]

	if r.Method== "POST"{
		var data Entities.Artist
		var err error
		if err = json.NewDecoder(r.Body).Decode(&data); err !=nil{
			log.Println("Unable to read json because ", err)
			middleware.ErrorResponse(writer, "Invalid data")
			return
		}
		log.Println(data)

		UserId, err = Models.CreateArtist(s.Db, data)
		if err != nil{
			middleware.ErrorResponse(writer, "error inserting artist")
			return
		}
		middleware.OkResponse(writer, http.StatusOK, data)
	}
	if r.Method== "GET" {
		log.Println("Getting Information for one artist", UserId)
		artist, err := Models.Fetchartist(s.DB, UserId)
		if err !=nil {
			middleware.JsonResponse(writer, http.StatusBadRequest, "an error occured")
			return
		}
		middleware.OkResponse (writer, http.StatusOK, artist)
	}
}