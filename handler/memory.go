package handler

import (
	"log"
	"net/http"
	"github.com/heroku/go-getting-started/httputil"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/usecase"

	"strconv"
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

func (m *MemoryHandler) GetMyMemories(c *gin.Context) {
	uid, err := httputil.GetUidFromToken(c)
	if err != nil {
		// clientが悪ければmiddlewareで弾かれるはずだから500
		c.JSON(http.StatusInternalServerError, gin.H{// TODO: 本当にstatus500でいい？
			"msg": err.Error(),
		})
		return
	}
	memories, err := m.memoryUseCase.GetMyMemories(uid)
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

func (m *MemoryHandler) CreateMemory(c *gin.Context) {
	var memory model.Memory
	if err := c.BindJSON(&memory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	uid, err := httputil.GetUidFromToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = m.memoryUseCase.CreateMemories(&memory, uid)
	if err != nil {
		panic(err)// TODO: error handling
	}

	log.Println("bind memory=")
	log.Printf("%v\n", memory)
	c.JSON(http.StatusCreated, gin.H{
		"memory": memory,
	})
}

func (m *MemoryHandler) DeleteMemory(c *gin.Context) {
	memoryId := c.Query("memory_id")
	aid, err := strconv.Atoi(memoryId); 
	if err != nil  {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "invalid query string",
		})
		return
	}
	
	// TODO: BadRequestではなくね？
	err = m.memoryUseCase.DeleteMemories(aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusNoContent, gin.H{
		"msg": "ok",
	})
}

func (m *MemoryHandler) SeenMemory(c *gin.Context) {
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
	
	uid, err := httputil.GetUidFromToken(c)
	if err != nil {
		// clientが悪ければmiddlewareで弾かれるはずだから500
		c.JSON(http.StatusInternalServerError, gin.H{// TODO: 本当にstatus500でいい？
			"msg": err.Error(),
		})
		return
	}
	err = m.memoryUseCase.SeenMemory(uid, memoryId.MemoryId)
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
