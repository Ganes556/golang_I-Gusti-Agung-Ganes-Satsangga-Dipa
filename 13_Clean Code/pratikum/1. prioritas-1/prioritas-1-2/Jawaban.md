# **Jawaban Sub Dari Prioriats-1-2**

## **Note : folder code-revisi adalah kode yang sudah menggunakan clean code. Serta penjelasannya agar tidak bolak-balik liat Jawaban.md**

## 1. Berapa banyak kekurangan dalam penulisan kode tersebut?

- Secara garis besar terdapat 12 kekurangan dalam penulisan kode tersebut, sisanya tinggal mengikuti apabila telah diperbaiki 12 kekurang tersebut

## 2. Bagian mana saja terjadi kekurangan tersebut?

```go
package main

// penamaan struct `kendaraan`
type kendaraan struct {
  // penamaan property `totalroda`
	totalroda       int
  // penamaan property `kecepatanperjam`
  // tipe data `int`
	kecepatanperjam int
}

// penamaan struct `mobil`
type mobil struct {
	kendaraan
}

// penamaan param instance `m`
func (m *mobil) berjalan() {
	m.tambahkecepatan(10)
}
// penamaan param instance `m`
// penamaan method `tamabahkecepatan`
// penamaan method param `kecepatanbaru`
// tipe data method param `int`
func (m *mobil) tambahkecepatan(kecepatanbaru int) {
	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
}

func main() {
  // penamaan variabel `mobilcepat`
	mobilcepat := mobil{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()
  // penamaan variabel `mobillamban`
	mobillamban := mobil{}
	mobillamban.berjalan()
}
```

## 3. Alasan dari setiap kekurangan tersebut?

```go
package main

// seharusnya merubahnya menjadi `Kendaraan`, agar dapat lebih mudah membedakan mana yang struct / fungsi / method /instance struct / parameter
type kendaraan struct {
  // seharusnya merubahnya menjadi bentuk CamelCase `totalRoda`, agar lebih mudah dieja dan dicari
	totalroda       int
  // seharusnya merubahnya menjadi bentuk CamelCase `kecepatanPerJam`, agar lebih mudah dieja dan dicari
  // seharusnya  tipe data `float64`, karena kecepatan per jam bisa dalam bentuk desimal
	kecepatanperjam int
}
// seharusnya menjadi `Mobil`, agar dapat lebih mudah membedakan mana yang struct / fungsi / method /instance struct / parameter
type mobil struct {
	kendaraan
}
// seharusnya merubah nama instance param menjadi `mobil`, agar mudah dipahami intance yang dimaksudkan
func (m *mobil) berjalan() {
	m.tambahkecepatan(10)
}
// seharusnya merubah nama instance param menjadi `mobil`, agar mudah dipahami intance yang dimaksudkan
// seharusnya merubah nama method menjadi bentuk CamelCase `tambahKecepatan`, agar lebih mudah dieja dan dicari
// seharusnya merubah nama param menjadi bentuk CamelCase `kecepatanBaru`, agar lebih mudah dieja dan dicari
// seharusnya merubah tipe data param menjadi `float64`, karena kecepatan baru bisa dalam bentuk desimal
func (m *mobil) tambahkecepatan(kecepatanbaru int) {
	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
}

func main() {
  // seharusnya merubah nama variabel menjadi bentuk CamelCase `mobilCepat`, agar lebih mudah dieja dan dicari
	mobilcepat := mobil{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()

  // seharusnya merubah nama variabel menjadi bentuk CamelCase `mobilLamban`, agar lebih mudah dieja dan dicari
	mobillamban := mobil{}
	mobillamban.berjalan()
}
```
