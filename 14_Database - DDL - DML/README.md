# **Rangkuman Materi Database**

## **1. Database**

- ### Kumpulan data yang terorganisasi dan terstruktur.
- ### 2 konsep dalam database

  - #### Schema

    - Sebuah struktur yang mengatur bagaimana data disimpan dalam database yang mencakup definisi data, hubungan antar data, dan batasan akses data.
    - Seperti dalam sebuah tabel karyawan. Schema tabelnya akan mendefinisikan field seperti employee_ID dan employee_Name beserta tipe data yang sesuai seperti string of 10 digits atau string of 45 characters

  - #### Model

    - Struktur keseluruhan dari sebuah database. Model database menggambarkan bagaimana data disimpan, bagaimana data terhubung, dan bagaimana data diakses.
    - Seperti model relasional yang menggambarkan data dalam bentuk tabel yang terhubung melalui relasi antar tabel

- ### Hubungan antara data dalam sebuah database dikenal sebagai relationship.
  - #### One-to-One Relationship
    - Hubungan antara dua entitas yang saling terkait satu sama lain
    - Misalnya, tabel karyawan dan tabel gaji. Setiap karyawan memiliki satu gaji dan setiap gaji hanya dimiliki oleh satu karyawan
  - #### One-to-Many Relationship
    - Hubungan antara dua entitas dimana satu entitas memiliki banyak entitas lainnya
    - Misalnya, tabel karyawan dan tabel gaji. Setiap karyawan memiliki banyak gaji dan setiap gaji hanya dimiliki oleh satu karyawan
  - #### Many-to-Many Relationship
    - Hubungan antara dua entitas dimana satu entitas memiliki banyak entitas lainnya dan satu entitas lainnya juga memiliki banyak entitas lainnya
    - Misalnya, tabel karyawan dan tabel proyek. Setiap karyawan bisa bekerja di banyak proyek dan setiap proyek juga memiliki banyak karyawan

## **2. Sistem Manajemen Basis Data (DBMS)**

- ### Aplikasi untuk mengelola dan mengakses database, seperti workbench pada SQL dan MongoDB Compass pada NoSQL.

- ### Sistem Manajemen Basis Data Relasional (RDBMS)

  - #### Jenis DBMS yang menggunakan model relasional dalam mengatur dan mengakses data dalam database, termasuk jenis SQL.

- ### Jenis database yang umum digunakan adalah SQL dan NoSQL.

  - #### SQL (Structured Query Language)
    - Database yang terstruktur dan memiliki relasi antar tabel
    - Contoh: MySQL, PostgreSQL, Oracle, Microsoft SQL Server
  - #### NoSQL (Not Only SQL)
    - Database yang tidak terstruktur dan menggunakan model dokumen, grafik, atau key-value dalam mengatur dan mengakses data dalam database.
    - Contoh: MongoDB, Cassandra, Redis, CouchDB

## **3. Perintah SQL**

- ### Jenis perintah SQL

  - #### DDL (Data Definition Language)

    - Digunakan untuk mendefinisikan struktur database, seperti membuat tabel, mengubah tabel, menghapus tabel, dan lain-lain
    - Contoh: CREATE, USE, DROP, RENAME, ALTER, dll

  - #### DML (Data Manipulation Language)

    - Digunakan untuk memanipulasi data dalam tabel, seperti menambahkan data, mengubah data, menghapus data, dan lain-lain
    - Contoh: INSERT, SELECT, UPDATE, DELETE, LIKE, IN, BETWEEN, AND, OR, dll

  - #### DCL (Data Control Language)

    - Digunakan untuk mengontrol akses data, seperti memberikan hak akses, menghapus hak akses, dan lain-lain
    - Contoh: GRANT, REVOKE, dll

  - ### DTCL (Data Transaction Control Language)

    - Digunakan untuk mengontrol transaksi, seperti mengatur transaksi, menggagalkan transaksi, dan lain-lain
    - Contoh: COMMIT, ROLLBACK, SAVEPOINT, dll
