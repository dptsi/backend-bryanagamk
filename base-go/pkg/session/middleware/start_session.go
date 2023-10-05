package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"its.ac.id/base-go/bootstrap/config"
	"its.ac.id/base-go/pkg/session"
)

func StartSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		storage, err := do.Invoke[session.Storage](do.DefaultInjector)
		if err != nil {
			panic(err)
		}
		if storage == nil {
			panic("Session storage not configured. Please configure it first in bootstrap/web/web.go")
		}
		cfg := do.MustInvoke[config.Config](do.DefaultInjector).Session()

		// Initialize session data
		var data *session.Data
		sessionId, err := ctx.Cookie(cfg.CookieName)

		if err == nil {
			// Get session data from storage
			sess, err := storage.Get(ctx, sessionId)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "unable_to_get_session_data",
					"data":    nil,
				})
				return
			}
			if sess != nil {
				data = sess
			}
		}
		if data == nil {
			data = session.NewEmptyData(ctx, storage)
			if err := data.Save(); err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "unable_to_save_session_data",
					"data":    nil,
				})
				return
			}
		}
		ctx.Set("session", data)
		session.AddCookieToResponse(ctx, data.Id())
		ctx.Next()
	}
}
