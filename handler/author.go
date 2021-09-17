package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	// "github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/repository"

	"database/sql"

	_ "github.com/lib/pq"
)


type AuthorHandler struct {
	db *sql.DB
}

func NewAuthorHandler(db *sql.DB) *AuthorHandler {
	return &AuthorHandler{
		db: db,
	}
}


func (m *AuthorHandler) RegisterAuthor(c *gin.Context) {
	err := repository.RegisterAuthor(m.db, "uuid")// TODO: uuidはmiddlewareでsetしたのを使う
	if err != nil {
		panic("ee") // TODO: エラーハンドリングは適切に
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
		panic("ee") // TODO: エラーハンドリングは適切に
	}
	c.JSON(http.StatusCreated, gin.H{
		"msg": "OK",
	})
}
