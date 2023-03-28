package main

// merubah menjadi Kendaraan, agar dapat lebih mudah membedakan mana yang struct / fungsi / method /instance struct / parameter
type Kendaraan struct {
	// merubah menjadi bentuk CamelCase agar lebih mudah dieja
	totalRoda       int
	// merubah menjadi bentuk CamelCase agar lebih mudah dieja
	// merubahnya ke tipe data float64 karena kecepatan per jam bisa dalam bentuk desimal
	kecepatanPerJam float64
}
// merubah menjadi Mobil, agar dapat lebih mudah membedakan mana yang struct / fungsi / method /instance struct / parameter
type Mobil struct {
	// mengikuti perubahan sebelumnya
	Kendaraan
}
// merubah nama instance param menjadi mobil, agar mudah dipahami intance yang dimaksudkan
func (mobil *Mobil) berjalan() {
	// mengikuti perubahan pada method tambahKecepatan
	mobil.tambahKecepatan(10)
}
// merubah nama method menjadi bentuk CamelCase, agar lebih mudah dieja dan dicari
// merubah nama parameter menjadi bentuk CamelCase `kecepatanBaru`, agar lebih mudah dieja dan dicari
// serta merubahnya dalam tipe data float64 karena kecepatan bisa dalam bentuk desimal
// perubahan lainnya mengikuti perubahan sebelumnya seperti pada param instance mobil, mobil.kecepatanPerjam
func (mobil *Mobil) tambahKecepatan(kecepatanBaru float64) {
	mobil.kecepatanPerJam = mobil.kecepatanPerJam + kecepatanBaru
}

func main() {
	// merubah variabel ke bentuk CamelCase, agar lebih mudah dieja dan dicari
	mobilCepat := Mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	// merubah variabel ke bentuk CamelCase, agar lebih mudah dieja dan dicari
	mobilLamban := Mobil{}
	mobilLamban.berjalan()
}