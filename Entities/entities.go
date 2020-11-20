package Entities

type Artist struct {
	UserId int `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Gender string `json:"gender"`
	Password string `json:"password"`
}

type Songs struct {
	Song_Id int `json:"song_id"`
	SongTitle string `json:"song_title"`
	ReleaseDate string `json:"release_date"`
}

type Category struct {
	CategoryId int `json:"category_id"`
	CategoryName int `json:"category_name"`
}

type Album struct {
	AlbumId int `json:"album_id"`
	AlbumTitle string `json:"album_title"`
}

type UserLogin struct {
	UserId int 
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`	
	
}