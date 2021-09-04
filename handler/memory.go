package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/repository"

	"database/sql"

	_ "github.com/lib/pq"
)

type MemoryHandler struct {
	db *sql.DB
}

func NewMemoryHandler(db *sql.DB) *MemoryHandler {
	return &MemoryHandler{
		db: db,
	}
}

func (m *MemoryHandler) GetMemories(c *gin.Context) {
	memories, err := repository.GetMemories(m.db)
	if err != nil {
		panic("ee") // TODO: エラーハンドリングは適切に
	}
	c.JSON(http.StatusOK, gin.H{
		"memories": memories,
	})
}

// func GetMyMemories(c *gin.Context) {
func (m *MemoryHandler) GetMyMemories(c *gin.Context) {
	// uuid := c.Query("uuid")
	uuid := "uuid"
	memories, err := repository.GetMyMemories(m.db, uuid)
	if err != nil {
		panic("err")
	}

	c.JSON(http.StatusOK, gin.H{
		"memories": memories,
	})
}

func (m *MemoryHandler) CreateMemory(c *gin.Context) {
	var mb model.Memory
	if err := c.BindJSON(&mb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	err := repository.CreateMemory(m.db, mb)
	if err != nil {
		panic("eerr")
	}

	log.Println("bind memory=")
	log.Printf("%v\n", mb)
	c.JSON(http.StatusCreated, gin.H{
		"msg": "OK",
	})
}

