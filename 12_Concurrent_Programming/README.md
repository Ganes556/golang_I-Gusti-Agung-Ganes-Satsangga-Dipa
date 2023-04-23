# **Rangkuman Materi Concurrent Programming**

## **Beberapa Jenis Programming**

- ### Sequential Programming

  - Intruksi yang dieksekusi secara berurutan

- ### Parallel Programming

  - Intruksi yang dieksekusi secara bersamaan (harus ada multi core)

- ### Concurrent Programming

  - Intruksi yang diselesaikan dalam waktu yang sama dan berjalan dengan independent (tidak harus ada multi core)

## Feature Concurrent pada Golang

- ### Goroutine

  - Sebagai proses concurrent
  - Sebuah fungsi atau method yang dapat dijalankan secara concurrent (indepedent) dengan fungsi atau method lainnya

- ### Channel

  - Sebagai tempat berbagi data antar goroutine
  - Terdapat 2 jenis,
    1. **unbuffered channel** : tidak menampung data, hanya mengirim data
    2. **buffered channel** : memiliki kapasitas untuk menampung data, sebelum dikirim

- ### Select

  - Sebagai mengontrol jalannya concurrent dari beberapa channel

## **Race Condition**

- Salah satu kelemanan dari concurrent programming
- Terjadi ketika terdapat 2 thread yang mengakses memory secara bersamaan, seperti membaca dan menulis data pada memory yang sama
- Solusi dari race condition dengan melakukan bloking terhadap goroutinenya. Dapat menggunakan **mutex**, **wait group**, **channel**
- Cara melakukan pengecekan race condition dengan menjalankan golang dengan flag `-race`, contoh: `go run -race nama_file.go`
