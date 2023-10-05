# Routing

## Routing Dasar

Routing pada base project ini menggunakan [gin-gonic](https://gin-gonic.com/) sebagai enginenya. Pada saat modul dibuat, prefix yang digunakan adalah nama modul dengan underscore (`_`) yang direplace dengan dash (`-`). Anda dapat mengubah prefix tersebut melalui konstanta `routePrefix`. Routing dapat diatur pada berkas `internal/app/routes/routes.go` di fungsi `registerRoutes()`.

```go
const routePrefix = "/nama-modul"
```

Handler pada Gin pada dasarnya sederhana yaitu dengan memanggil method sesuai dengan HTTP method dengan parameter uri, serta handler.

```go
g := r.Group(routePrefix)

// Register routes below
g.GET("/test", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "test",
    })
})
```

## Aturan Dalam Routing

- Prefix modul harus unik
- Prefix modul harus terdiri dari alfanumerik dan dash (`-`)

## Route Group

Route group berfungsi untuk memberi atribut seperti prefix maupun middleware pada beberapa route sekaligus tanpa perlu mendefinisikan pada setiap route. Untuk menggunakan group, anda perlu menggunakan method `Group()` pada pointer `gin.Engine`.

```go
func registerRoutes(r *gin.Engine) {
    // Membuat route group dengan prefix contoh-group
    g := r.Group("/contoh-group")
}
```

### Middleware

Middleware dapat digunakan pada route group dengan memanggil method `Use()` pada pointer `gin.RouterGroup`.

```go
func registerRoutes(r *gin.Engine) {
    // Membuat route group dengan prefix contoh-group
    g := r.Group("/contoh-group")

    // Memasangkan middleware
	g.Use(middleware.Auth())
}
```

## Cross-Origin Resource Sharing (CORS)

Base project ini otomatis menghandle request dengan method `OPTIONS` untuk melayani CORS. Akan tetapi, beberapa skenario mungkin memerlukan konfigurasi yang berbeda yang dapat diatur pada `.env`. Berikut konfigurasi bawaan dari base project ini:

```bash
CORS_PATH=*
CORS_ALLOWED_METHODS=*
CORS_ALLOWED_ORIGINS=*
CORS_ALLOWED_HEADERS=
CORS_EXPOSED_HEADERS=
CORS_MAX_AGE=0
CORS_SUPPORT_CREDENTIALS=false
```
