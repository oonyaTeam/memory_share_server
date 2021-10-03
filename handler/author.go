package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/httputil"
	"github.com/heroku/go-getting-started/usecase"
)

type AuthorHandler struct {
	authorUseCase *usecase.AuthorUseCase
}

func NewAuthorHandler(
	authorUseCase *usecase.AuthorUseCase,
) *AuthorHandler {
	return &AuthorHandler{
		authorUseCase: authorUseCase,
	}
}


func (m *AuthorHandler) RegisterAuthor(c *gin.Context) {
	uid, err := httputil.GetUidFromToken(c)
	if err != nil {
		// clientが悪ければmiddlewareで弾かれるはずだから500
		c.JSON(http.StatusInternalServerError, gin.H{// TODO: 本当にstatus500でいい？
			"msg": err.Error(),
		})
		return
	}
	err = m.authorUseCase.RegisterAuthor(uid)
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
	
	uid, err := httputil.GetUidFromToken(c)
	if err != nil {
		// clientが悪ければmiddlewareで弾かれるはずだから500
		c.JSON(http.StatusInternalServerError, gin.H{// TODO: 本当にstatus500でいい？
			"msg": err.Error(),
		})
		return
	}
	err = m.authorUseCase.SeenMemory(uid, memoryId.MemoryId)
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
