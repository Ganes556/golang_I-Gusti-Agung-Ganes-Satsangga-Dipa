# **Resume Materi Basic Programming**

## **Pada materi basic programming terdapat 3 poin utama yaitu:**

## **1. Tentang Golang**

- Bahasa pemrograman open-source yang dikembangkan oleh Google pada tahun 2007.
- Bertujuan untuk memudahkan proses development, sederhana, handal, dan efisien.
- Dirilis ke publik di tahun 2009.
- Golang dikembangkan oleh tim yang terdiri dari Rob Pike, Ken Thompson, Robert Griesemer, Ian Lenz Taylor dan Russ Cox.
- Sering digunakan untuk membuat aplikasi e-commerce, musik player, sosial media, FPI, game engine, dan lain-lain.

## **2. Cara Instal Golang di VS Code**

Dapat dengan cara Mengunduh [VS Code](https://code.visualstudio.com/download) dan menginstalnya, mengunduh [Golang](https://go.dev/dl/) dan menginstalnya di. Kemudian, setting environment path golang yang di sesuaikan pada os masing-masing.

## **3. Perintah Terminal dan Penjelasan Sintak dasar di Golang**

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
    - Pperator, melakukan operasi pada variabel, literal, dan ekspresi)
    - Expression kombinasi antara konstanta, variabel, fungsi, dan operator yang mengembalikan sebuah nilai.

  - Control structure yang merupakan struktur yang digunakan untuk mengatur alur program yang dimana pada golang hanya terdapat **if**, **switch** dan **for**
