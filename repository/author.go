package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"log"
)

func RegisterAuthor(db *sqlx.DB, uuid string) error {
	log.Println("register author")
	// insert
	return nil
}

func SeenMemory(db *sqlx.DB, uuid string, memoryId int64) error {
	log.Println("seen memory")
	// insert
	return nil
}
