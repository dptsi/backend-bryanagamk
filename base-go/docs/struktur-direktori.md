# Struktur Direktori

Base Go ini didesain sedemikian rupa sehingga cocok untuk aplikasi berskala besar maupun kecil dengan sifat yang modular. Secara default, arsitektur yang digunakan adalah Modular Monolithic dengan Aggregate Pattern maupun Transaction Script Pattern yang ada pada Domain-Driven Design (DDD).

Sebagian besar kode disimpan pada folder [internal](https://go.dev/doc/go1.4#internalpackages) untuk menjaga agar penggunaan nama/istilah yang sama tidak terimpor secara tidak sengaja pada Bounded Context/modul lain.

## Direktori Root

### Direktori `bootstrap`

Direktori ini berfungsi untuk melakukan setup pada config, event handler, dan routing pada aplikasi dan setiap modul di dalam aplikasi.

### Direktori `docs`

Direktori ini berisi terkait dokumentasi penggunaan base project berupa markdown yang dicompile menggunakan [MkDocs](https://www.mkdocs.org/) dengan tema [mkdocs-material](https://github.com/squidfunk/mkdocs-material).

### Direktori `examples`

Direktori ini berisi mengenai kode snippet untuk contoh pada setiap kasus, misal penggunaan autentikasi, dan lain-lain.

### Direktori `modules`

Direktori ini adalah core dari aplikasi yang akan dieksplor lebih detail pada bagian selanjutnya. Sebagian besar dari kode yang akan Anda tulis ada pada direktori ini.

### Direktori `pkg`

Direktori ini berisikan shared library yang dapat digunakan oleh semua modul pada aplikasi.

### Direktori `script`

Direktori ini berisikan kode sumber untuk automated script yang membantu dalam pengembangan aplikasi.

## Direktori `modules`

Direktori `modules` berisi modul-modul yang ada dalam aplikasi. Setiap modul merupakan implementasi dari sebuah Bounded Context pada Domain-Driven Design. Direktori ini harus berisi beberapa folder yang menyatakan nama dari modul. Struktur setiap modul disesuaikan dengan Onion/Hexagonal/Clean Architecture dengan Command Query Responsibility Segregation (CQRS) Pattern yang terdiri dari direktori sebagai berikut:

### Direktori `internal/app/config`

Direktori ini berisikan konfigurasi dari setiap modul. Alur yang dikerjakan dari layer ini adalah sebagai berikut:

1. Mengambil data environment variable
2. Melakukan mapping ke struct config
3. Inject config agar dapat dipakai di seluruh aplikasi

### Direktori `internal/app/controllers`

Direktori ini berisikan handler dari request yang masuk ke dalam aplikasi. Alur yang dikerjakan dari layer ini adalah sebagai berikut:

1. Menerima request masuk
2. Validasi request (request body/query/params)
3. Memanggil command/query yang diperlukan untuk mengerjakan use case
4. Handle error yang dikeluarkan oleh command/query terkait
5. Memberikan response terhadap request

### Direktori `internal/app/commands`

Direktori ini berisikan command handler yang memiliki sifat **mengubah state** dari sistem yang artinya melakukan perubahan data yang disimpan pada database sistem. Alur yang dikerjakan dari layer ini adalah sebagai berikut:

1. Menerima request masuk dari controller
2. Memanggil repository terkait untuk mengambil state eksisting pada sebuah agregat
3. Memanggil exported/public method yang ada pada agregat untuk mengubah state
4. (Opsional) Publish event ke event dispatcher
5. Memanggil repository untuk menyimpan state ke database

### Direktori `internal/app/listeners`

Direktori ini memiliki banyak kesamaan peran dan alur seperti `internal/app/commands`. Perbedaannya adalah listeners dipanggil oleh aplikasi (lebih tepatnya event dispatcher) ketika event sudah dipublish. Listener tidak dapat memberikan return value apapun (void method) dikarenakan layer ini tidak dihandle oleh controller. Apapun hal yang ingin dicatat pada layer ini harus dilakukan melalui logging.

### Direktori `internal/app/queries`

Direktori ini berisikan query object yang hanya bersifat read-only yang berarti tidak dapat mengubah state yang tersimpan pada sistem. Layer ini berperan untuk melayani use-case yang hanya membaca data. Alur kerja layer ini adalah sebagai berikut:

1. Menerima panggilan dari controller
2. Mengambil data dari database
3. Mengembalikan nilainya ke controller

### Direktori `internal/app/routes`

Direktori ini berisikan pengaturan routing dari setiap endpoint yang akan disediakan (method, middleware, URI, dan handler). Direktori ini juga berperan dalam memberikan implementasi konkrit ke Dependency Injection.

### Direktori `internal/infrastructures`

Direktori ini berisikan implementasi dari interface yang didefinisikan pada query handler, repository, atau lainnya.

### Direktori `internal/domain/services`

Jika Anda menggunakan Transaction Script Pattern, maka direktori ini memiliki peran sebagai berikut:

- Melakukan validasi business rules/invariant

Jika Anda menggunakan Aggregate Pattern, maka penggunaan dari domain services hanya untuk melakukan komputasi dari beberapa agregat dengan alur sebagai berikut:

1. Memanggil repositori dari beberapa agregat
2. Melakukan komputasi
3. Mengembalikan nilai

### Direktori `internal/domain/repositories`

Direktori ini berisikan interface untuk mendapatkan dan menyimpan data yang diperlukan oleh domain services/command handler/event listener dari setiap agregat.

## Struktur Direktori Khusus Aggregate Pattern

### Direktori `internal/domain/entities`

Direktori ini berisikan entitas maupun aggregate root dari aplikasi yang berfungsi untuk mengenkapsulasi business rules/invariant dan menjaga statenya selalu valid.

### Direktori `internal/domain/events`

Direktori ini berisikan event yang dipublish oleh aggregate root.

### Direktori `internal/domain/valueobjects`

Direktori ini berisikan value object yang digunakan untuk mengenkapsulasi business rules/invariant dan menjaga statenya selalu valid.
