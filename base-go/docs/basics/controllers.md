# Controllers

Daripada mendefinisikan semua logika penanganan permintaan Anda di file `routes.go`, Anda mungkin ingin mengatur perilaku ini menggunakan kelas "controllers". Controller dapat mengelompokkan logika handle request terkait ke dalam satu kelas. Misalnya, kelas UserController mungkin menangani semua request yang terkait dengan pengguna, termasuk menampilkan, membuat, memperbarui, dan menghapus pengguna. Secara default, controllers disimpan di direktori internal/presentation/controllers pada setiap modul.

## Controller Dasar

Untuk membuat controller baru dengan cepat, Anda dapat menjalankan command `make:controller`.

```bash
go run ./script/script.go make:controller <nama modul> <nama controller>
```

> Nama controller harus berupa pascal case tanpa suffix "Controller"

Berikut adalah contoh controller sederhana yang memiliki satu public/exported method.

```go
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type SemesterController struct {
	i *do.Injector
}

func NewSemesterController() *SemesterController {
	i := do.DefaultInjector
	return &SemesterController{i: i}
}


func (c *SemesterController) Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello SemesterController",
	})
}
```

Dengan demikian, kode pada `registerRoutes()` dapat diubah dengan inisiasi dan memanggil controller.

```go
func registerRoutes(r *gin.Engine) {
	g := r.Group(routePrefix)

	// Controller initialization
	smtController := controllers.NewSemesterController()

	// Register routes below
	g.GET("/hello", smtController.Hello)
}
```
