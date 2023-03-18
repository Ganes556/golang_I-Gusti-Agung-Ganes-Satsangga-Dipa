# **Jawaban Sub Dari Prioriats-1-1**

## **Note : folder code-revisi adalah kode yang sudah menggunakan clean code. Serta penjelasannya agar tidak bolak-balik liat Jawaban.md**

## 1. Berapa banyak kekurangan dalam penulisan kode tersebut?

- Secara garis besar terdapat 10 kekurangan dalam penulisan kode tersebut, sisanya tinggal mengikuti apabila telah diperbaiki 10 kekurang tersebut

## 2. Bagian mana saja terjadi kekurangan tersebut?

```go
package main
// penamaan struct `user`
type user struct {
  id       int
  // tipe data `int`
  username int
  // tipe data `int`
  password int
}
// penamaan struct `userservice`
type userservice struct {
  // penamaan property `t`
  t []user
}

// penamaan method `getallusers`
// penamaan param instance `u`
func (u userservice) getallusers() []user {
  return u.t
}
// penamaan method `getuserbyid`
// penamaan param instance `u`
func (u userservice) getuserbyid(id int) user {
  // penamaan variabel `r`
  for _, r := range u.t {
    if id == r.id {
      return r
    }
  }

  return user{}
}
```

## 3. Alasan dari setiap kekurangan tersebut?

```go
package main

// penamaan struct user, seharusnya menggunakan huruf kapital diawal `User` untuk membedakan mana yang struct / fungsi / method /instance struct / parameter
type user struct {

  id       int

  // tipe data username, seharusnya menggunakan tipe data `string` karena tidak ada username yang menggunakan integer/numeric saja
  username int

  // tipe data password, seharusnya menggunakan tipe data `string` karena tidak ada password yang hanya menggunakan integer/numeric saja
  password int
}

// penamaan struct userservice, seharusnya menggunakan CamelCase dan untuk struct menggunakan huruf kapital diawal jadi `UserService` untuk mempermudah membedakan mana yang struct / fungsi / method /instance struct / parameter dan juga mudah untuk dieja
type userservice struct {
  // penamaan property t, yang seharusnya `users []User` karena merupakan array atau banyak user dari struct `User`. sehingga akan lebih mudah untuk dipahami maksud dari propertinya
  t []user
}

// penamaan method getallusers, dirubah dengan menggunakan CamelCase agar lebih mudah dieja dan dicari
// nama instance parameter u, dirubah menjadi `userService` agar lebih menjelaskan maksud dari instance tersebut yakni `UserService` yang merupakan instance dari struct `UserService`
func (u userservice) getallusers() []user {
  return u.t
}

// penamaan method getuserbyid, dirubah dengan menggunakan CamelCase agar lebih mudah dieja dan dicari
// nama instance parameter u, dirubah menjadi `userService` agar lebih menjelaskan maksud dari instance tersebut yakni `UserService` yang merupakan instance dari struct `UserService`
func (u userservice) getuserbyid(id int) user {
  // untuk u.t bila sudah diperbaiki di sebelumnya akan menjadi `userService.users`, sehingga penamaan variabel r lebih cocok menjadi `user` karena akan lebih mudah dipahami maksud dari variabel tersebut sesuai dengan bahasa inggris dimana secara umum penambahan s untuk jamak dan tanpa s untuk tunggal
  for _, r := range u.t {
    if id == r.id {
      return r
    }
  }

  return user{}
}
```
