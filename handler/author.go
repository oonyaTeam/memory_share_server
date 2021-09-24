package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


type AuthorHandler struct {
	db *sqlx.DB
}

func NewAuthorHandler(db *sqlx.DB) *AuthorHandler {
	return &AuthorHandler{
		db: db,
	}
}


func (m *AuthorHandler) RegisterAuthor(c *gin.Context) {
	uid, ok := c.Get("UID")
	if !ok {
		panic("not exist UID")
	}
	err := repository.RegisterAuthor(m.db, uid.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"msg": "OK",
	})
}

func (m *AuthorHandler) SeenMemory(c *gin.Context) {
	var memoryId struct{
		MemoryId int64 `json:"memory_id"`
	}
	if err := c.BindJSON(&memoryId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	log.Printf("memoryId struct: %v", memoryId)
	err := repository.SeenMemory(m.db, "uuid", memoryId.MemoryId)// TODO: uuidはmiddlewareでsetしたのを使う
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"msg": "OK",
	})
}
