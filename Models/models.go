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



func FetchArtist(db *sql.DB, UserId int) (artist Entities.Artist, err error) {

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

func CreateAlbum(db *sql.DB, data Entities.Album) (id int64, err error){
	query:= "insert into album (album_id, album_title) values (?,?)"
	db, err = utils.Connecttodb()
	if err !=nil{
		log.Println("Unable to connect to db")
		return
	}
	row, err:= db.Exec(query, data.AlbumId, data.AlbumTitle)
	if err != nil{
		log.Println(err)
	}
	id, _= row.LastInsertId()
	return
}

func FetchAlbum (db *sql.DB, AlbumId int) (album Entities.Album, err error) {

	db, err = utils.Connecttodb()
	if err != nil {
		log.Println("unable to connct to db")
		return
	}
	query:= "select * from album where album_title =?"
	row, err:= db.Query(query, AlbumId)
	if row == nil{
		fmt.Println("unable to fetch album")
		return
	}

	for row.Next() {
		err = row.Scan(&album.AlbumId, &album.AlbumTitle)
		if err != nil {
			log.Println(err)
			return
		}
	}
	return
}

func CreateCategory(db *sql.DB, data Entities.Category) (id int64, err error){
	query := "insert into category (category_id, category_name) values (?,?)"
	db, err = utils.Connecttodb()
	if err!=nil{
		log.Println("Unable to connect to db")
		return
	}
	row, err:= db.Exec(query, data.CategoryId, data.CategoryName)
	if err != nil{
		log.Println(err)
	}
	id, _= row.LastInsertId()
	return
}

func FetchCategory(db *sql.DB, CategoryId int) (category Entities.Category, err error){
	db, err =utils.Connecttodb()
	if err != nil{
		log.Println("unable to connect to db")
		return
	}
	query := "select * from category where category_name =?"
	row, err := db.Query(query, CategoryId)
	if row == nil{
		fmt.Println("Cannot read category")
		return
	}

	for row.Next(){
		err = row.Scan(&category.CategoryId, &category.CategoryName)
		if err != nil{
			log.Println(err)
			return
		}
	}
	return
}

func CreateSong (db *sql.DB, data Entities.Songs) (id int64, err error) {
	query := "inserts into songs (song_id, song_title, release_date) values (?,?,?)"
	db, err = utils.Connecttodb()
	if err != nil {
		log.Println("Unable to connect to db")
		return
	}
	row, err := db.Exec(query, data.Song_Id, data.SongTitle, data.ReleaseDate)
	if err != nil {
		log.Println(err)
	}
	id, _= row.LastInsertId()
	return
}

func FetchSongs (db *sql.DB, SongId int) (songs Entities.Songs, err error) {
	db, err =utils.Connecttodb()
	if err != nil{
		log.Println("unable to connect to db")
		return
	}
	query := "select * from songs where song_title = ?"
	row, err := db.Query(query, SongId)
	if row == nil{
		fmt.Println("Cannot find song")
		return
	}
	for row.Next(){
		err = row.Scan(&songs.Song_Id, &songs.SongTitle, &songs.ReleaseDate)
		if err != nil{
			log.Println(err)
			return
		}
	}
	return
}