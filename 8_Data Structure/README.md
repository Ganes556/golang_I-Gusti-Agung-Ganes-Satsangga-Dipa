# **Rangkuman Data Structure**

# Struktur Data

- Struktur data adalah cara untuk menyimpan dan mengorganisasi data agar dapat digunakan secara efisien dan efektif
- Struktur data di golang terdiri dari 3 jenis yaitu array, slice dan map

# Array

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

# Slice

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

# Map

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

# Perbedaan Array, Slice dan Map

| Fitur               | Array          | Slice           | Map                     |
| ------------------- | -------------- | --------------- | ----------------------- |
| Deklarasi           | var arr [5]int | var slice []int | var map1 map[string]int |
| Ukuran Tetap        | Ya             | Tidak           | Tidak                   |
| Elemen dapat diubah | Tidak          | Ya              | Ya                      |
| Akses Elemen        | arr[index]     | slice[index]    | map1["key"]             |
| Referensi           | By value       | By reference    | By reference            |

# Fungsi

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
    param3, param4 := fungsi3("param1",1)
  ```
