# Modul

Aplikasi modular monolithic dibagi menjadi beberapa modul yang merupakan implementasi dari Bounded Context pada Domain-Driven Desain. Pada base project ini, modul didesain sedemikian rupa sehingga memiliki sifat loosely coupled dari/ke aplikasi utama maupun modul lainnya.

## Aturan Dalam Modul

Untuk menjaga agar modul tetap bersifat loosely coupled, maka aturan-aturan berikut harus dipenuhi:

- Modul hanya diizinkan mengimpor apa yang ada pada `bootstrap` (terutama hook) dan `pkg` (shared library) dari aplikasi utama
- Setiap modul dilarang mengimpor apa yang ada pada modul lain
- Setiap modul wajib memiliki nama dan URI prefix yang unik
- Nama modul wajib menggunakan huruf kecil dan snake case

## Pembuatan modul

Modul dapat dibuat dengan menjalankan command berikut:

```bash
go run ./script/script.go make:module <nama modul>
```

Anda akan diberikan pertanyaan sebagai berikut setelah menjalankan command untuk memilih pattern yang digunakan:

```bash
Do you want to use transaction script instead of aggregate pattern? (y/N):
```

Jika sudah, maka modul akan ada pada direktori `modules`.

## Mendaftarkan Modul

Modul dapat didaftarkan/diaktifkan dengan mengimpornya pada `main.go`. Anda hanya perlu memberi alias `_` pada modul yang diimpor.

```go

import (
    // Modules
	_ "its.ac.id/base-go/modules/auth"
    _ "its.ac.id/base-go/modules/<nama modul>" // Modul baru
)
```

## Menonaktifkan Modul

Modul dapat dinonaktifkan tanpa perlu menghapus codenya dengan menghapus/memberi komentar pada baris impor di `main.go`.
