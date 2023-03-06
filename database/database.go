package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Open() *sql.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic(errEnv.Error())
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, pass, name)
	db, errDb := sql.Open("postgres", dsn)

	if errDb != nil {
		panic(errDb.Error())
	}
	if pingErr := db.Ping(); pingErr != nil {
		panic(pingErr.Error())
	}
	return db
}

func Close(db *sql.DB) {

	if err := db.Close(); err != nil {
		panic(err.Error())
	}
}
