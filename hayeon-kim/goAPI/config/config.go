package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := "123qweasdzxc"
	dbName := "go_api"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	return
}
