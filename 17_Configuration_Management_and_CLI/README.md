# **Rangkuman Materi Configuration Management and CLI**

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

### **1. Directory**

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

### **2. File**

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

### **3. Network**

- **ping** : untuk mengecek koneksi ke host
- **ifconfig** : untuk melihat konfigurasi jaringan
- **netstat** : untuk melihat status jaringan
- **curl** : untuk mengirim request ke host
- **wget** : untuk mengunduh file dari host
- **ssh** : untuk mengakses host lain
- **netcat** : untuk mengirim dan menerima data melalui jaringan
- **telnet** : untuk mengakses host lain
- **nmap** : untuk melakukan scanning jaringan

### **4. Utility**

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
