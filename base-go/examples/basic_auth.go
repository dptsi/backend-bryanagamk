package examples

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"its.ac.id/base-go/pkg/auth/adapters"
	"its.ac.id/base-go/pkg/auth/middleware"
	"its.ac.id/base-go/pkg/auth/services"
)

func basicAuthRoutes(r *gin.Engine) {
	// Load semua user dari env atau hardcoded
	users := make([]adapters.SliceUser, 0)
	users = append(users, adapters.SliceUser{
		Id:             "9a08d515-522c-4f5b-954a-13755bf61b2f",
		Username:       "zydhanlinnar11",
		HashedPassword: "$2a$10$nY0Xkk.8zTUNQFL4fiKCKe5djphBJN5AyAq.LDnxKwOWJ1W/D9Txa",
	})

	// Inisialisasi userRepo dengan adapter untuk mendapatkan user dari slice / array
	userRepo := adapters.NewSliceUserRepository(users)

	// Inisiaisasi middleware dengan userRepo
	authMiddleware := middleware.NewBasicAuthMiddleware(userRepo)

	// Tambahkan middleware ke route yang ingin di-protect
	r.GET("/basic-auth", authMiddleware.Handle(), func(c *gin.Context) {
		user := services.User(c)

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "user",
			"data": gin.H{
				"id":          user.Id(),
				"active_role": nil,
				"roles":       make([]string, 0),
			},
		})
	})
}
