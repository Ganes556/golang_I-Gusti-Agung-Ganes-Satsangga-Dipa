# **Rangkuman Materi Concurrent Programming**

## Sequential Programming

- Intruksi yang dieksekusi secara berurutan

## Parallel Programming

- Intruksi yang dieksekusi secara bersamaan (harus ada multi core)

## Concurrent Programming

- Intruksi yang diselesaikan dalam waktu yang sama dan berjalan dengan independent (tidak harus ada multi core)

## Feature Concurrent pada Golang

- ### Goroutine

  - Sebagai proses concurrent
  - Sebuah fungsi atau method yang dapat dijalankan secara concurrent (indepedent) dengan fungsi atau method lain

- ### Channel

  - Sebagai tempat berbagi data antar goroutine

- ### Select

  - Sebagai mengontrol jalannya concurrent dari beberapa channel

## Race Condition

- Terjadi ketika terdapat 2 thread yang mengakses memory secara bersamaan, seperti membaca dan menulis data pada memory yang sama
