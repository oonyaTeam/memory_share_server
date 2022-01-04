package httputil

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func GetUidFromToken(c *gin.Context) (string, error) {
	v, ok := c.Get("UID")
	if !ok {
		return "", errors.New("not exist UID")
	}
	uid, ok := v.(string)
	if !ok {
		return "", errors.New("cannot cast UID to string, uid must be string")
	}
	return uid, nil
}
