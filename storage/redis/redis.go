package redis

import (
	"database/sql"
	"log"
)

type Storage struct {
	db *sql.DB
}

func New(url string) *sql.DB {

	db, err := sql.Open("postgers", "./halyard-bot.db")
	if err != nil {
		log.Fatal("Can't open Redis connection")
	}
	return db
}
