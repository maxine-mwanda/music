package utils

import (
	"database/sql"
	"fmt"
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
