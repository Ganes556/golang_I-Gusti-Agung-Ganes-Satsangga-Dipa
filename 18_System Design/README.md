# **Rangkuman Materi System Design**

## **1. Diagram**

- Representasi visual dari suatu sistem yang berisi komponen-komponen sistem dan hubungan antar komponen tersebut.
- Software design yang biasa digunakan:
  - **Flowchart** : Representasi grafis dari sebuah proses atau algoritma
  - **Use Case** : Deskripsi dari perilaku sebuah sistem dari sudut pandang pengguna
  - **ERD (Entity Relationship Diagram)** : Diagram yang menunjukkan hubungan antara entitas dalam sebuah database
  - **HLA (High Level Architecture)** : Desain arsitektur tingkat atas yang menguraikan komponen dan interaksi dari sebuah sistem

## **2. Distributed System**

- Sistem yang terdiri dari beberapa komputer yang saling berkomunikasi dan saling bekerja sama untuk mencapai tujuan tertentu.
- Karakteristik:
  - scalability : kemampuan untuk menambahkan komponen baru ke sistem
    - vertical scaling : menambahkan komponen baru dengan spesifikasi yang lebih tinggi (lebih lama)
    - horizontal scaling : menambahkan komponen baru dengan spesifikasi yang sama (lebih cepat)
  - realibility : keandalan dari sistem dalam menjalankan tugasnya
  - availability : ketersediaan dari sistem dalam menjalankan tugasnya (tidak down)
  - efiency : kecepatan dalam menjalankan tugasnya (response time)
  - serviceability or manageability : harus mudah untuk di maintain dan di manage (good clean code)

## **3. Job Queue, Load Balancer, Monolitic dan Microservice**

- Job Queue : sistem yang memproses tugas-tugas yang masuk dan menyimpannya dalam antrian (queue) untuk di proses oleh framework atau sistem lainnya (worker)
- Load Balancer : sistem yang membagi tugas-tugas yang masuk ke beberapa komputer (server) secara merata
- Microservice : sistem yang terdiri dari beberapa komponen yang saling berkomunikasi dan saling bekerja sama untuk mencapai tujuan tertentu
- Monolitic : sistem yang terdiri dari satu komponen utama yang menjalankan semua tugas

## **4. SQL vs NoSQL**

- SQL (Structured Query Language) : bahasa yang digunakan untuk mengakses dan memanipulasi data dalam database relasional
  - Kelebihan :
    - Dirancang untuk segala keperluan
    - Memiliki standar yang jelas
    - Memiliki banyak tools (administrasi, reporting, framework, dll)
  - Prinsip :
    - A : Atomicity, transaksi terjadi semua atau tidak sama sekali
    - C : Consistency, data tertulis merupakan data valid yang ditentukan berdasarkan aturan tertentu
    - I : Isolation, pada saat terjadi request yang bersamaan (concurrent), memastikan bahwa transaksi dieksekusi seperti dijalankan secara sekuensial
    - D : Durability, jaminan bahwa transaksi yang telah tersimpan, tetap tersimpan
      - Soft delete : data yang dihapus tidak dihapus secara permanen, tetapi hanya diubah statusnya menjadi tidak aktif
- NoSQL (Not Only SQL) : database yang tidak menggunakan SQL
  - Kelebihan : skema lebih setikit, cepat di kembangkan dan di maintain, tidak perlu melakukan join, tidak perlu melakukan normalisasi
  - Kapan diperlukan: membutuhkan skema fleksibel, ACID tidak diperlukan, data logging (json), data sementara (chace), dll
  - Kapan tidak direkomendasikan: ketika data"nya penting diharuskan memiliki struktur yang jelas seperti: data finansial, data transaksi, business critical, ACID - compliant

## **5. Cache dan Indexing**

- Cache : data yang disimpan di memori untuk mengurangi waktu akses ke data yang sama
- Redudancy dan Replication

  - Redudancy : data yang disimpan di beberapa tempat
  - Replication : data yang disimpan di beberapa tempat dan diupdate secara sinkron

- Indexing : cara untuk mengoptimasi pencarian data dalam database dengan cara menambahkan kolom tambahan index agar dapat
  - Lebih cepat untuk mencari data
  - Lebih lambat untuk menambahkan data
