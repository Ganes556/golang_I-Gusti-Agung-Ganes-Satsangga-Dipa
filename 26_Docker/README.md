# **Rangkuman Docker**

## **Docker**

- ### Perbedaan Container dengan VM

  Container adalah sebuah proses dengan file system dan resource yang terisolasi sedangkan VM adalah sebuah proses yang menjalankan sebuah OS. Container berjalan diatas OS yang sudah ada sedangkan VM menjalankan OS yang baru. Container lebih ringan dibandingkan VM karena tidak perlu menjalankan OS yang baru. Container juga lebih cepat dibandingkan VM karena tidak perlu menjalankan OS yang baru. Container juga lebih efisien dibandingkan VM karena tidak perlu menjalankan OS yang baru. Container juga lebih aman dibandingkan VM karena tidak perlu menjalankan OS yang baru.

- ### Docker Infrastructure

  - Client :
    - docker build
    - docker pull
    - docker run
  - Docker Host :
    - Docker Daemon: containers, images
  - Docker Registry : Docker Hub

- ### Dockerfile
  file yang mendeskripsikan langkah-langkah untuk membuild image
- ### Docker Image
  file yang berisi semua yang dibutuhkan untuk menjalankan sebuah aplikasi. Image ini bisa dijalankan menjadi container. dan yang akan di push ke docker registry dan dari docker registry yang akan di pull ke docker host.
- ### Docker Container
  proses yang berjalan diatas OS yang sudah ada. Container ini bisa dijalankan menjadi container.
- ### **Docker Volume**
  Tempat penyimpanan data yang bisa diakses oleh container, sehingga dengan menggunakan volume container bisa dihapus tanpa menghapus penyimpanan data yang ada di dalamnya.
- ### **Docker Network**
  Sebuah virtual network yang bisa digunakan untuk menghubungkan container-container yang berbeda. Docker Network juga bisa digunakan untuk menghubungkan container dengan host.

## **Strategi Deployment**

- ### Big-Bang Deployment (Replace Deployment)
- ### Rollout Deployment
- ### Blue/Green Deployment
- ### Canary Deployment
