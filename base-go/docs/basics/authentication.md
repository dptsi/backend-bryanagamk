# Authentication

## Melakukan Autentikasi Manual

Anda dapat melakukan autentikasi dengan memanggil fungsi `Login()` pada package `its.ac.id/base-go/pkg/auth/services`. Fungsi tersebut menerima argumen berupa `*gin.Context` dan instance user yang akan dilakukan proses autentikasi. Setelah autentikasi berhasil, Anda harus melakukan generate ulang user session untuk mencegah [session fixation](https://owasp.org/www-community/attacks/Session_fixation).

```go
func HandleLogin(ctx *gin.Context) {
    // Melakukan login
    if err := services.Login(ctx, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "login_failed",
			"data":    nil,
		})
		return
	}

    // Regenerate session ID dan menyimpannya
	sess := session.Default(ctx)
	sess.Regenerate()
	if err := sess.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "unable_to_save_session",
			"data":    nil,
		})
		return
	}

    // Mengganti session ID pada cookie ke nilai yang baru
	session.AddCookieToResponse(ctx, sess.Id())

    // Response login berhasil
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "login_success",
		"data":    nil,
	})
}
```

Route yang ingin diproteksi dapat diproteksi dengan menggunakan middleware `Auth()` pada package `its.ac.id/base-go/pkg/auth/middleware`.

```go
func registerRoutes(r *gin.Engine) {
	g := r.Group("/auth")
	authController := controllers.NewAuthController()

	g.GET("/user", middleware.Auth(), authController.User)
}
```

## HTTP Basic Authentication

Metode ini mengizinkan Anda untuk mengautentikasi pengguna menggunakan header `Authorization: Basic xxxxxx`.

### Prasyarat

Anda perlu membuat adapter untuk mendapatkan user dari penyimpanan yang mengimplementasikan interface berikut:

```go
type UserRepository interface {
	FindByUsername(username string) (*User, error)
}
```

Secara bawaan, tersedia adapter untuk mendapatkan user dari sebuah slice/array:

```go
func initiateUserRepository() contracts.UserRepository {
    // Load semua user dari env atau hardcoded
	users := make([]adapters.SliceUser, 0)
	users = append(users, adapters.SliceUser{
		Id:             "9a08d515-522c-4f5b-954a-13755bf61b2f",
		Username:       "zydhanlinnar11",
		HashedPassword: "$2a$10$nY0Xkk.8zTUNQFL4fiKCKe5djphBJN5AyAq.LDnxKwOWJ1W/D9Txa",
	})

	// Inisialisasi userRepo dengan adapter untuk mendapatkan user dari slice / array
	return adapters.NewSliceUserRepository(users)
}
```

### Proses Autentikasi

Anda perlu menginisiasi middleware basic auth dengan memberikan argumen berupa user repository di atas kemudian menggunakannya pada handler function pada rute yang ingin diproteksi.

```go
func basicAuthRoutes(r *gin.Engine) {
	// Inisiaisasi middleware dengan userRepo
	authMiddleware := middleware.NewBasicAuthMiddleware(initiateUserRepository())

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
```

## Mendapatkan User

User dapat diperoleh melalui method `User()` dengan argumen `*gin.Context` pada package `its.ac.id/base-go/pkg/auth/services`.

```go
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
```

## Log Out

Untuk melakukan log out, Anda dapat memanggil fungsi `Logout()`. Direkomendasikan untuk invalidate user session dan melakukan regenerate CSRF token setelah proses log out.

```go
func HandleLogout(ctx *gin.Context) {
    // Melakukan proses log out
    if err = services.Logout(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "logout_failed",
			"data":    nil,
		})
		return
	}

    // Invalidate session, regenerate CSRF token, dan menyimpannya
	sess := session.Default(ctx)
	sess.Invalidate()
	sess.RegenerateCSRFToken()
	if err := sess.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "unable_to_save_session",
			"data":    nil,
		})
		return
	}

    // Mengganti session ID pada cookie ke nilai yang baru
	session.AddCookieToResponse(ctx, sess.Id())

    // Response logout berhasil
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "logout_success",
		"data":    nil,
	})
}
```
