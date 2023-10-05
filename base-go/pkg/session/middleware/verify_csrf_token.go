package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"its.ac.id/base-go/pkg/session"
)

var ErrInvalidCSRFToken = errors.New("invalid_csrf_token")
var MethodsWithoutCSRFToken = []string{"GET", "HEAD", "OPTIONS"}

func VerifyCSRFToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess := session.Default(ctx)
		sessionCSRFToken := sess.CSRFToken()
		requestCSRFToken := ctx.Request.Header.Get("X-CSRF-TOKEN")

		// Skip CSRF token verification for some methods
		for _, method := range MethodsWithoutCSRFToken {
			if ctx.Request.Method == method {
				ctx.Next()
				return
			}
		}

		if sessionCSRFToken == "" || sessionCSRFToken != requestCSRFToken {
			ctx.AbortWithStatusJSON(403, gin.H{
				"code":    403,
				"message": ErrInvalidCSRFToken.Error(),
				"data":    nil,
			})
			return
		}

		ctx.Next()
	}
}
