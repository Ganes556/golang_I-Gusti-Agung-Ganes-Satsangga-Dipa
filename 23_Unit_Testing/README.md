# Rangkuman Unit Testing

## Unit Testing

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

## Testing Framework

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

## Implementasi Unit Testing Pada Golang

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
