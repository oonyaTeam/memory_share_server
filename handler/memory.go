package handler

import (
	"log"
	"net/http"
	"github.com/heroku/go-getting-started/httputil"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/usecase"
)

type MemoryHandler struct {
	memoryUseCase *usecase.MemoryUseCase
}

func NewMemoryHandler(
	memoryUseCase *usecase.MemoryUseCase,
) *MemoryHandler {
	return &MemoryHandler{
		memoryUseCase: memoryUseCase,
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

	memories, err := m.memoryUseCase.GetMemories(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"memories": memories,
	})
}

// func (m *MemoryHandler) GetMyMemories(c *gin.Context) {
// 	uuid := "uuid"// TODO: uuidはmiddlewareでsetしたのを使う
// 	memories, err := repository.GetMyMemories(m.db, uuid)
// 	if err != nil {
// 		panic(err)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"memories": memories,
// 	})
// }

func (m *MemoryHandler) CreateMemory(c *gin.Context) {
	var mb model.Memory
	if err := c.BindJSON(&mb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	_, err := httputil.GetUidFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// err = repository.CreateMemory(
	// 	m.db,
	// 	mb.Memory,
	// 	mb.Image,
	// 	mb.Longitude,
	// 	mb.Latitude,
	// 	mb.Angle,
	// 	mb.Episodes,
	// 	uid,
	// )
	err = m.memoryUseCase.CreateMemories()
	if err != nil {
		panic(err)
	}

	log.Println("bind memory=")
	log.Printf("%v\n", mb)
	c.JSON(http.StatusCreated, gin.H{
		"msg": "OK",
	})
}

