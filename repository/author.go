package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"log"
)

func RegisterAuthor(db *sqlx.DB, uuid string) error {
	log.Println("register author")
	// TODO: すでにある場合実行しない
	_, err := db.Exec(`insert into authors(uuid) values ($1)`, uuid)
	return err
}

func SeenMemory(db *sqlx.DB, uuid string, memoryId int64) error {
	log.Println("seen memory")
	// insert
	return nil
}

func GetAuthorId(db *sqlx.DB, uuid string) (int64, error) {
	var authorId int64
	err := db.Get(&authorId, `select id from authors where uuid = $1`, uuid)
	if err != nil {
		return 0, err
	} else {
		return authorId, nil
	}
}
