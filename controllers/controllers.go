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
	"strconv"
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


func (s *Server) Artist(w http.ResponseWriter, r*http.Request) {
	params := mux.Vars(r)
	UserId , err := strconv.ParseInt(params["UserId"], 10, 64)
	if err !=nil {
		log.Println("invalid user id")
		middleware.ErrorResponse(w, "Invalid data")
		return
	}

	if r.Method== "POST"{
		var data Entities.Artist
		var err error
		if err = json.NewDecoder(r.Body).Decode(&data); err !=nil{
			log.Println("Unable to read json because ", err)
			middleware.ErrorResponse(w, "Invalid data")
			return
		}
		log.Println(data)

		UserId, err = Models.CreateArtist(s.Db, data)
		if err != nil{
			middleware.ErrorResponse(w, "error inserting artist")
			return
		}
		middleware.OkResponse(w, http.StatusOK, data)
	}
	if r.Method== "GET" {
		log.Println("Getting Information for one artist", UserId)
		artist, err := Models.FetchArtist(s.Db, int(UserId))
		if err !=nil {
			middleware.JsonResponse(w, http.StatusBadRequest, "an error occured")
			return
		}
		middleware.OkResponse (w, http.StatusOK, artist)
	}
	return
}

func (s *Server) Album (w http.ResponseWriter, r*http.Request){
	if r.Method== "POST"{
		var data Entities.Album
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil{
			log.Println("unable to read json because ", err)
			middleware.ErrorResponse(w, "invalid data")
			return
		}
		AlbumId, err := Models.CreateAlbum(s.Db, data)
		if err !=nil{
			middleware.ErrorResponse(w, "error inserting album")
		}
		data.AlbumId = int(AlbumId)
		middleware.OkResponse(w, http.StatusOK, data)
	}

	if r.Method== "GET" {
		params := mux.Vars(r)
		AlbumId , err := strconv.ParseInt(params["AlbumId"], 10, 64)
		album, err := Models.FetchAlbum(s.Db, int(AlbumId))
		if err !=nil{
			middleware.ErrorResponse(w, "an error occurred")
		}
		middleware.OkResponse(w,http.StatusOK, album)
	}
	return
}

func (s *Server) Category (w http.ResponseWriter, r*http.Request) {
	if r.Method == "POST"{
		var data Entities.Category
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil{
			log.Println("unable to read data because ", err)
			middleware.ErrorResponse(w, "invalid data")
			return
		}
		CategoryId, err := Models.CreateCategory(s.Db, data)
		if err != nil{
			middleware.ErrorResponse(w, "errpr creating category")
		}
		data.CategoryId = int(CategoryId)
		middleware.OkResponse(w, http.StatusOK, data)
	}

	if r.Method == "GET"{
		params :=mux.Vars(r)
		CategoryId, err := strconv.ParseInt(params["CategoryId"], 10, 64)
		category, err := Models.FetchCategory(s.Db, int(CategoryId))
		if err !=nil {
			middleware.ErrorResponse(w, "an error occurred")
		}
		middleware.OkResponse(w, http.StatusOK, category)
	}
	return
}

func (s *Server) Songs (w http.ResponseWriter, r*http.Request) {
	if r.Method == "POST"{
		var data Entities.Songs
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil{
			log.Println("unable to read data because ", err)
			middleware.ErrorResponse(w, "invalid data")
			return
		}
		SongId, err := Models.CreateSong(s.Db, data)
		if err != nil{
			middleware.ErrorResponse(w, "errpr creating category")
		}
		data.Song_Id = int(SongId)
		middleware.OkResponse(w, http.StatusOK, data)
	}

	if r.Method == "GET"{
		params :=mux.Vars(r)
		SongId, err := strconv.ParseInt(params["SongId"], 10, 64)
		song, err := Models.FetchCategory(s.Db, int(SongId))
		if err !=nil {
			middleware.ErrorResponse(w, "an error occurred")
		}
		middleware.OkResponse(w, http.StatusOK, song)
	}
	return
}