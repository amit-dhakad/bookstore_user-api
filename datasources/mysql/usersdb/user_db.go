package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/joho/godotenv"
)

const (
	mysqlUsersUsername = "mysqlUsersUsername"
	mysqlUsersPassword = "mysqlUsersPassword"
	mysqlUsersHost     = "mysqlUsersHost"
	mysqlUserSchema    = "mysqlUserSchema"
)

// Client return db connection object

var (
	Client *sql.DB

	username = getEnvVariable(mysqlUsersUsername)
	password = getEnvVariable(mysqlUsersPassword)
	host     = getEnvVariable(mysqlUsersHost)
	schema   = getEnvVariable(mysqlUserSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema)
	Client, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}

func getEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
