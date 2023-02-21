# **Rangkuman Materi Version Control and Branch Management (Git)**

# Version Control System (VCS)

Version Control adalah sebuah sistem yang digunakan untuk melacak suatu perubahan pada sebuah file. Version Control System (VCS) adalah sebuah sistem yang digunakan untuk mengelola perubahan pada sebuah file atau sekumpulan file. VCS memungkinkan kita untuk melihat perubahan apa saja yang terjadi pada file atau sekumpulan file, kapan perubahan tersebut terjadi, siapa yang melakukan perubahan tersebut, dan mengembalikan file ke versi sebelumnya.

# Sejarah Singkat Version Control

1.  Local Version Control (Single User)

    - Kelebihan :
      - mudah digunakan
      - tidak memerlukan koneksi internet
    - Kekurangan :
      - tidak ada backup
      - jika terjadi kerusakan pada database, maka tidak bisa melakukan perubahan pada file
    - Contoh:
      - SCCS - 1972 hanya untuk sistem operasi Unix
      - RCS - 1982 cross platform, hanya untuk text

2.  Centralize (Client-Server)
    - Kelebihan :
      - mudah digunakan, tidak memerlukan koneksi internet
    - Kekurangan :
      - jika server down, maka tidak bisa melakukan perubahan pada file
    - Contoh :
      - CVS - 1995 hanya untuk file
      - Perforce - 1995
      - Subversion - 2000 - melacak perubahan struktur direktori
      - Microsoft Team Foundation Server - 2005
3.  Distributed (Peer-to-Peer)
    - Kelebihan :
      - jika server down, tetap bisa melakukan perubahan pada file
    - Kekurangan :
      - memerlukan koneksi internet
    - Contoh :
      - Git - 2005
      - Mercurial - 2005
      - Bazaar - 2005

# 3 Dasar Staging Area

1. Working Directory : tempat kerja kita saat ini
2. Staging Area : tempat untuk menyimpan perubahan yang akan di commit
3. Repository : tempat penyimpanan file yang sudah di commit

# Command git

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

# Tips untuk mengoptimasi alur kerja dengan github

1. Hindari perubahaan di branch master, branch master hanya digunakan untuk menyimpan versi yang sudah stabil
2. Hindari melakukan perubahan secara langsung di branch development, sebaiknya buat branch baru untuk setiap fitur yang akan dikerjakan
3. Setiap branch fitur yang telah selesai minta pull request ke branch development
4. Setiap branch development yang telah selesai minta pull request ke branch master
