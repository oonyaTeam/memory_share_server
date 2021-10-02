package handler

import (
	"log"
	"net/http"
	"github.com/heroku/go-getting-started/httputil"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type MemoryHandler struct {
	db *sqlx.DB
}

func NewMemoryHandler(db *sqlx.DB) *MemoryHandler {
	return &MemoryHandler{
		db: db,
	}
}

func (m *MemoryHandler) GetMemories(c *gin.Context) {
	uid, err := httputil.GetUidFromToken(c)
	if err != nil {
		// clientが悪ければmiddlewareで弾かれるはずだから500
		c.JSON(http.StatusInternalServerError, gin.H{// TODO: 本当にstatus500でいい？
			"msg": err.Error(),
		})
		return
	}

	memories, err := repository.GetMemories(m.db)
	if err != nil {
		panic(err) // TODO: エラーハンドリングは適切に
	}
	// TODO: Seenを埋める
	seenMemoryList, err := repository.SeenMemoryIds(m.db, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	
	for i, memory := range memories {
		if contains(seenMemoryList, memory.Id) {
			memories[i].Seen = true
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"memories": memories,
	})
}

func contains(s []int64, e int64) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func (m *MemoryHandler) GetMyMemories(c *gin.Context) {
	uuid := "uuid"// TODO: uuidはmiddlewareでsetしたのを使う
	memories, err := repository.GetMyMemories(m.db, uuid)
	if err != nil {
		panic(err)
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

	// authorに対するerrとmemoryに対するerrを分けたいのでauthorIdを取得する処理はCreateMemoryとは分けた
	authorId, err := httputil.GetAuthorIdFromToken(m.db, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	mb.AuthorId = authorId

	err = repository.CreateMemory(m.db, mb)
	if err != nil {
		panic(err)
	}

	log.Println("bind memory=")
	log.Printf("%v\n", mb)
	c.JSON(http.StatusCreated, gin.H{
		"msg": "OK",
	})
}

