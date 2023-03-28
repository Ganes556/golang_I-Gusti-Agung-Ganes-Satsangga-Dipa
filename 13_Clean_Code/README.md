# **Rangkuman Materi Clean Code**

## **Pengertian**

- Clean code adalah kode yang mudah dibaca, mudah dimengerti, dan mudah diperbaiki.

## **Manfaat**

- Mempermudah dalam berkolaborasi dengan programmer lain.
- Mempercepat proses development.

## **Ciri-ciri**

- **Mudah dipahami**,
  - Seperti, penamaan variabel **locations** untuk array yang berisi kumpulan lokasi dan **location** untuk variabel yang berisi satu lokasi.
- **Mudah dieja dan dicari**
  - Seperti, penamaan variabel **fullName** untuk nama lengkap. sehingga mudah dieja dan dicari.
- **Singkat namun mendeskripsikan konteks**
  - Seperti, mengambil field merek yang berada pada array index ke 1, maka penamaan variabelnya **brand** = `car[1]`
- **Konsisten**
  - Seperti, penamaan variabel **konstanta** menggunakan huruf kapital semua dan yang lainnya **camelCase**
- **Hindari penambahan konteks yang tidak perlu**
  - Seperti, penamaan variabel **fullName** yang sudah pasti merupakan string. sehingga tidak perlu menambahkan konteks **fullNameString**.
- **Komentar**
  - Komentar yang menjelaskan logika atau apa yang dilakukan oleh suatu kode. bukan menjelaskan kode yang ditulis.
- **Good Function**
  - Jangan teralu banyak argumen
  - Berikan komentar pada argumen return jika perlu
- **Gunakan Konvensi**
  - Mengikuti style code yang sudah ada di perusahaan, seperti airbnb, google, dll.
- **Formating**
  - Lebar baris code 80-120 karakter
  - Satu class 300-500 baris
  - Baris code yang berhubungan saling berdekatan
  - Dekatkan fungsi dengan pemanggilnya
  - Deklarasi variabel berdekatan dengan penggunanya
  - Perhatikan indentasi
  - Menggunakan extension **prettier** atau **formatter** untuk format code secara otomatis

## **Principle**

- **KISS**
  - Keep It Simple, Stupid
  - Kode harus mudah dipahami dan mudah diperbaiki.
- **DRY**
  - Don't Repeat Yourself
  - Hindari duplikasi kode.
- **Refactoring**
  - Refactoring adalah proses mengubah kode menjadi clean code tanpa mengubah fungsionalitas dari suatu kode.
