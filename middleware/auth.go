package middleware

import (
	"github.com/gin-gonic/gin"
	"firebase.google.com/go/auth"
	"context"
	"log"
	"github.com/pkg/errors"
)

type Auth struct {
	client *auth.Client
}

func NewAuth(client *auth.Client) *Auth {
	return &Auth{
		client: client,
	}
}

func (auth *Auth) AuthRequired(c *gin.Context) {
	log.Println("auth middle")
	idToken, err := getTokenFromHeader(c)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	token, err := auth.client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		log.Println(err)
		c.JSON(403, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	log.Printf("token: %v\n", token)
}

func getTokenFromHeader(c *gin.Context) (string, error) {
	const bearer string = "Bearer"

	header := c.GetHeader("Authorization")
	if header == "" {
		return "", errors.New("authorization header not found")
	}

	l := len(bearer)
	if len(header) > l+1 && header[:l] == bearer {
		return header[l+1:], nil
	}

	return "", errors.New("authorization header format must be 'Bearer {token}'")
}
