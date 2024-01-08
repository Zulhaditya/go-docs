# PENGENALAN VALIDATION
- Saat membuat aplikasi, validasi adalah salah satu hal yang selalu dibuat
- Validasi dilakukan untuk memastikan bahwa data yang diproses sudah benar
- Validasi adalah sesuatu yang wajib dilakukan saat pembuatan aplikasi, agar kesalahan pada data bisa ditemukan secepat mungkin sebelum data tersebut diproses

# TEMPAT MELAKUKAN VALIDATION
- Validasi sering dilakukan di banyak bagian pada aplikasi seperti,
- **Web**, validasi request dari pengguna
- **Business logic**, validasi data
- **Database**, validasi constraint
- Pada beberapa bagian kadang menggunakan validasi yang sama. oleh karena itu pembuatan validasi secara manual sangat memakan waktu, dan kesalahan sedikit bisa menyebabkan validasi tidak konsisten

# KENAPA BUTUH VALIDASI?
- Sederhananya untuk memastikan request atau data yang dikirimkan oleh pengguna sudah sesuai dengan yang kita inginkan
- **Never trust user input!**

# MANUAL VALIDATION
- Saat melakukan validasi, biasanya melakukan validasi secara manual
- Rata-rata, validasi manual menggunakan **if statement**
- Semakin banyak validasi yang diperlukan, semakin banyak if statement yang harus dibuat

# VALIDATION LIBRARY
- Penggunaan library untuk melakukan validasi sangat direkomendasikan
- Hal ini agar kode validasi bisa lebih mudah diterapkan dan rapi antar programmer
- kita bisa menggunakan library **validator package**

# VALIDATOR PACKAGE
- Validator package adalah opensource library untuk melakukan validation di Golang
- Validator package memiliki banyak sekali fitur yang bisa kita gunakan untuk mempermudah kita melakukan validasi
- Lihat di: https://github.com/go-playground/validator

# MEMBUAT PROJECT
- Buat folder belajar-golang-validation
- Buat sebagai **go module:**
```bash
go mod init belajar_golang_validation
```
- Tambahkan library **validator** package:
```bash
go get github.com/go-playground/validator/v10
```