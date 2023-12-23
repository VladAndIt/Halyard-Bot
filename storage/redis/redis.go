package redis

import "database/sql"

type Storage struct {
	db *sql.DB
}

func New(url string) (*Storage, error) {

	db, err := sql.Open("postgers", "./halyard-bot.db")
}
