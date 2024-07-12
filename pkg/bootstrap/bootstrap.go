package bootstrap

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func NewDB() (*sql.DB, error) {
	dbUrl := os.ExpandEnv("$DATABASE_USER:$DATABASE_PASSWORD@tcp($DATABASE_HOST:$DATABASE_PORT)/$DATABASE_NAME")
	log.Println(dbUrl)
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return nil, err
	}
	return db, nil
}
