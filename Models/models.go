package Models

import (
	"database/sql"
	"fmt"
	"log"
	"music/Entities"
	"music/utils"
)

func CreateArtist(db *sql.DB, data Entities.Artist) (id int64, err error) {
	query:= "insert into artist (first_name, last_name, gender, password) values (?,?,?,?)"
	db, err = utils.Connecttodb()
	if err !=nil{
		log.Println("Unable to connect to db")
		return
	}
	row, err:= db.Exec(query, data.FirstName, data.LastName, data.Gender, data.Password)
	if err != nil{
		log.Println(err)
	}
	id, _= row.LastInsertId()
	return 
}



func Fetchartist(db *sql.DB, UserId string) (artist Entities.Artist, err error) {

	db, err = utils.Connecttodb()
	if err != nil {
		log.Println("unable to connct to db")
		return
	}
	query:= "select * from artist where first_name =?"
	row, err:= db.Query(query, UserId)
	if row == nil{
		fmt.Println("unable to fetch artist")
		return
	}
	
	for row.Next() {
		err = row.Scan(&artist.FirstName, &artist.LastName, &artist.Gender, &artist.Password)
		if err != nil {
			log.Println(err)
			return
		}
	}
	return 
}
