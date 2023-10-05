package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"its.ac.id/base-go/pkg/app/common"
	"its.ac.id/base-go/pkg/auth/contracts"
	"its.ac.id/base-go/pkg/auth/internal/utils"
	"its.ac.id/base-go/pkg/session"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess := session.Default(ctx)
		idIf, ok := sess.Get("user.id")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
			return
		}
		// TODO: Unserialize roles
		activeRoleIf, ok := sess.Get("user.active_role")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
			return
		}
		rolesJsonIf, ok := sess.Get("user.roles")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
			return
		}
		activeRole, ok := activeRoleIf.(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
			return
		}
		rolesJson, ok := rolesJsonIf.(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
			return
		}
		var roles []contracts.Role
		err := json.Unmarshal([]byte(rolesJson), &roles)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
		}

		id, ok := idIf.(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
			return
		}

		u := contracts.NewUser(id)
		for _, role := range roles {
			u.AddRole(role.Name, role.Permissions, role.IsDefault)
		}
		u.SetActiveRole(activeRole)

		ctx.Set(utils.UserKey, u)
		ctx.Next()
	}
}
