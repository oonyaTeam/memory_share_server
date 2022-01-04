package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"log"
)

func RegisterAuthor(db *sqlx.DB, uuid string) error {
	log.Println("register author")
	stmt := `insert into authors(uuid)
			select $1
			where not exists (
				select uuid from authors where uuid=$2
			)`
	_, err := db.Exec(stmt, uuid, uuid)
	return err
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
