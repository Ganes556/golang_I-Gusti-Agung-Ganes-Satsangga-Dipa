# **Rangkuman Clean dan Hexagonal Architecture**

## Hexagonal Architecture

- ### Pengertian

  Hexagonal Architecture adalah sebuah arsitektur yang memisahkan aplikasi menjadi dua bagian utama, yaitu domain dan infrastructure. Domain berisi data dan logika bisnis, sedangkan infrastructure berisi tampilan dan logika untuk mengatur data dan tampilan.

- ### Struktur Hexagonal Architecture

  - API interface, Berisi hal yang berhubungan dengan API, seperti rest controller/framework/CLI
  - Domain service, Berisi data dan logika bisnis, seperti model, repository, dan service
  - SPI interface, Berisi hal yang berhubungan dengan database, seperti mysql, redis, third party, dll

## Clean Architecture

- ### Pengertian

  Clean Architecture adalah sebuah arsitektur yang memisahkan aplikasi menjadi beberapa layer, dimana setiap layer memiliki tanggung jawab yang berbeda-beda. Layer yang paling luar adalah layer yang paling dekat dengan user, sedangkan layer yang paling dalam adalah layer yang paling dekat dengan database.

- ### Constraint

  - Independent of Frameworks
  - Testable
  - Independent of UI
  - Independent of Database
  - Independent of any external agency

- ### Keuntungan

  - Struktur yang standar, sehingga mempermudah dalam melakukan integrasi dengan suatu project
  - Mempercepat dalam melakukan development (untuk jangka panjang)
  - Mempermudah dalam melakukan mocking denpendencies di unit test
  - Mempermudah dalam melakukan perubahan pada suatu layer tanpa harus melakukan perubahan pada layer lainnya

- ### Layer

  - (optional) Entities: Berisi objek bisnis
  - Use Cases/Domain/services: Berisi logika bisnis
  - Controller/Presentation: Menyajikan data dan menghandle interaksi dengan user
  - Drivers/Data: Mengelola data aplikasi seperti menerima data dari network, mengelola data cache

## Ubiquitous Language

- Sebuah istilah yang digunakan untuk mendeskripsikan bahasa yang digunakan oleh semua orang yang terlibat dalam suatu project. Ubiquitous Language ini biasanya berupa kata-kata yang digunakan untuk mendeskripsikan suatu hal, seperti kata-kata yang digunakan untuk mendeskripsikan suatu objek, kata-kata yang digunakan untuk mendeskripsikan suatu fungsi, kata-kata yang digunakan untuk mendeskripsikan suatu proses, dll.
