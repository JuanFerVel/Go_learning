package main

import (
	"go_mysql/database"
	"go_mysql/handlers"
	"log"

	_ "github.com/go-sql-driver/mysql" // la "_" es porque se esta usando la libreria de forma indirecta
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handlers.ListContacts(db)
}
