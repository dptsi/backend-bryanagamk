package services

import (
	"github.com/gin-gonic/gin"
	"its.ac.id/base-go/pkg/session"
)

func Logout(ctx *gin.Context) error {
	sess := session.Default(ctx)
	sess.Delete("user.id")
	sess.Delete("user.active_role")
	sess.Delete("user.roles")
	sess.Save()

	return nil
}
