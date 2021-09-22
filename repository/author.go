package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"log"
)

func RegisterAuthor(db *sqlx.DB, uuid string) error {
	log.Println("register author")
	// insert
	_, err := db.Exec(`insert into authors(uuid) values ($1)`, uuid)
	return err
}

func SeenMemory(db *sqlx.DB, uuid string, memoryId int64) error {
	log.Println("seen memory")
	// insert
	return nil
}
