# **Rangkuman Introduction to RESTful API**

## **API (Application Programming Interface)**

- Sebuah interface yang memungkinkan dua aplikasi untuk berkomunikasi satu sama lain.
- Integrasi API:
  - Dari Frontend ke Backend, seperti: dari web ke service
  - Dari Backend ke Backend, seperti: dari service ke service
- Salah satu contoh API adalah REST API (Representational State Transfer)
  - REST API adalah API yang menggunakan REST (Representational State Transfer) sebagai arsitektur utamanya.
- API tools
  - Postman
  - Katalon
  - Apache JMeter
  - SoapUI
  - Insomnia
  - etc.

## **Design REST API**

- Menggunakan kata benda dibandingkan dengan kata kerja
- Menamai dengan menggunakan kata benda jamak
- Menggunakan percabangan untuk memperlihatkan relasi atau hubungan antar resource
- Menggunakan format response JSON
- Filtering, Sorting, Paging, dan Field Selection
  - Filtering
    - GET /users?name=John
  - Sorting
    - GET /users?sort=name
  - Paging
    - GET /users?page=2&size=10
  - Field Selection
    - GET /users?fields=name,age
- Menghandle garis miring trailing dengan baik
  - POST: /users
  - POST: /users/
- Menggunakan versioning
  - GET: /v1/users
  - GET: /v2/users
- Membuat API documentation, seperti
  - https://mailchimp.com/developer/marketing/docs/fundamentals/
  - https://docs.github.com/en/rest/guides/getting-started-with-the-rest-api
  - https://developer.twitter.com/en/docs/twitter-api/v1/tweets/search/overview
  - [etc](https://github.com/public-apis/public-apis)

## **Cara Implementasi REST API dengan Go**

- Menggunakan package build-in:
  - `net/http` : untuk membuat server dan mengconsume API
  - `encoding/json` : untuk menghandle JSON
- Menggunakan third party lib:
  - [echo](https://echo.labstack.com/)
  - [cobra](https://github.com/spf13/cobra)
  - [gorm](https://gorm.io/docs/)
  - [logrus](https://github.com/sirupsen/logrus)
  - go kit
  - [etc](https://github.com/go-kit/kit)

## **Reference**

- [Awesome Package Golang](https://awesome-go.com/)
- [Check Third Party Lib](https://osv.dev/)
