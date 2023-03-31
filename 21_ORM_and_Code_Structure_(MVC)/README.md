# **Rangkuman Materi ORM dan Code Structure (MVC)**

## **ORM (Object Relational Mapping)**

- ### Pengertian
  Teknik pemrograman yang digunakan untuk memetakan antara objek dalam sebuah bahasa pemrograman dengan tabel dalam sebuah basis data relasional.
- ### Keuntungan

  - Mengurangi repetisi query
  - Secara otomatis **fetch data** dari database dan mengubahnya menjadi objek
  - Cara termudah if ingin **screening data** sebelum menyimpannya ke database
  - Mempunyai fiture cache query secara otomatis

- ### Kekurangan

  - Membutuhkan operasi yang lebih yang dapat memberatkan di bagian operasi database
  - Semakin complex querynya maka semakin panjang penulisan ORM-nya (> 10 tabel join)
  - Terkadang ORM tidak mendukung fitur-fitur yang ada di SQL atau belum tersedia di ORM yang digunakan

- ### Database migration

  Cara mengupdate database tanpa harus menghapus database yang sudah ada

  - Keuntungan:
    - Update database lebih mudah
    - Rollback database lebih mudah
    - Mentrack perubahan struktur database
    - History database dibuat dalam sebuah kode
    - Selalu compatible dengan versi yang sebelumnya atau yang baru

## **GORM (Go Object Relational Mapping)**

- ### Cara install
  - `go get -u gorm.io/gorm`
  - `go get -u gorm.io/driver/mysql`
- ### Cara koneksi

  ```go
    package main

    import (
      "gorm.io/driver/mysql"
      "gorm.io/gorm"
    )
    func main(){
      dsn := "root:@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
      db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    }
  ```

- ### Lebih detail tentang GORM pada link berikut
  - https://gorm.io/

## **Code Structure (MVC)**

- ### Pengertian

  Model View Controller adalah sebuah pola arsitektur yang memisahkan aplikasi menjadi tiga bagian utama, yaitu model, view, dan controller. Model berisi data dan logika bisnis, view berisi tampilan, dan controller berisi logika untuk mengatur data dan tampilan.

- ### Struktur MVC

  - Model
    - Berisi data dan logika bisnis
    - Biasanya berupa struct
    - Biasanya berada di folder `models`
  - View
    - Berisi tampilan
    - Biasanya berupa file HTML
    - Biasanya berada di folder `views`
  - Controller
    - Berisi logika untuk mengatur data dan tampilan
    - Biasanya berupa file `.go`
    - Biasanya berada di folder `controllers`
