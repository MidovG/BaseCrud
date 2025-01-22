package connectors

import (
	"database/sql"
	"log"
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:@/base_crud_bd")

	if err != nil {
		log.Println(err)
	}

	return db
}
