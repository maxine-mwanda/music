package controllers

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"music/utils"
	"net/http"
	"os"
)

type Server struct {
	Db *sql.DB
	Router *mux.Router

}

func (s *Server) Init(){
	utils.InitLogger()
	s.Db= utils.DbConnect()
	//s.InitRoutes ()
}

func (s *Server) Run(){
	port := ":" + os.Getenv("PORT")
	log.Println("Listening on port ", port)

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
	})

	handler:= c.Handler(s.Router)
	log.Fatalln(http.ListenAndServe(port, handler))
}
