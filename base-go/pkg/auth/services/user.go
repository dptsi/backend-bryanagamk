package services

import (
	"github.com/gin-gonic/gin"
	"its.ac.id/base-go/pkg/auth/contracts"
	"its.ac.id/base-go/pkg/auth/internal/utils"
)

func User(c *gin.Context) *contracts.User {
	uInterface, exist := c.Get(utils.UserKey)
	if !exist {
		return nil
	}
	u, ok := uInterface.(*contracts.User)
	if !ok {
		return nil
	}

	return u
}
