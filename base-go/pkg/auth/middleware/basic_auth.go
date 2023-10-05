package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"its.ac.id/base-go/pkg/app/common"
	"its.ac.id/base-go/pkg/auth/contracts"
	"its.ac.id/base-go/pkg/auth/internal/utils"
)

type BasicAuthMiddleware struct {
	userRepo contracts.UserRepository
}

func NewBasicAuthMiddleware(userRepo contracts.UserRepository) *BasicAuthMiddleware {
	return &BasicAuthMiddleware{
		userRepo: userRepo,
	}
}

func (m *BasicAuthMiddleware) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		raw := ctx.GetHeader("Authorization")

		if raw == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
			return
		}

		authorization := strings.Split(raw, " ")
		if len(authorization) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
			return
		}
		if authorization[0] != "Basic" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
			return
		}
		base64Creds := authorization[1]

		decoded, err := base64.StdEncoding.DecodeString(base64Creds)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
			return
		}

		creds := strings.Split(string(decoded), ":")
		if len(creds) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
			return
		}

		username := creds[0]
		password := creds[1]

		user, err := m.userRepo.FindByUsername(username)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.InternalServerErrorResponse)
			return
		}
		if user == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword()), []byte(password))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.UnauthorizedResponse)
			return
		}

		ctx.Set(utils.UserKey, user)
		ctx.Next()
	}
}
