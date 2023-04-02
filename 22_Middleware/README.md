# **Rangkuman Middleware**

## Middleware

- Sebuah fungsi yang dapat dijalankan sebelum atau sesudah sebuah request dijalankan.

## Middleware pada Echo

- ### Middleware dalam echo dibagi menjadi 2:

  1. Echo#Pre() : Middleware yang dijalankan sebelum router memproses request
  2. Echo#Use() : Middleware yang dijalankan sesudah router memproses request dan memiliki akses penuh terhadap echo.Context API

- ### Contoh middleware pada Echo

  - Middleware Logger, middleware yang digunakan untuk menampilkan log request yang masuk ke server.
  - Middleware BasicAuth, middleware yang digunakan untuk melakukan autentikasi terhadap request yang masuk ke server.
  - Middleware JWT, middleware yang digunakan untuk melakukan autentikasi terhadap request yang masuk ke server dengan menggunakan JWT.
  - [etc.](https://echo.labstack.com/middleware/)
