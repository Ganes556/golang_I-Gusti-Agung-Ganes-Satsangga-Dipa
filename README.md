# **Rangkuman Seluruh Materi Become a Master of Golang Programming** <!-- omit from toc -->

## **Daftar Isi** <!-- omit from toc -->

- [**Indroduction to Algorithm and Golang**](#indroduction-to-algorithm-and-golang)
  - [**Pengertian Algoritma dan Karakteristiknya**](#pengertian-algoritma-dan-karakteristiknya)
  - [**Konsep Algoritma Dasar**](#konsep-algoritma-dasar)
  - [**Cara Mempresentasikan Algoritma**](#cara-mempresentasikan-algoritma)
- [**Basic Programming**](#basic-programming)
  - [**Tentang Golang**](#tentang-golang)
  - [**Cara Instal Golang di VS Code**](#cara-instal-golang-di-vs-code)
  - [**Perintah Terminal dan Penjelasan Sintak dasar di Golang**](#perintah-terminal-dan-penjelasan-sintak-dasar-di-golang)
- [**Version Control and Branch Management (Git)**](#version-control-and-branch-management-git)
  - [**Version Control System (VCS)**](#version-control-system-vcs)
  - [**Sejarah Singkat Version Control**](#sejarah-singkat-version-control)
  - [**3 Dasar Staging Area**](#3-dasar-staging-area)
  - [**Command git**](#command-git)
  - [**Tips untuk mengoptimasi alur kerja dengan github**](#tips-untuk-mengoptimasi-alur-kerja-dengan-github)
- [**Time Complexity \& Space Complexity**](#time-complexity--space-complexity)
  - [**Time Complexity**](#time-complexity)
  - [**Space Complexity**](#space-complexity)
  - [**Beberapa Jenis Time Complexity**](#beberapa-jenis-time-complexity)
  - [**Beberapa Jenis Space Complexity**](#beberapa-jenis-space-complexity)
- [**Data Structure**](#data-structure)
  - [**Struktur Data**](#struktur-data)
  - [**Array**](#array)
  - [**Slice**](#slice)
  - [**Map**](#map)
  - [**Perbedaan Array, Slice dan Map**](#perbedaan-array-slice-dan-map)
  - [**Fungsi**](#fungsi)
- [**String, Advance Function, Pointer, Method, Struct and Interface**](#string-advance-function-pointer-method-struct-and-interface)
  - [**Beberapa hal yang dapat dilakukan pada string golang**](#beberapa-hal-yang-dapat-dilakukan-pada-string-golang)
  - [**Advance Function**](#advance-function)
  - [**Pointer**](#pointer)
  - [**Struct dan Interface**](#struct-dan-interface)
  - [**Method**](#method)
  - [**Error Handling**](#error-handling)
- [**Recursive, Number Theory, Searching, Sorting**](#recursive-number-theory-searching-sorting)
  - [**Recursive**](#recursive)
  - [**Number Theory**](#number-theory)
  - [**Searching**](#searching)
  - [**Sorting**](#sorting)
- [**Problem Solving Paradigm - Brute Force Greedy and Dynamic Programming**](#problem-solving-paradigm---brute-force-greedy-and-dynamic-programming)
  - [**Problem Solving Paradigm**](#problem-solving-paradigm)
- [**Concurrent Programming**](#concurrent-programming)
  - [**Beberapa Jenis Programming**](#beberapa-jenis-programming)
  - [**Feature Concurrent pada Golang**](#feature-concurrent-pada-golang)
  - [**Race Condition**](#race-condition)
- [**Clean Code**](#clean-code)
  - [**Pengertian**](#pengertian)
  - [**Manfaat**](#manfaat)
  - [**Ciri-ciri**](#ciri-ciri)
  - [**Principle**](#principle)
- [**Database**](#database)
  - [**Database**](#database-1)
  - [**Sistem Manajemen Basis Data (DBMS)**](#sistem-manajemen-basis-data-dbms)
  - [**Perintah SQL**](#perintah-sql)
- [**Comunication Skill**](#comunication-skill)
  - [**Komunikasi Efektif**](#komunikasi-efektif)
  - [**Fungsi Komunikasi Efektif**](#fungsi-komunikasi-efektif)
  - [**5W Element Komunikasi**](#5w-element-komunikasi)
  - [**7C Komunikasi**](#7c-komunikasi)
  - [**Jenis Komunikasi**](#jenis-komunikasi)
- [**Join, Union, Agregasi, Subquery, Function (DBMS)**](#join-union-agregasi-subquery-function-dbms)
  - [**Join**](#join)
  - [**Union**](#union)
  - [**Fungsi Agregasi**](#fungsi-agregasi)
  - [**Subquery**](#subquery)
  - [**Trigger**](#trigger)
  - [**Function**](#function)
- [**Configuration Management and CLI**](#configuration-management-and-cli)
  - [**Command Line Interface (CLI)**](#command-line-interface-cli)
  - [**Command yang sering digunakan**](#command-yang-sering-digunakan)
  - [**Shell Script**](#shell-script)
- [**System Design**](#system-design)
  - [**Diagram**](#diagram)
  - [**Distributed System**](#distributed-system)
  - [**Job Queue, Load Balancer, Monolitic dan Microservice**](#job-queue-load-balancer-monolitic-dan-microservice)
  - [**SQL vs NoSQL**](#sql-vs-nosql)
  - [**Cache dan Indexing**](#cache-dan-indexing)
- [**Introduction to RESTful API**](#introduction-to-restful-api)
  - [**API (Application Programming Interface)**](#api-application-programming-interface)
  - [**Design REST API**](#design-rest-api)
  - [**Cara Implementasi REST API dengan Go**](#cara-implementasi-rest-api-dengan-go)
  - [**Reference**](#reference)
- [**Introduction Echo Golang**](#introduction-echo-golang)
  - [**Framework Echo**](#framework-echo)
  - [**Basic Echo**](#basic-echo)
- [**ORM dan Code Structure (MVC)**](#orm-dan-code-structure-mvc)
  - [**ORM (Object Relational Mapping)**](#orm-object-relational-mapping)
  - [**GORM (Go Object Relational Mapping)**](#gorm-go-object-relational-mapping)
  - [**Code Structure (MVC)**](#code-structure-mvc)
- [**Middleware**](#middleware)
  - [**Middleware**](#middleware-1)
  - [**Middleware pada Echo**](#middleware-pada-echo)
- [**Unit Testing**](#unit-testing)
  - [**Unit Testing**](#unit-testing-1)
  - [**Testing Framework**](#testing-framework)
  - [**Implementasi Unit Testing Pada Golang**](#implementasi-unit-testing-pada-golang)
- [**Clean dan Hexagonal Architecture**](#clean-dan-hexagonal-architecture)
  - [**Hexagonal Architecture**](#hexagonal-architecture)
  - [**Clean Architecture**](#clean-architecture)
  - [**Ubiquitous Language**](#ubiquitous-language)
- [**Proffessional Skill 2 Building Rangkuman and CV**](#proffessional-skill-2-building-rangkuman-and-cv)
  - [**Kesalahan Umum**](#kesalahan-umum)
  - [**Tips Terpenting**](#tips-terpenting)
  - [**Konten pada CV**](#konten-pada-cv)
- [**Docker**](#docker)
  - [**Docker**](#docker-1)
  - [**Strategi Deployment**](#strategi-deployment)
- [**Compute Service**](#compute-service)
  - [**Simple Dev-Ops Cycle**](#simple-dev-ops-cycle)
  - [**Compute Service**](#compute-service-1)
  - [**AWS RDS (Relational Database Service)**](#aws-rds-relational-database-service)
  - [**Link Deploy dengan AWS EC2 dan RDS**](#link-deploy-dengan-aws-ec2-dan-rds)
- [**CI/CD**](#cicd)
  - [**CI (Continuous Integration)**](#ci-continuous-integration)
  - [**CD (Continuous Delivery/Deployment)**](#cd-continuous-deliverydeployment)
  - [**Tools CI/CD**](#tools-cicd)

# **Indroduction to Algorithm and Golang**

## **Pengertian Algoritma dan Karakteristiknya**

Algoritma adalah prosedur komputasi yang didefinisikan dengan baik, serta harus memiliki batasan awalan dan akhiran, setiap tahapan instruksinya harus terdefinisi dengan baik, dan harus efektif dan efisien

## **Konsep Algoritma Dasar**

1. **Sequential** (Urutan)
   : Urutan dalam menyusun langkah-langkah untuk menyelesaikan sebuah masalah
2. **Branching** (Percabangan)
   : Menentukan sebuah alur dari sebuah program dari suatu masalah
3. **Looping** (Perulangan)
   : Melakukan sebuah aksi secara berulang sampai dengan kondisi tertentu terpenuhi

## **Cara Mempresentasikan Algoritma**

1. **Pseudo Code**
   : Deskripsi bahasa yang sederhana untuk menuangkan Ide atau algoritma itu sendiri
2. **Flowchart**
   : Bagan dengan urutan tertentu yang saling berhubungan antar proses secara mendetail

# **Basic Programming**

## **Tentang Golang**

- Bahasa pemrograman open-source yang dikembangkan oleh Google pada tahun 2007.
- Bertujuan untuk memudahkan proses development, sederhana, handal, dan efisien.
- Dirilis ke publik di tahun 2009.
- Golang dikembangkan oleh tim yang terdiri dari Rob Pike, Ken Thompson, Robert Griesemer, Ian Lenz Taylor dan Russ Cox.
- Sering digunakan untuk membuat aplikasi e-commerce, musik player, sosial media, FPI, game engine, dan lain-lain.

## **Cara Instal Golang di VS Code**

- Dapat dengan cara Mengunduh [VS Code](https://code.visualstudio.com/download) dan menginstalnya, mengunduh [Golang](https://go.dev/dl/) dan menginstalnya di. Kemudian, setting environment path golang yang di sesuaikan pada os masing-masing.

## **Perintah Terminal dan Penjelasan Sintak dasar di Golang**

- Perintah terminal golang

  - go run `<nama_file>.go`
  - go build `<nama_file>.go`
  - go install `<nama_file>.go`
  - go get `<nama_file>.go`

- Sintak dasar golang

  - Package "fmt" yang merupakan package standar untuk melakukan input dan output. Seperti,
    - Output
      - fmt.Println()
      - fmt.Printf()
      - fmt.Print()
    - Input
      - fmt.Scanln()
    - Format Verb
      - %T, %v, %s, %q, dll.
  - Variabel, tempat untuk menyimpan data yang memiliki nama, sehingga nama tersebut yang akan dipanggil pada saat membutuhkan datanya

    - Tipe-tipe variabel pada Golang yakni, **Boolean** : true / false, **String** , **Number** : int, float, complex
      - Cara deklarasi variabel dengan cara,
        - var `<nama_variabel> <tipe_data>`
        - var `<nama_variabel> <tipe_data> = <nilai>`
        - var `<nama_variabel> <tipe_data>`
        - var `<nama_variabel> = <nilai>`
        - `<nama_variabel> := <nilai>`
    - Zero value tiap variabel, dimana **Boolean** : false, **Number** : (float : 0.0), (int : 0), (complex : 0i), **String** : ""
    - Constants variabel yang nilainya tidak dapat diubah dan ada 2 yakni Single Constant dan Multiple Constant
    - Operand, nilai yang akan dioperasikan oleh operator
    - Pperator, melakukan operasi pada variabel, literal, dan ekspresi
    - Expression kombinasi antara konstanta, variabel, fungsi, dan operator yang mengembalikan sebuah nilai.

  - Control structure yang merupakan struktur yang digunakan untuk mengatur alur program yang dimana pada golang hanya terdapat **if**, **switch** dan **for**

# **Version Control and Branch Management (Git)**

## **Version Control System (VCS)**

- ### Version Control System (VCS) adalah sebuah sistem yang digunakan untuk mengelola perubahan pada sebuah file atau sekumpulan file. VCS memungkinkan kita untuk melihat perubahan apa saja yang terjadi pada file, kapan perubahan tersebut terjadi, siapa yang melakukan perubahan tersebut, dan mengembalikan file ke versi sebelumnya.

## **Sejarah Singkat Version Control**

1.  Local Version Control (Single User)

    - Kelebihan :
      - Mudah digunakan
      - Tidak memerlukan koneksi internet
    - Kekurangan :
      - Tidak ada backup
      - Jika terjadi kerusakan pada database, maka tidak bisa melakukan perubahan pada file
    - Contoh:
      - SCCS - 1972 hanya untuk sistem operasi Unix
      - RCS - 1982 cross platform, hanya untuk text

2.  Centralize (Client-Server)
    - Kelebihan :
      - Mudah digunakan, tidak memerlukan koneksi internet
    - Kekurangan :
      - Jika server down, maka tidak bisa melakukan perubahan pada file
    - Contoh :
      - CVS - 1995 hanya untuk file
      - Perforce - 1995
      - Subversion - 2000 - melacak perubahan struktur direktori
      - Microsoft Team Foundation Server - 2005
3.  Distributed (Peer-to-Peer)
    - Kelebihan :
      - Jika server down, tetap bisa melakukan perubahan pada file
    - Kekurangan :
      - Memerlukan koneksi internet
    - Contoh :
      - Git - 2005
      - Mercurial - 2005
      - Bazaar - 2005

## **3 Dasar Staging Area**

1. Working Directory : Tempat kerja kita saat ini
2. Staging Area : Tempat untuk menyimpan perubahan yang akan di commit
3. Repository : Tempat penyimpanan file yang sudah di commit

## **Command git**

Beberapa command git yang dipelajari pada materi ini yakni,

1. git init : untuk membuat repository di local
2. git add : untuk menambahkan file ke stagging area
3. git commit : untuk menyimpan file ke repository
4. git push : untuk mengirimkan file ke repository
5. git status : untuk melihat status dari file
6. git diff : untuk melihat perubahaan apa saja yang terjadi pada file
7. git stash : untuk menjadi penyimpanan sementara atau menyembunyikan perubahan yang belum di commit dan akan masuk ke stash area
8. git log : untuk melihat history dari commit
9. git reset : untuk mengembalikan file ke versi sebelumnya
   - git reset --soft : perubahan tidak dicommit, tetapi tetap ada di staging area
   - git reset --hard : perubahannya tidak dicommit dan tidak ada di staging area
10. git fetch : untuk mengambil perubahan data git di remote repository dan langsung di simpan di git local tetapi file/perubahan belum ada (tidak di merge)
11. git pull : untuk mengambil file dari repository dan langsung di merge
12. git branch : untuk melihat branch apa saja yang ada di repository atau bisa untuk CURD branch bila di tambahkan opsi setelah branch
13. git checkout : untuk berpindah branch atau bisa juga untuk pindah sekaligus membuat branch baru dengan menambah opsi `-b`
14. git merge : untuk menggabungkan branch yang berbeda

## **Tips untuk mengoptimasi alur kerja dengan github**

1. Hindari perubahaan di branch master, branch master hanya digunakan untuk menyimpan versi yang sudah stabil
2. Hindari melakukan perubahan secara langsung di branch development, sebaiknya buat branch baru untuk setiap fitur yang akan dikerjakan
3. Setiap branch fitur yang telah selesai minta pull request ke branch development
4. Setiap branch development yang telah selesai minta pull request ke branch master

# **Time Complexity & Space Complexity**

## **Time Complexity**

- Ukuran waktu yang dibutuhkan oleh sebuah algoritma untuk menyelesaikan sebuah masalah dengan input tertentu.
- Dinyatakan dalam notasi Big O dan dapat dinyatakan sebagai O(1), O(log n), O(n), O(n^2), dll.
- O(1) merupakan kompleksitas waktu tercepat, sedangkan O(n^2) merupakan kompleksitas waktu terlama.
- Menyederhanakan proses perhitungan

## **Space Complexity**

- Sebuah metode untuk mengukur seberapa banyak memori yang digunakan oleh sebuah algoritma.
- Biasanya diukur dalam satuan bit atau byte, dan meliputi ruang yang digunakan untuk menyimpan variabel, data struktur, dan tumpukan (stack) selama eksekusi algoritma.
- Space Complexity juga dinyatakan sebagai O(1), O(n), O(n^2), atau O(log n), tergantung pada seberapa banyak memori yang dibutuhkan oleh algoritma seiring dengan peningkatan ukuran inputnya

## **Beberapa Jenis Time Complexity**

1. Constant Time Complexity - O(1)

   - Algoritma yang dimana, jumlah input yang diberikan tidak mempengaruhi waktu proses (runtime) dari algoritma tersebut. Sehingga membentuk notasi O(1).
   - Contoh kasus: Akses elemen dalam array dengan indeks tertentu atau mencari panjang dari sebuah string.

2. Linear Time Complexity - O(n)

   - Algoritma yang dimana, ketika runtime dari fungsi akan berbanding lurus dengan jumlah input yang diberikan. Sehingga membentuk notasi O(n).
   - Contoh kasus: Mencari elemen tertentu dalam array atau string,menghitung total dari seluruh elemen dalam array.

3. Quadratic Time Complexity - O(n^2)

   - Algoritma yang dimana, ketika runtime dari fungsi adalah sebesar n^2, dimana n adalah jumlah input dari fungsi tersebut. Hal tersebut bisa terjadi karena menjalankan fungsi linear didalam fungsi linear (n\*n). Sehingga membentuk notasi O(n^2).
   - Contoh kasus: Mencari sebuah kombinasi atau permutasi, melakukan pengurutan menggunakan metode bubble sort

4. Logarithmic Time Complexity - O(log n)

   - Algoritma yang dimana, ketika memberikan input sebesar n terhadap sebuah fungsi, jumlah tahapan yang dilakukan oleh fungsi tersebut berkurang berdasarkan suatu faktor. Sehingga membentuk notasi O(n).
   - Contoh kasus: Mencari nilai tertentu dalam sebuah array yang diurutkan menggunakan binary search, atau menghitung hasil operasi exponentiation dengan cepat menggunakan metode divide and conquer.

## **Beberapa Jenis Space Complexity**

1. Constant Space Complexity - O(1)

   - Algoritma yang memerlukan jumlah memori tetap tidak tergantung pada input yang diberikan, sehingga membentuk notasi O(1).
   - Contoh kasus: mengalokasikan satu variabel dalam memori.

2. Linear Space Complexity - O(n)

   - Algoritma yang memerlukan jumlah memori yang berbanding lurus dengan jumlah input yang diberikan, sehingga membentuk notasi O(n).
   - Contoh kasus: menyimpan seluruh elemen dalam array.

3. Quadratic Space Complexity - O(n^2)

   - Algoritma yang memerlukan jumlah memori sebesar kuadrat dari jumlah input yang diberikan, sehingga membentuk notasi O(n^2).
   - Contoh kasus: menyimpan matriks n x n.

4. Logarithmic Space Complexity - O(log n)

   - Algoritma yang memerlukan jumlah memori yang berkurang secara logaritmik seiring dengan peningkatan input, sehingga membentuk notasi O(log n).
   - Contoh kasus: rekursi dengan memori tunggal.

# **Data Structure**

## **Struktur Data**

- Struktur data adalah cara untuk menyimpan dan mengorganisasi data agar dapat digunakan secara efisien dan efektif
- Struktur data di golang terdiri dari 3 jenis yaitu array, slice dan map

## **Array**

- Struktur data yang berisi kumpulan data yang memiliki tipe data yang sama
- Memiliki ukuran yang tetap
- Tipe data yang dapat disimpan dalam array dapat berupa number, string dan boolean
- Diawali dengan index 0 dan diakhiri dengan index n-1
- Cara deklarasi array pada golang:
  ```go
    var arr [5]int
    var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
    arr2 := [5]int{1, 2, 3, 4, 5}
    arr3 := [...]int{1, 2, 3, 4, 5}
  ```
- Cara mengakses array dengan looping:
  ```go
    // cara pertama
    for i := 0; i < len(arr1); i++ {
      fmt.Println(arr1[i])
    }
    // cara kedua
    for i, v := range arr1 {
      fmt.Println(i, v)
    }
    // cara ketiga
    i := 0
    for range arr1{
      fmt.Println(i, arr1[i])
      i++
    }
  ```

## **Slice**

- Sama dengan array namun memiliki ukuran yang dinamis
- Merupakan mereferensi dari sebuah array
- Slice memiliki pointer, length dan capacity

- Cara deklarasi slice pada golang:
  ```go
    var slice []int
    slice1 := []int{1, 2, 3, 4, 5}
    slice2 := make([]int, 5)
    slice3 := make([]int, 5, 10)
  ```
- Cara menambah & mengkopy pada slice di goang:
  ```go
    // Menambah
    slice1 = append(slice1, 6)
    slice2 = append(slice2, slice1...)
    // Mengkopy
    slice3 = copy(slice3, slice1)
  ```

## **Map**

- Memiliki struktur data yang berupa pasangan **key** dan **value**, dimana key itu merupakan nilai yang unik
- Cara deklarasi map pada golang:
  ```go
    var map1 map[string]int
    map1 = make(map[string]int)
    map2 := map[string]int{
      "key1": 1,
      "key2": 2,
    }
  ```
- Cara menambah & menghapus pada map di goang:
  ```go
    // Menambah
    map1["key1"] = 1
    // Menghapus
    delete(map1, "key1")
  ```
- Cara mengakses map dengan looping:
  ```go
    for key, value := range map1 {
      fmt.Println(key, value)
    }
  ```

## **Perbedaan Array, Slice dan Map**

| Fitur               | Array          | Slice           | Map                     |
| ------------------- | -------------- | --------------- | ----------------------- |
| Deklarasi           | var arr [5]int | var slice []int | var map1 map[string]int |
| Ukuran Tetap        | Ya             | Tidak           | Tidak                   |
| Elemen dapat diubah | Tidak          | Ya              | Ya                      |
| Akses Elemen        | arr[index]     | slice[index]    | map1["key"]             |
| Referensi           | By value       | By reference    | By reference            |

## **Fungsi**

- Fungsi adalah sekumpulan perintah yang dipanggil dipanggil dengan nama tertentu dan dapat dipanggil berulang-ulang
- Berfungsi untuk membuat clean code, mempermudah dalam pemanggilan perintah dan modular code
- Cara deklarasi fungsi pada golang:

  ```go
    // tanpa parameter
    func fungsi1() {
      fmt.Println("Hello World")
    }
    // dengan parameter & single return
    func fungsi2(param1 string) (string) {
      return param1
    }
    // dengan multi return
    func fungsi3(param1 string, param2 int) (string, int) {
      return param1, param2
    }

  ```

- Cara memanggil fungsi pada golang:
  ```go
    fungsi1()
    // memanggil sekaligus memasukan hasil fungsi single return ke variabel
    param1 := fungsi2("param1")
    // memanggil sekaligus memasukan hasil fungsi multi return ke variabel
    param2, param3 := fungsi3("param1",1)
  ```

# **String, Advance Function, Pointer, Method, Struct and Interface**

## **Beberapa hal yang dapat dilakukan pada string golang**

### Menghitung panjang string

```go
  sentence := "Hello World"
  fmt.Println(len(sentence))
```

### Compare String

```go
  sentence := "Hello World"
  sentence2 := "Hello World"
  fmt.Println(sentence == sentence2)
```

### Contain, yang ada pada library strings

```go
  sentence := "Hello World"
  fmt.Println(strings.Contains(sentence, "Hello"))
```

### Slicesing String, berdasarkan index

```go
  sentence := "Hello World"
  fmt.Println(sentence[0:5])
```

### Replace String

```go
  sentence := "Hello World"
  // -1 untuk mengganti semua huruf e menjadi o
  fmt.Println(strings.Replace(sentence, "e", "o", -1))

  // n untuk mengganti satu huruf ld menjadi ll sebanyak n
  fmt.Println(strings.Replace(sentence, "ld", "ll", 1))
```

### Insert String

```go
  sentence := "Hello World"
  insert := sentence[:5] + " Golang" + sentence[5:]
  fmt.Println(insert)
```

## **Advance Function**

### Variadic Function

- Untuk menskip pembuatan slice sementara
- Mempermudah dalam pembacaan fungsi yang memiliki banyak parameter
- Digunakan ketika jumlah parameter fungsi tidak diketahui
- Contoh:

  ```go
    func sumAll(numbers ...int) int {
      total := 0
      for _, number := range numbers {
        total += number
      }
      return total
    }

    func main() {
      total := sumAll(10, 10, 10, 10)
      fmt.Println(total)
    }
  ```

### Anonymous Function

- Fungsi tanpa nama
- Biasanya digunakan ketika ingin membuat inline function
- Contoh:

  ```go
    func main() {
      anonymous := func() {
        fmt.Println("Hello World")
      }

      anonymous()
    }
  ```

### Closure

- Fungsi yang dapat mengakses variabel diluar fungsi
- Contoh:

  ```go
  func NewCounter() func() int{
    count := 0
    return func() int {
      count++
      return count
    }
  }
  func main() {
    counter := NewCounter()
    fmt.Println(counter())
    fmt.Println(counter())
    fmt.Println(counter())
  }
  ```

### Defer Function

- Fungsi yang akan dieksekusi setelah fungsi yang memanggilnya selesai dieksekusi
- Bersifat LIFO (Last In First Out)
- Biasanya digunakan untuk menutup koneksi database, file, dll
- Contoh:

  ```go
    func main() {
      defer fmt.Println("Start")
      fmt.Println("Second")
    }
  ```

## **Pointer**

- Variable yang menyimpan sebuah alamat memory dari variabel lain
- Digunakan untuk mengubah nilai variabel yang ada di alamat memory tersebut
- "&" untuk mendapatkan alamat memory dari sebuah variabel
- "\*" untuk mendapatkan nilai dari alamat memory yang ditunjuk oleh pointer
- Contoh:

  ```go
    func main() {
      var name string = "Golang"
      var pName *string = &name

      pName = "Hello World"
      fmt.Println(*pName)
    }
  ```

## **Struct dan Interface**

### Struct

- Kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu (seperti class di OOP)
- Contoh:

  ```go
  type Student struct {
    Name string
    Grade int
  }

  func main() {
    // cara 1
    var student1 Student
    student1.Name = "Golang"
    student1.Grade = 2

    fmt.Println(student1)

    // cara 2
    var student1 = Student{Name: "Golang", Grade: 2}

    fmt.Println(student1)

    // cara 3 (urutan diperhatikan)
    var student1 = Student{"Golang", 2}

    fmt.Println(student1)
  }
  ```

### Interface

- Kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan nama tertentu.
- Bisa berfungsi untuk menyimpan dinamic type data
- Contoh:

  ```go
  type Student struct {
    Name string
    Grade int
  }

  type StudentInterface interface {
    sayHello()
  }

  func (student Student) sayHello() {
    fmt.Println("Hello", student.Name)
  }

  func main() {
    var student1 Student
    student1.Name = "Golang"
    student1.sayHello()
  }
  ```

## **Method**

- Fungsi yang menempel pada type (bisa struct atau tipe data lainnya).
- Method bisa diakses lewat variabel objek.
- Digunakan untuk menulis kode dengan style OOP di golang
- Mencegah terjadinya konflik pada nama fungsi
- Lebih mudah dibaca dan dimengerti dari pada fungsi biasa
- Contoh:

  ```go
    type Student struct {
      Name string
      Grade int
    }

    func (student Student) sayHello() {
      fmt.Println("Hello", student.Name)
    }

    func main() {
      var student1 Student
      student1.Name = "Golang"
      student1.sayHello()
    }
  ```

## **Error Handling**

### Panic

- Fungsi yang digunakan untuk menghentikan proses program
- Contoh:

  ```go
    func endApp() {
      fmt.Println("Aplikasi selesai")
    }

    func runApp(error bool) {
      defer endApp()
      if error {
        panic("Aplikasi Error")
      }
      fmt.Println("Aplikasi berjalan")
    }

    func main() {
      runApp(true)
      fmt.Println("Mantap")
    }
  ```

### Recover

- Fungsi yang digunakan untuk menangkap error yang terjadi pada fungsi panic
- Contoh:

  ```go
    func endApp() {
      message := recover()
      fmt.Println(message)
      fmt.Println("Aplikasi selesai")
    }

    func runApp(error bool) {
      defer endApp()
      if error {
        panic("Aplikasi Error")
      }
      fmt.Println("Aplikasi berjalan")
    }

    func main() {
      runApp(true)
      fmt.Println("Mantap")
    }
  ```

# **Recursive, Number Theory, Searching, Sorting**

## **Recursive**

- Sebuah fungsi yang memanggil dirinya sendiri
- Jika masalahnya ringan maka hasil akan didapatkan dengan cepat
- Jika masalahnya besar maka fungsi rekursif akan membagi masalahnya menjadi masalah yang lebih kecil
- Contoh kasus : faktorial, fibonacci, dll

## **Number Theory**

- Sebuah cabang matematika yang mempelajari tentang angka
- Contoh number theory : prime number, gcd, lcm, dll

## **Searching**

- Sebuah algoritma yang digunakan untuk mencari suatu data dari kumpulan data
- Contoh searching : binary search, linear search, dll

## **Sorting**

- Sebuah algoritma yang digunakan untuk mengurutkan data
- Contoh sorting : bubble sort, insertion sort, dll

# **Problem Solving Paradigm - Brute Force Greedy and Dynamic Programming**

## **Problem Solving Paradigm**

- Paradigma adalah cara untuk menyelesaikan suatu masalah

- Paradigma yang paling sering digunakan,

  **1. Brute Force (complete search)**

  - Mencoba semua kemungkinan
  - Mudah untuk dipahami dan diimplementasikan
  - Digunakan ketika tidak ada algoritma lain yang lebih cepat
  - Secara teori semua masalah dapat diselesaikan dengan brute force
  - Kompleksitasnya sangat tinggi
  - Contoh kasus : mencari nilai maximum dan minimum dari suatu array, dll

  **2. Divide and Conquer**

  - Memecah masalah menjadi sub masalah yang lebih kecil
  - Langkah yang dilakukan adalah

    1.  **Divide** : memecah masalah menjadi sub masalah yang lebih kecil
    2.  **Conquer** : menyelesaikan sub masalah yang lebih kecil
    3.  **Combine** : menggabungkan solusi dari sub masalah yang lebih kecil menjadi solusi dari masalah yang lebih besar

  - Contoh kasus : binary search, mencari pangkat suatu bilangan, dll

  **3. Greedy**

  - Algoritma yang mencari lokal optimal (solusi yang terbaik pada saat itu) tanpa mempertimbangkan solusi global optimal (solusi terbaik dari keseluruhan masalah)
  - Contoh kasus : coin change, activity selection, huffman coding, dijkstra algorithm, dll

  **4. Dynamic Programming**

  - Algoritma yang mencari solusi optimal dari suatu masalah dengan memecah masalah menjadi sub masalah yang lebih kecil dan menyimpan solusi dari sub masalah tersebut ke dalam tabel untuk digunakan kembali

  - Ciri utama dari dynamic programming adalah

    1.  **Overlapping sub problem** : sub masalah yang lebih kecil dan sama dipanggil berkali-kali
    2.  **Optimal sub structure** : setiap masalah dapat dipecah menjadi sub masalah yang lebih kecil yang memiliki solusi optimal

  - Method yang digunakan adalah

    1.  **Top down with memoization** : mencari solusi dari masalah yang lebih besar sampai ke sub masalah yang lebih kecil dan menyimpan solusi dari sub masalah tersebut ke dalam tabel untuk digunakan kembali (biasanya diselesaikan rekursif).
    2.  **Bottom up with tabulation** : mencari solusi dari sub masalah yang lebih kecil sampai ke masalah yang lebih besar dan menyimpan solusi setiap sub masalah dalam sebuah tabel untuk digunakan kembali (mengindari rekursif).

  - Contoh kasus : fibonacci, knapsack, longest common subsequence, dll

# **Concurrent Programming**

## **Beberapa Jenis Programming**

- ### Sequential Programming

  - Intruksi yang dieksekusi secara berurutan

- ### Parallel Programming

  - Intruksi yang dieksekusi secara bersamaan (harus ada multi core)

- ### Concurrent Programming

  - Intruksi yang diselesaikan dalam waktu yang sama dan berjalan dengan independent (tidak harus ada multi core)

## **Feature Concurrent pada Golang**

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

# **Clean Code**

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

# **Database**

## **Database**

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

## **Sistem Manajemen Basis Data (DBMS)**

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

## **Perintah SQL**

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

# **Comunication Skill**

## **Komunikasi Efektif**

- Proses pertukaran informasi baik verbal maupun non-verbal di dalam sebuah organisasi dengan tujuan menyampaikan, menerima dan memahami sebuah pesan dengan baik

## **Fungsi Komunikasi Efektif**

- Meningkatkan efektivitas kerja
- Meningkatkan kepuasan karyawan
- Mengurangi tingkat pergantian karyawan

## **5W Element Komunikasi**

- Who (Siapa) : Siapa yang memberikan pesan
- (Says) What (Apa) : Pesan apa yang akan disampaikan
- (In Which) Channel (Media) : Melalui media apa pesan akan disampaikan
- (To) Whom (Siapa) : Siapa yang akan menerima pesan
- (With What) Effect

## **7C Komunikasi**

- **Completeness** (Kelengkapan) : Menyampaikan semua fakta yang dibutukan oleh audiens. Informasi yang lengkap meliputi :

  - Tidak ada informasi yang terlewatkan
  - Tidak membuat penerima semakin bingung
  - Membantu penerima pesan untuk mengambil keputusan

- **Conciseness** (Keringkasan) : Menyampaikan informasi dengan singkat dan jelas. Informasi yang singkat meliputi :

  - Menghemat waktu dan biaya
  - Menyoroti pesan utama
  - Menghindari penggunaan kata yang berlebihan dan tidak diperlukan
  - Tidk mengulang-ulang (bertele-tele)

- **Consideration** (Pertimbangan) : Pahami posisi orang lain. Informasi yang mempertimbangkan meliputi :

  - Utamakan pendekatan "Kamu"
  - Mendengarkan dan mempertimbangkan hal yang diminati penerima
  - Penekanan pada "apa yang mungkin"

- **Clarity** (Kejelasan) : Fokus pada pesan atau tujuan yang spesifik dalam satu waktu. Informasi yang jelas meliputi :

  - Kejelasan terhadap ide dan gagasan
  - Gunakan kata-kata yang tepat, sesuai dan konkret (sesuatu yang nyata)

- **Concreteness** (Konkret) : Tidak kabur dan terlalu umum. Informasi yang konkret meliputi :

  - Di dukung oleh fakta dan angka tertentu
  - Pesan konkret tidak di misinterpretasikan

- **Courtesy** (Kesopanan) : Pesan tersebut harus menunjukan eskpresi komunikator dan juga harus menghormati komunikan. Informasi yang sopan meliputi :

  - Mempertimbangkan sudut pandang kedua belah pihak dan juga apa yang dirasakan oleh komunikan
  - Pesan yang sopan adalah pesan yang positif dan fokus pada audiens
  - Menunjukan penghormatan kepada lawan penerima pesan
  - Tidak bias

- **Correctness** (Kesesuaian) : Bahasa yang tepat dan tidak salah. Informasi yang sesuai meliputi :

  - Pesan yang disampaikan tepat dan di waktu yang tepat
  - Menggunakan bahas ayang benar dan sesuai
  - Memvalidasi ketepatan dan keakuratan fakta dan angka

## **Jenis Komunikasi**

- ### Synchronous (Komunikasi Real Time)

  - Komunikasi yang dilakukan secara langsung dan bersamaan
  - Pro :

    - Menyelesaikan masalah dengan lebih cepat
    - Memudahkan Interperonal (komunikasi antar 2 orang atau lebih) comunication
    - Menjalin hubungan yang baik dengan pihak lain
    - Teamwork lebih terjalin
    - Memungkinkan untuk mendapatkan lebih banyak ide / masukan dari orang lain

  - Cons :

    - Lebih banyak tenaga dan pikiran
    - Memungkinkan untuk membuat keputusan yang tidak berkualitas
    - Output dari meeting bisa tidak jelas
    - Terkadang tidak memiliki agenda

  - Contoh : Telepon, Video Conference (zoom, google meet, dll)

- ### Asynchronous (Komunikasi Non Real Time)

  - Komunikasi yang dilakukan secara tidak langsung dan tidak bersamaan
  - Pro :

    - Memudahkan komunikasi distributed teams dengan zona waktu yang berbeda
    - Komunikasi yang terdokumentasi dengan baik
    - Meningkatkan proses komunikasi setiap saat. Sehingga memperjelas pesan yang ingin disampaikan
    - Biasanya menghasilkan solusi yang lebih berkualitas
    - Memberdayakan tim untuk bekerja sesuai waktu produktif mereka

  - Cons :

    - Tidak segera
    - Memakan waktu yang lebih lama
    - Kurang hubungan interpersonal
    - Terkadang kurang bisa memahami pesan yang disampaikan

  - Tips :
    - Membuat rencana komunikasi :
    - Tingkat visibilitas
    - Jadikan sebagai budaya
    - Komunikasikan jam kerja
  - Contoh : Email, Chat, SMS, dll

# **Join, Union, Agregasi, Subquery, Function (DBMS)**

## **Join**

- Sebuah klausa untuk mengkombinasikan dari dua atau lebih tabel
- Jenis Join standar SQL :
  - Inner Join : Mengembalikan baris-baris dari dua tabel atau lebih yang memenuhi syarat
    - Contoh :
      ```sql
      SELECT * FROM table1
      INNER JOIN table2
      ON table1.column_name = table2.column_name;
      ```
  - Left Join : Mengembalikan semua baris dari tabel kiri yang dikenai kondisi ON dan hanya baris dari baris tabel di kanan yang memenuhi kondisi join
    - Contoh :
      ```sql
      SELECT * FROM table1
      LEFT JOIN table2
      ON table1.column_name = table2.column_name;
      ```
  - Right Join : Mengembalikan semua baris dari tabel kanan yang dikenai kondisi ON dan hanya baris dari baris tabel di kiri yang memenuhi kondisi join
    - Contoh :
      ```sql
      SELECT * FROM table1
      RIGHT JOIN table2
      ON table1.column_name = table2.column_name;
      ```

## **Union**

- Menggabungkan hasil dari dua atau lebih SELECT statement
- Syarat:
  - Jumlah kolom harus sama
  - Tipe data dari kolom harus sama (tapi bila memang bisa di konversi ke tipe data yang sama, maka bisa)
  - Urutan kolom harus sama
- Contoh :
  ```sql
  SELECT column_name FROM table1
  UNION
  SELECT column_name FROM table2;
  ```

## **Fungsi Agregasi**

- Fungsi di mana nilai dari beberapa baris dikelompokan bersama untuk membentuk nilai ringkasan tunggal
- Beberapa fungsi agregasi yang umum digunakan :
  - COUNT() : jumlah baris yang tidak kosong
  - SUM() : jumlah dari nilai dalam kolom
  - AVG() : rata-rata dari nilai dalam kolom
  - MAX() : nilai terbesar dalam kolom
  - MIN() : nilai terkecil dalam kolom
- Contoh penggunaan fungsi agregasi `COUNT`:
  ```sql
  SELECT COUNT(column_name) FROM table_name;
  ```
- HAVING : digunakan untuk menambahkan kondisi pada fungsi agregasi
  - Contoh penggunaan fungsi agregasi `SUM`:
  ```sql
  SELECT user_id FROM tweets GROUP BY user_id HAVING SUM(retweet_count) > 2;
  ```

## **Subquery**

- Subquery adalah query yang berada di dalam query lain
- Subquery dapat berada di dalam SELECT, FROM, WHERE, HAVING, dan JOIN
- Contoh:
  ```sql
  SELECT * FROM table1
  WHERE column_name IN (SELECT column_name FROM table2);
  ```

## **Trigger**

- Trigger adalah sebuah prosedur yang akan dijalankan secara otomatis ketika terjadi suatu event
- Jenis-jenis trigger :
  - BEFORE : trigger akan dijalankan sebelum event terjadi
  - AFTER : trigger akan dijalankan setelah event terjadi
- Contoh :
  ```sql
    DELIMITER $$
    CREATE TRIGGER after_del_transactions
    BEFORE DELETE ON transactions
    FOR EACH ROW
    BEGIN
      DELETE FROM transaction_details WHERE transaction_id = OLD.transaction_id;
    END$$
    DELIMITER ;
  ```

## **Function**

- Sekumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya
- Contoh:

  ```sql
    -- Membuat function
    -- deterministic (private) --> kepemilikan nya si user yang login saja,
    DELIMITER $$
    CREATE FUNCTION tweet_per_user
      (user_id_p int) returns int deterministic
      BEGIN
        DECLARE total_tweet int;
        SELECT count(*) into total_tweet from tweet
        WHERE user_id = user_id_p;
        return total_tweet;
      END$$
    DELIMITER ;
  ```

  ```sql
    -- Memanggil function
    SELECT tweet_per_user(1);
  ```

# **Configuration Management and CLI**

## **Command Line Interface (CLI)**

- Berfungsi untuk mengirimkan perintah-perintah kepada sistem operasi.
- CLI biasanya berupa teks yang ditampilkan di layar monitor.
- Lebih efisien dan efektif ketika mengelola tugas sistem operasi yang kompleks.
- Contoh CLI
  - UNIX Shell
  - Windows Command Prompt
- Applikasi CLI lainnya
  - Python REPL
  - MYSQL CLI Client
  - Mongo shell
  - redis-cli
  - etc

## **Command yang sering digunakan**

### Directory

- **cd** : untuk berpindah direktori
  - urutan permissions file : file_type:owner:group:other
    - **file_type** : tipe file
      - **d** : directory
      - **-** : file
      - **l** : link
    - **owner** : pemilik file
    - **group** : grup pemilik file
    - **other** : pemilik lain
- **ls** : untuk melihat isi direktori
- **pwd** : untuk melihat direktori yang sedang aktif
- **mkdir** : untuk membuat direktori baru
- **rm** : untuk menghapus file
- **mv**: untuk memindahkan file
- **cp** : untuk menyalin file
- **ln** : untuk membuat link
  - hard link : link yang mengarah ke file yang sama (file masih bisa dibuka walaupun file yang di link dihapus)
  - soft link : link yang mengarah ke file yang berbeda (file tidak bisa dibuka jika file yang di link dihapus)

### File

- Create
  - **touch** : untuk membuat file kosong
  - **cat** : untuk membuat file dengan isi
- View
  - **cat** : untuk membaca isi file
  - **less** : untuk membaca isi file dengan fitur scroll
  - **head** : untuk membaca isi file dari awal
  - **tail** : untuk membaca isi file dari akhir
- Editor
  - **nano** : untuk mengedit file
  - **vim** : untuk mengedit file
- Permission
  - **chmod** : untuk mengubah permission file
  - **chown** : untuk mengubah pemilik file
  - permmission file code:
    - **r** : read (4)
    - **w** : write (2)
    - **x** : execute (1)
  - jenis user:
    - normal user : diperbolehkan untuk mengubah file yang dimiliki sendiri
    - root user : diperbolehkan untuk mengubah file yang dimiliki orang lain

### Network

- **ping** : untuk mengecek koneksi ke host
- **ifconfig** : untuk melihat konfigurasi jaringan
- **netstat** : untuk melihat status jaringan
- **curl** : untuk mengirim request ke host
- **wget** : untuk mengunduh file dari host
- **ssh** : untuk mengakses host lain
- **netcat** : untuk mengirim dan menerima data melalui jaringan
- **telnet** : untuk mengakses host lain
- **nmap** : untuk melakukan scanning jaringan

### Utility

- **man** : untuk melihat manual command
- **env** : untuk melihat environment variable
- **export** : untuk menambahkan environment variable
- **echo** : untuk menampilkan teks
- **date** : untuk melihat tanggal dan waktu
- **which** : untuk melihat lokasi command
- **watch** : untuk menjalankan command secara otomatis
- **sudo** : untuk menjalankan command sebagai root user
- **history** : untuk melihat history command
- **grep** : untuk mencari teks
- **locate** : untuk mencari file
- **find** : untuk mencari file

## **Shell Script**

- Shell adalah program yang berfungsi sebagai penghubung antara user dan kernel.
- Shell script adalah sebuah bahasa pemrograman yang dicompile berdasarkan perintah shell.
- Contoh shell script
  - Bash
  - Zsh
  - Fish
  - etc
- Contoh shell script
  ```bash
    #!/bin/bash
    echo "Hello World"
  ```
- Cara menjalankan shell script
  - **chmod +x** : untuk memberikan permission rwx-rwx-rwx pada file
  - **./** : untuk menjalankan script

# **System Design**

## **Diagram**

- Representasi visual dari suatu sistem yang berisi komponen-komponen sistem dan hubungan antar komponen tersebut.
- Software design yang biasa digunakan:
  - **Flowchart** : Representasi grafis dari sebuah proses atau algoritma
  - **Use Case** : Deskripsi dari perilaku sebuah sistem dari sudut pandang pengguna
  - **ERD (Entity Relationship Diagram)** : Diagram yang menunjukkan hubungan antara entitas dalam sebuah database
  - **HLA (High Level Architecture)** : Desain arsitektur tingkat atas yang menguraikan komponen dan interaksi dari sebuah sistem

## **Distributed System**

- Sistem yang terdiri dari beberapa komputer yang saling berkomunikasi dan saling bekerja sama untuk mencapai tujuan tertentu.
- Karakteristik:
  - scalability : kemampuan untuk menambahkan komponen baru ke sistem
    - vertical scaling : menambahkan komponen baru dengan spesifikasi yang lebih tinggi (lebih lama)
    - horizontal scaling : menambahkan komponen baru dengan spesifikasi yang sama (lebih cepat)
  - realibility : keandalan dari sistem dalam menjalankan tugasnya
  - availability : ketersediaan dari sistem dalam menjalankan tugasnya (tidak down)
  - efiency : kecepatan dalam menjalankan tugasnya (response time)
  - serviceability or manageability : harus mudah untuk di maintain dan di manage (good clean code)

## **Job Queue, Load Balancer, Monolitic dan Microservice**

- Job Queue : sistem yang memproses tugas-tugas yang masuk dan menyimpannya dalam antrian (queue) untuk di proses oleh framework atau sistem lainnya (worker)
- Load Balancer : sistem yang membagi tugas-tugas yang masuk ke beberapa komputer (server) secara merata
- Microservice : sistem yang terdiri dari beberapa komponen yang saling berkomunikasi dan saling bekerja sama untuk mencapai tujuan tertentu
- Monolitic : sistem yang terdiri dari satu komponen utama yang menjalankan semua tugas

## **SQL vs NoSQL**

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

## **Cache dan Indexing**

- Cache : data yang disimpan di memori untuk mengurangi waktu akses ke data yang sama
- Redudancy dan Replication

  - Redudancy : data yang disimpan di beberapa tempat
  - Replication : data yang disimpan di beberapa tempat dan diupdate secara sinkron

- Indexing : cara untuk mengoptimasi pencarian data dalam database dengan cara menambahkan kolom tambahan index agar dapat
  - Lebih cepat untuk mencari data
  - Lebih lambat untuk menambahkan data

# **Introduction to RESTful API**

## **API (Application Programming Interface)**

- Sebuah interface yang memungkinkan dua aplikasi untuk berkomunikasi satu sama lain.
- Integrasi API:
  - Dari Frontend ke Backend, seperti: dari web ke service
  - Dari Backend ke Backend, seperti: dari service ke service
- Salah satu contoh API adalah REST API (Representational State Transfer)
  - REST API adalah API yang menggunakan REST (Representational State Transfer) sebagai arsitektur utamanya.
- API tools
  - Postman
  - Katalon
  - Apache JMeter
  - SoapUI
  - Insomnia
  - etc.

## **Design REST API**

- Menggunakan kata benda dibandingkan dengan kata kerja
- Menamai dengan menggunakan kata benda jamak
- Menggunakan percabangan untuk memperlihatkan relasi atau hubungan antar resource
- Menggunakan format response JSON
- Filtering, Sorting, Paging, dan Field Selection
  - Filtering
    - GET /users?name=John
  - Sorting
    - GET /users?sort=name
  - Paging
    - GET /users?page=2&size=10
  - Field Selection
    - GET /users?fields=name,age
- Menghandle garis miring trailing dengan baik
  - POST: /users
  - POST: /users/
- Menggunakan versioning
  - GET: /v1/users
  - GET: /v2/users
- Membuat API documentation, seperti
  - https://mailchimp.com/developer/marketing/docs/fundamentals/
  - https://docs.github.com/en/rest/guides/getting-started-with-the-rest-api
  - https://developer.twitter.com/en/docs/twitter-api/v1/tweets/search/overview
  - [etc](https://github.com/public-apis/public-apis)

## **Cara Implementasi REST API dengan Go**

- Menggunakan package build-in:
  - `net/http` : untuk membuat server dan mengconsume API
  - `encoding/json` : untuk menghandle JSON
- Menggunakan third party lib:
  - [echo](https://echo.labstack.com/)
  - [cobra](https://github.com/spf13/cobra)
  - [gorm](https://gorm.io/docs/)
  - [logrus](https://github.com/sirupsen/logrus)
  - go kit
  - [etc](https://github.com/go-kit/kit)

## **Reference**

- [Awesome Package Golang](https://awesome-go.com/)
- [Check Third Party Lib](https://osv.dev/)

# **Introduction Echo Golang**

## **Framework Echo**

- performa tinggi, ekstensibilitas, dan minimalist
- keuntungan:
  - optimized router
  - middleware
  - data rendering
  - scalable
  - data binding
- struktur pada framework echo bebas yang bisa diintegrasikan dengan third party lib lainnya

## **Basic Echo**

- `echo.New()` untuk membuat instance echo
- routing adalah url yang akan diakses oleh user, dapat berupa GET, POST, PUT, DELETE, dan lainnya

  - Contoh pada Echo:
    - `e.GET("/hello", hello)`
    - `e.POST("/hello", hello)`
    - `e.PUT("/hello", hello)`
    - `e.DELETE("/hello", hello)`
    - etc

- controller pada echo menggunakan `echo.Context` untuk mempermudah routing
  - Contoh:
    ```go
    func hello(c echo.Context) error {
    	return c.String(http.StatusOK, "Hello, World!")
    }
    ```
- mengambil param dari url
  - Contoh:
    ```go
    func hello(c echo.Context) error {
    	name := c.Param("name")
    	return c.String(http.StatusOK, "Hello, "+name)
    }
    ```
- mengambil query dari url
  - Contoh:
    ```go
    func hello(c echo.Context) error {
    	name := c.QueryParam("name")
    	return c.String(http.StatusOK, "Hello, "+name)
    }
    ```
- mengambil data dari form
  - Contoh:
    ```go
    func hello(c echo.Context) error {
    	name := c.FormValue("name")
    	return c.String(http.StatusOK, "Hello, "+name)
    }
    ```
- binding data json ke struct

  - Contoh:

    ```go
    type User struct {
    	Name  string `json:"name"`
    	Email string `json:"email"`
    }

    func hello(c echo.Context) error {
    	u := new(User)
    	if err := c.Bind(u); err != nil {
    		return err
    	}
    	return c.JSON(http.StatusOK, u)
    }
    ```

# **ORM dan Code Structure (MVC)**

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

# **Middleware**

## **Middleware**

- Sebuah fungsi yang dapat dijalankan sebelum atau sesudah sebuah request dijalankan.

## **Middleware pada Echo**

- ### Middleware dalam echo dibagi menjadi 2:

  1. Echo#Pre() : Middleware yang dijalankan sebelum router memproses request
  2. Echo#Use() : Middleware yang dijalankan sesudah router memproses request dan memiliki akses penuh terhadap echo.Context API

- ### Contoh middleware pada Echo

  - Middleware Logger, middleware yang digunakan untuk menampilkan log request yang masuk ke server.
  - Middleware BasicAuth, middleware yang digunakan untuk melakukan autentikasi terhadap request yang masuk ke server.
  - Middleware JWT, middleware yang digunakan untuk melakukan autentikasi terhadap request yang masuk ke server dengan menggunakan JWT.
  - [etc.](https://echo.labstack.com/middleware/)

# **Unit Testing**

## **Unit Testing**

- ### Pengertian
  Sebuah proses untuk memastikan bahwa setiap unit dari kode program berjalan dengan benar.
- ### Tujuan
  - **Preventing Regression** : Ketika ada perubahan kode, maka unit test akan menunjukkan apakah ada perubahan yang tidak diinginkan.
  - **Confidence Level in Refactoring** : Membuat lebih percaya diri ketika melakukan refactoring.
  - **Improve Code Design** : Unit test akan menunjukkan apakah kode kita sudah sesuai dengan desain yang kita inginkan.
  - **Code Documentation** : Unit test akan menjadi dokumentasi dari kode kita.
  - **Scaling The Team** : Unit test akan memudahkan tim dalam melakukan scaling.
- ### Level of Testing
  - **UI** : Testing yang dilakukan pada tampilan aplikasi.
  - **Integration** : Spesific testing pada modul-modul atau sub system melalui api
  - **Unit** : Testing yang dilakukan pada fungsi atau method.
- ### Runner
  Sebuah program yang digunakan untuk menjalankan unit test.
- ### Mocking
  Sebuah proses untuk membuat object palsu yang mensimulasikan prilaku object aslinya.
- ### Jangan Melakukan Test Pada,
  - 3rd party modules
  - Database
  - 3rd party api atau sistem external
  - Object class yang sudah di test di tempat lain
- ### Coverage Testing

  Sebuah proses untuk mengukur seberapa banyak kode program yang tercakup oleh unit testing, biasanya dilakukan dengan menggunakan code coverage tool. Coverage testing biasanya dilakukan dengan menggunakan 4 jenis coverage testing:

  - Statement coverage: mengukur seberapa banyak pernyataan kode program yang tercakup oleh unit testing.
  - Branch coverage: mengukur seberapa banyak percabangan atau jalur alternatif dari kode program yang tercakup oleh unit testing.
  - Function coverage: mengukur seberapa banyak fungsi atau metode dalam kode program yang tercakup oleh unit testing.
  - Line coverage: mengukur seberapa banyak baris kode program yang tercakup oleh unit testing.

  Salah satu tool yang dapat digunakan untuk melakukan coverage testing adalah [sonarqube](https://www.sonarqube.org/).

- ### Test Case
  - Sekumpulan dari positif dan negatif test case yang digunakan untuk mengetes suatu fungsi atau method dari skenario tertentu.
  - Test case memiliki sekumpulan dari pre-condition, steps, expected, status, dan actual result

## **Testing Framework**

- ### Pengertian
  Sebuah library yang digunakan untuk membuat unit test.
- ### Jenis Pada Golang
  - Ginkgo
  - Gomega
  - Testify
  - [etc](https://wikipedia.org/wiki/List_of_unit_testing_frameworks#Go)
- ### Structure

  - Stategi yang di implementasikan dalam menaruh file test
  - 2 Usual Pattern:
    1. **Centralized** test file di dalam folder `tests`
    2. Simpan file test **Together** dengan production file
  - Komponen penting di dalam unit test:
    1. Test File (Koleksi dari test suites)
    2. Test Suites (Koleksi dari test cases)
    3. Test Fixtures (Setup & Teardown), Test Cases (Input,Process,Output)

## **Implementasi Unit Testing Pada Golang**

- Buat file baru, dengan format `<library>_test.go`. didalam package yang sama
- Tulis kode berikut pada file tersebut

  ```go
  package <nama_package>

  import (
  	"testing"
  )

  func Test<nama_function>(t *testing.T) {
  	// test code
  }
  ```

- Jalankan unit test dengan perintah `go test ./<nama_package>/... -cover`. -cover untuk melihat coverage testing
- Jika ingin melihat coverage testing secara detail, gunakan perintah `go test <nama_package>/... -coverprofile=coverage.out && go tool cover -html=coverage.out`

# **Clean dan Hexagonal Architecture**

## **Hexagonal Architecture**

- ### Pengertian

  Hexagonal Architecture adalah sebuah arsitektur yang memisahkan aplikasi menjadi dua bagian utama, yaitu domain dan infrastructure. Domain berisi data dan logika bisnis, sedangkan infrastructure berisi tampilan dan logika untuk mengatur data dan tampilan.

- ### Struktur Hexagonal Architecture

  - API interface, Berisi hal yang berhubungan dengan API, seperti rest controller/framework/CLI
  - Domain service, Berisi data dan logika bisnis, seperti model, repository, dan service
  - SPI interface, Berisi hal yang berhubungan dengan database, seperti mysql, redis, third party, dll

## **Clean Architecture**

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

## **Ubiquitous Language**

- Sebuah istilah yang digunakan untuk mendeskripsikan bahasa yang digunakan oleh semua orang yang terlibat dalam suatu project. Ubiquitous Language ini biasanya berupa kata-kata yang digunakan untuk mendeskripsikan suatu hal, seperti kata-kata yang digunakan untuk mendeskripsikan suatu objek, kata-kata yang digunakan untuk mendeskripsikan suatu fungsi, kata-kata yang digunakan untuk mendeskripsikan suatu proses, dll.

# **Proffessional Skill 2 Building Rangkuman and CV**

## **Kesalahan Umum**

- ### Salah ketik
- ### Tidak berstruktur
- ### Sangat klise dan banyak buzzwords
- ### Tidak ada achivement
- ### Terlalu panjang
- ### CV umum atau tidak spesific

## **Tips Terpenting**

- ### Empati
- ### Secara pyschological HRD akan memperhatikan berdasarkan:
  - #### Pola Membaca
  - #### Struktur Informasi
  - #### Kelelahan
- ### Secara visual HRD akan memperhatikan berdasarkan:
  - #### Ukuran & Kontras
  - #### Negative Spaces

## **Konten pada CV**

- ### Relevan
- ### Ada angka lebih baik
- ### Hindari buzzwords
- ### Singkat
- ### Perhatikan kata kunci
- ### Pakai font yang tepat
- ### Baca ulang

# **Docker**

## **Docker**

- ### Perbedaan Container dengan VM

  Container adalah sebuah proses dengan file system dan resource yang terisolasi sedangkan VM adalah sebuah proses yang menjalankan sebuah OS. Container berjalan diatas OS yang sudah ada sedangkan VM menjalankan OS yang baru. Container lebih ringan dibandingkan VM karena tidak perlu menjalankan OS yang baru. Container juga lebih cepat dibandingkan VM karena tidak perlu menjalankan OS yang baru. Container juga lebih efisien dibandingkan VM karena tidak perlu menjalankan OS yang baru. Container juga lebih aman dibandingkan VM karena tidak perlu menjalankan OS yang baru.

- ### Docker Infrastructure

  - Client :
    - docker build
    - docker pull
    - docker run
  - Docker Host :
    - Docker Daemon: containers, images
  - Docker Registry : Docker Hub

- ### Dockerfile
  file yang mendeskripsikan langkah-langkah untuk membuild image
- ### Docker Image
  file yang berisi semua yang dibutuhkan untuk menjalankan sebuah aplikasi. Image ini bisa dijalankan menjadi container. dan yang akan di push ke docker registry dan dari docker registry yang akan di pull ke docker host.
- ### Docker Container
  proses yang berjalan diatas OS yang sudah ada. Container ini bisa dijalankan menjadi container.
- ### **Docker Volume**
  Tempat penyimpanan data yang bisa diakses oleh container, sehingga dengan menggunakan volume container bisa dihapus tanpa menghapus penyimpanan data yang ada di dalamnya.
- ### **Docker Network**
  Sebuah virtual network yang bisa digunakan untuk menghubungkan container-container yang berbeda. Docker Network juga bisa digunakan untuk menghubungkan container dengan host.

## **Strategi Deployment**

- ### Big-Bang Deployment (Replace Deployment)
- ### Rollout Deployment
- ### Blue/Green Deployment
- ### Canary Deployment

# **Compute Service**

## **Simple Dev-Ops Cycle**

- Dev:
  - Plan
  - Code
  - Build
  - Test
- Ops:
  - Release
  - Deploy
  - Operate
  - Monitor

## **Compute Service**

- Layanan yang menyediakan komputasi untuk VM. Seperti ec2 pada AWS, compute engine pada GCP, dll.

## **AWS RDS (Relational Database Service)**

- Layanan yang menyediakan database relasional yang dapat diatur secara otomatis. Layanan ini dapat digunakan untuk mengelola database MySQL, PostgreSQL, MariaDB, Oracle, dan Microsoft SQL Server. Layanan ini dapat digunakan untuk mengelola database MySQL, PostgreSQL, MariaDB, Oracle, dan Microsoft SQL Server.

## **Link Deploy dengan AWS EC2 dan RDS**

- POST: [register](http://ec2-18-136-195-242.ap-southeast-1.compute.amazonaws.com:8000/users)
- GET: [login](http://ec2-18-136-195-242.ap-southeast-1.compute.amazonaws.com:8000/users/login)
- dll

# **CI/CD**

## **CI (Continuous Integration)**

Sebuah proses otomatis untuk mengintegrasikan berbagai kode dari potensial sumber yang berbeda untuk di build atau di test/

## **CD (Continuous Delivery/Deployment)**

- ### **Continuous Deployment**

  Sebuah proses deploy yang dilakukan secara otomatis setiap build atau test yang dilakukan pada proses CI berhasil.

- ### **Continuous Delivery**

  Sebuah proses deploy yang dilakukan secara otomatis setiap build atau test yang dilakukan pada proses CI berhasil. Namun, proses deploy ini tidak langsung di deploy ke production, melainkan ke environment yang berbeda terlebih dahulu.

## **Tools CI/CD**

- ### **Code & Commit**
  - Git
  - Github
  - Visual Studio
  - dll
- ### **Build & Config**
  - Docker
  - Maven
  - Visual Studio
  - dll
- ### **Scan & Test**
  - Sonar
  - SOASTA
  - Selenium Webdriver
- ### **Release**
  - uDeploy
  - SERENA
  - CollabNet
  - dll
- ### **Deploy**
  - Docker
  - Pivotal
  - Kubernetes
  - dll
