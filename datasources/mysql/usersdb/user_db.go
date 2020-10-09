package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// Client return db connection object
var (
	Client *sql.DB
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?character=utf8", "root", "toor", "127.0.0.1", "users_db")
	Client, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
