package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func Connecttodb() (connection *sql.DB, err error)  {
	dburi:="root:daddy@tcp(172.17.0.1:3306)/music"
	connection, err= sql.Open("mysql", dburi)
	if err != nil {
		fmt.Println("unable to connect to db")
		return
	}
	return
}

func InitLogger(){
	logFolder :=os.Getenv("LOG_FOLDER")
	log.Println(logFolder)
}

func DbConnect()(db *sql.DB){
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:3306)%s", dbUsername, dbPassword, dbHost,dbName)
	log.Println("connecting to db ", dbURI)
	
	db, err :=sql.Open("mysql", dbURI)
	if err!=nil{
		log.Println("unable to open db because ", err)
		os.Exit(3)
	}
	if err = db.Ping(); err!=nil{
		log.Println("Unable to confirm db connection because ", err)
	}
	
	db.SetMaxIdleConns(64)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(10 *time.Second)
	log.Println("Succesfully connected to DB")
	
	return 
}