package services

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"its.ac.id/base-go/pkg/auth/contracts"
	"its.ac.id/base-go/pkg/session"
)

func Login(ctx *gin.Context, u *contracts.User) error {
	sess := session.Default(ctx)
	sess.Set("user.id", strings.ToLower(u.Id()))
	sess.Set("user.active_role", u.ActiveRole())
	rolesJson, err := json.Marshal(u.Roles())
	if err != nil {
		return err
	}
	sess.Set("user.roles", string(rolesJson))
	sess.Save()

	return nil
}
