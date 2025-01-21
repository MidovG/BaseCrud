package connectors

import (
	"database/sql"
	"log"
)

var database *sql.DB

func Connection() {
	db, err := sql.Open("mysql", "root:@/base_crud_bd")

	if err != nil {
		log.Println(err)
	}

	database = db
	defer db.Close()
}
