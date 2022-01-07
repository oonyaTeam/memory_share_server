package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oonyaTeam/memory_share_server/httputil"
	"github.com/oonyaTeam/memory_share_server/usecase"
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
