# HTTP Session

## Pengenalan

Karena aplikasi berbasis HTTP tidak memiliki state, session menyediakan cara untuk menyimpan informasi tentang pengguna di beberapa request. Informasi pengguna tersebut biasanya ditempatkan di penyimpanan/backend persisten yang dapat diakses dari request berikutnya.

## Konfigurasi

Konfigurasi dari session aplikasi disimpan pada berkas `bootstrap/web/web.go`. Secara bawaan, base project ini dikonfigurasi menggunakan SQLite sebagai adapter dari session. Adapter tersebut menentukan dimana data dari session akan disimpan untuk setiap request. Berikut adalah adapter bawaan yang ada di dalam base project ini:

- Firestore - session disimpan di dalam database [Firestore](https://cloud.google.com/firestore)
- GORM - session disimpan di dalam database yang didukung oleh [GORM](https://gorm.io/)

## Prasyarat Adapter

### Firestore

Dengan menggunakan adapter ini, maka Anda perlu mengkonfigurasi Firestore dengan mengisi nilai `SESSION_FIRESTORE_PROJECT_ID` pada `.env` dengan nilai yang sesuai dengan Project ID yang akan digunakan.

### GORM

Buat koneksi ke database sesuai dengan [instruksi pada GORM](https://gorm.io/docs/connecting_to_the_database.html) kemudian panggil fungsi `adapters.NewGorm()` dengan argumen koneksi database yang telah dibuat.

## Interaksi Dengan Session

Instance dari session dapat diambil melalui method `Default()` pada `its.ac.id/base-go/pkg/session` dengan memberi argumen berupa `*gin.Context`.

```go
package something

import "its.ac.id/base-go/pkg/session"

func Handler(ctx *gin.Context) {
    // Mendapatkan instance dari session saat ini
    sess := session.Default(ctx)

}
```

### Mendapatkan data

Data dari session dapat diperoleh dari method `Get()` dari instance session. Method tersebut mengembalikan dua nilai, yaitu data yang ada dan sebuah `boolean` yang menunjukkan apakah key yang diberikan ada di dalam session tersebut.

```go
package something

import "its.ac.id/base-go/pkg/session"

func Handler(ctx *gin.Context) {
    // Mendapatkan instance dari session saat ini
    sess := session.Default(ctx)

    // Mendapatkan data dari session dengan key `user`
    user, exists := sess.Get("user")
}
```

### Menyimpan Data

Untuk menyimpan data ke session, Anda dapat menggunakan method `Set()` pada instance session dengan memberikan argumen yaitu key dan value.

```go
package something

import "its.ac.id/base-go/pkg/session"

func Handler(ctx *gin.Context) {
    // Mendapatkan instance dari session saat ini
    sess := session.Default(ctx)

    // Menyimpan ke session
    sess.Set("user.name", "DPTSI ITS")
    sess.Set("user.id", 11)

    // Setiap melakukan perubahan, wajib memanggil method `Save()` setidaknya sekali.
    // Umumnya dipanggil sekali setelah semua action yang mengubah session telah dipanggil.
	if err := sess.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "unable_to_save_session",
			"data":    nil,
		})
		return
	}
}
```

### Menghapus Data

Method `Delete()` dapat digunakan untuk menghapus data dengan key tertentu dari session.

```go
package something

import "its.ac.id/base-go/pkg/session"

func Handler(ctx *gin.Context) {
    // Mendapatkan instance dari session saat ini
    sess := session.Default(ctx)

    // Menghapus dari session
    sess.Delete("user.name")
    sess.Delete("user.id")

    // Setiap melakukan perubahan, wajib memanggil method `Save()` setidaknya sekali.
    // Umumnya dipanggil sekali setelah semua action yang mengubah session telah dipanggil.
	if err := sess.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "unable_to_save_session",
			"data":    nil,
		})
		return
	}
}
```

### Regenerate Session ID

Proses ini dilakukan untuk menghindari eksploitasi dengan [Session fixation](https://owasp.org/www-community/attacks/Session_fixation) pada aplikasi. Pada umumnya, hal ini dilakukan ketika ada perubahan privilege dari user, contohnya adalah ketika user berhasil login atau mengganti role.

```go
package something

import "its.ac.id/base-go/pkg/session"

func Handler(ctx *gin.Context) {
    // Mendapatkan instance dari session saat ini
    sess := session.Default(ctx)

    // Regenerate session ID
	sess.Regenerate()

    // Setiap melakukan perubahan, wajib memanggil method `Save()` setidaknya sekali.
    // Umumnya dipanggil sekali setelah semua action yang mengubah session telah dipanggil.
	if err := sess.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "unable_to_save_session",
			"data":    nil,
		})
		return
	}

    // Fungsi ini wajib dipanggil untuk mengganti session ID yang ada pada cookie dengan nilai baru yang digenerate
	session.AddCookieToResponse(ctx, sess.Id())
}
```

## Menambahkan Custom Session Adapter

### Mengimplementasikan Adapter

Jika adapter yang sudah disediakan tidak memenuhi kriteria Anda, maka Anda perlu mengimplementasikan adapter yang memenuhi kontrak sebagai berikut:

```go
type Storage interface {
    // Membaca instance dari session data dari penyimpanan
	Get(ctx *gin.Context, id string) (*Data, error)
    // Menyimpan instance session data ke penyimpanan
	Save(ctx *gin.Context, id string, data map[string]interface{}, expiredAt time.Time, csrfToken string) error
    // Menghapus data terkait dengan ID tersebut dari penyimpanan
	Delete(ctx *gin.Context, id string) error
}
```

### Mendaftarkan Adapter

Anda dapat melakukan inject adapter tersebut pada fungsi `init()` di berkas `bootstrap/web/web.go`.

```go
func init() {
    do.Provide[session.Storage](e.Msg, func(i *do.Injector) (session.Storage, error) {
        // Return implementasi dari adapter
    })
}
```
