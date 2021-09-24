package httputil

import (
	"github.com/heroku/go-getting-started/repository"
	"github.com/jmoiron/sqlx"
)

func GetUidFromToken(db *sqlx.DB,v interface{}) int64 {
	uid := v.(string)
	author_id, err := repository.GetAuthorId(db, uid)
	if err != nil {
		panic(err)
	}
	return author_id
}