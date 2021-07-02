# Belajar Go GraphQL
Belajar GraphQL dengan bahasa pemrograman Go.

# Daftar Isi
* [Fitur](#fitur)
* [Konsep Arsitektural](#konsep-arsitektural)
* [Stack Teknologi](#stack-teknologi)
* [Utilitas Pihak Ketiga](#utilitas-pihak-ketiga)
* [Menjalankan Projek](#menjalankan-projek)
  * [Persiapan](#persiapan)
  * [Mode Development](#mode-development)
  * [Mode Production](#mode-production)

# Fitur
Projek ini sudah siap dijalankan dan dapat dimodifikasi sesuai kebutuhan masing-masing. Adapun kemampuan projek ini:
1. Dapat menjadi backend dari GraphQL.
2. Dilengkapi Graphiql.
3. Dapat menjadi service dalam rangkaian microservices.
4. Support berbagai koneksi database, seperti MySQL, PostgreSQL, SQLite, dan SQL Server.
5. Support migrasi database dalam bentuk file SQL.
6. Support [dependency injection](https://en.wikipedia.org/wiki/Dependency_injection).

# Konsep Arsitektural
Struktur proyek ini mengikuti konsep-konsep arsitektural seperti:
- DRY (Don't Repeat Yourself)
- SOLID Principle
- Clean Architecture

# Stack Teknologi
Nama | Kegunaan
-|-
[Go](https://golang.org) | Bahasa Pemrograman
[Fiber](https://docs.gofiber.io) | HTTP Framework
[PostgreSQL](https://www.postgresql.org) | Database Management System
[Gorm](https://gorm.io/index.html) | Object-Relational Mapper
[Jwt](https://jwt.io) | Token untuk Otentikasi
[GraphQL Go](https://github.com/graphql-go/graphql) | Engine GraphQL

# Utilitas Pihak Ketiga
Nama | Kegunaan
-|-
[Validator](https://github.com/go-playground/validator) | Validasi data
[Soda CLI](https://gobuffalo.io/en/docs/db/toolbox#from-a-release-archive) | Migrasi database
[Swag](https://github.com/swaggo/swag) | Otomatis generate dokumentasi REST API (swagger)
[GNU/Make](https://www.gnu.org/software/make/) | Build tool

# Menjalankan Projek
## Persiapan
Agar projek dapat dijalankan, diperlukan beberapa hal untuk dipersiapkan, antara lain:
1. Menjalankan migration dengan menyesuaikan akses database.
```bash
./.bin/soda m -c db/database.yml -p db/migrations
```

## Mode Development
1. Clone projek ke komputer lokal, dengan perintah:
  ```bash
  git clone https://github.com/hadihammurabi/belajar-go-graphql
  ```
2. Menjalankan projek dengan mode development (pengembangan) dengan perintah:
  ```bash
  make dev
  ```

## Mode Production
1. Clone projek ke komputer lokal, dengan perintah:
  ```bash
  git clone https://github.com/hadihammurabi/belajar-go-graphql
  ```
2. Menjalankan projek siap guna dengan perintah:
  ```bash
  make && ./main
  ```
  
