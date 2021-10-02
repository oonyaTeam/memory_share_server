package httputil

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/repository"
	"github.com/jmoiron/sqlx"
)

func GetAuthorIdFromToken(db *sqlx.DB, c *gin.Context) (int64, error) {
	v, ok := c.Get("UID")
	if !ok {
		return 0, errors.New("not exist UID")
	}
	uid := v.(string)
	author_id, err := repository.GetAuthorId(db, uid)
	return author_id, err
}
