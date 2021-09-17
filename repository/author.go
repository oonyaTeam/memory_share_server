package repository

import (
	// "github.com/heroku/go-getting-started/model"

	"database/sql"
	_ "github.com/lib/pq"

	"log"
)

func RegisterAuthor(db *sql.DB, uuid string) error {
	log.Println("register author")
	// insert
	return nil
}

func SeenMemory(db *sql.DB, uuid string, memoryId int64) error {
	log.Println("seen memory")
	// insert
	return nil
}
