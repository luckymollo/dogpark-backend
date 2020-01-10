package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type CrudRepository interface {
	FindById(id string) interface{}
	FindBy(interface{}) []interface{}
	Create(interface{})
	Delete(id string)
}

const (
	host   = "localhost"
	port   = 5432
	user   = "wenjuntan"
	dbname = "wenjuntan"
)

var db *sql.DB

func InitDb() {
	psglInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := sql.Open("postgres", psglInfo)

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
