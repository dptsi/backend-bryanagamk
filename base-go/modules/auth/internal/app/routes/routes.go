package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mikestefanello/hooks"
	"its.ac.id/base-go/bootstrap/web"
	"its.ac.id/base-go/modules/auth/internal/presentation/controllers"
	"its.ac.id/base-go/pkg/auth/middleware"
)

func registerRoutes(r *gin.Engine) {
	g := r.Group("/auth")
	authController := controllers.NewAuthController()

	g.POST("/login", authController.Login)
	g.GET("/callback", authController.Callback)
	g.GET("/user", middleware.Auth(), authController.User)
	g.DELETE("/logout", middleware.Auth(), authController.Logout)
}

func init() {
	web.HookBuildRouter.Listen(func(event hooks.Event[*gin.Engine]) {
		registerRoutes(event.Msg)
	})
}
