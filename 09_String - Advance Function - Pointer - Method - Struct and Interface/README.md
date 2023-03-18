# **Rangkuman Materi String, Advance Function, Pointer, Method, Struct and Interface**

# Beberapa hal yang dapat dilakukan pada string golang

## 1. Menghitung panjang string

```go
  sentence := "Hello World"
  fmt.Println(len(sentence))
```

## 2. Compare String

```go
  sentence := "Hello World"
  sentence2 := "Hello World"
  fmt.Println(sentence == sentence2)
```

## 3. Contain, yang ada pada library strings

```go
  sentence := "Hello World"
  fmt.Println(strings.Contains(sentence, "Hello"))
```

## 4. Slicesing String, berdasarkan index

```go
  sentence := "Hello World"
  fmt.Println(sentence[0:5])
```

## 5. Replace String

```go
  sentence := "Hello World"
  // -1 untuk mengganti semua huruf e menjadi o
  fmt.Println(strings.Replace(sentence, "e", "o", -1))

  // n untuk mengganti satu huruf ld menjadi ll sebanyak n
  fmt.Println(strings.Replace(sentence, "ld", "ll", 1))
```

## 6. Insert String

```go
  sentence := "Hello World"
  insert := sentence[:5] + " Golang" + sentence[5:]
  fmt.Println(insert)
```

# Advance Function

## 1. Variadic Function

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

## 2. Anonymous Function

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

## 3. Closure

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

## 4. Defer Function

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

# Pointer

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

# Struct dan Interface

## 1. Struct

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

# 2. Interface

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

# Method

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

# Error Handling

## 1. Panic

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

## 2. Recover

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
