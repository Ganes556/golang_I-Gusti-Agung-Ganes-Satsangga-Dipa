# **Rangkuman Materi Problem Solving Paradigm - Brute Force Greedy and Dynamic Programming**

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
