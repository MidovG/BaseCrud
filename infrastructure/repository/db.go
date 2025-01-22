package repository

import (
	"database/sql"
)

type SQLUserRepository struct {
	DB *sql.DB
}
