# Backend Ticketing Event API
Backend Ticketing Event API adalah sebuah RESTful API yang dibangun menggunakan
Golang dan Fiber untuk mendukung sistem pemesanan tiket event secara online.
Sistem ini memungkinkan admin untuk mengelola event dan user untuk melakukan
booking tiket dengan aman tanpa risiko overselling.

---

## 🚀 Tech Stack
- Golang
- Fiber
- GORM
- MySQL
- JWT (JSON Web Token)
- Postman (API Documentation)

---

## 📦 Fitur Utama

### 🔐 Autentikasi
- Register User
- Login User / Admin (JWT)
- Logout

### 📅 Event
- Create Event (Admin)
- Update Event (Admin)
- Delete Event (Admin)
- Get All Event (Public)
- Get Event by ID (Public)

### 🎟️ Booking
- Create Booking (User)
- Get My Booking (User)
- Delete My Booking (User)
- Get All Booking (Admin)
- Get Booking per Event (Admin)

---

## 🔐 Authentication & Authorization
API ini menggunakan JWT (JSON Web Token) sebagai mekanisme autentikasi.
Token dikirim melalui HTTP Header dengan format:

Sistem menerapkan role-based access control:
- **Admin**: mengelola event dan melihat seluruh booking
- **User**: melakukan booking dan melihat booking miliknya sendiri

---

## 🛡️ Keamanan & Validasi
- Password disimpan menggunakan hashing (bcrypt)
- Booking menggunakan database transaction untuk mencegah overselling tiket
- Event tidak dapat dihapus apabila sudah memiliki data booking
- Validasi JWT dilakukan melalui middleware

---

## 📄 API Documentation
Dokumentasi lengkap API tersedia dan dapat diakses secara publik melalui Postman:

🔗 **Postman API Documentation**  
* Autentikasi API: https://documenter.getpostman.com/view/44006656/2sBXVbGDER
* Booking API: https://documenter.getpostman.com/view/44006656/2sBXVbGDES
* Event API: https://documenter.getpostman.com/view/44006656/2sBXVbGDET

## ▶️ Cara Menjalankan Project
1. Clone repository
```bash
https://github.com/mauinGit/BE_Program_MaulanaAdiatma.git
```

2. Install dependency
```bash
go mod tidy
```

3. Jalankan sistem
```bash
go run main.go
```

4. Sistem akan berjalan pada localhost kamu
```bash
http://127.0.0.1:3000
```

Siap 👍 **simpel, ringkas, dan siap copy–paste ke GitHub**.
Ini **STRUKTUR FOLDER LENGKAP** sesuai **punya kamu**, tanpa penjelasan panjang.

---
## 📁 Struktur Folder

```bash
BE_Program_MaulanaAdiatma/
├── assets/
│   └── *.jpg / *.png / *.webp
├── config/
│   └── env.go
├── controller/
│   ├── user.go
│   ├── event.go
│   └── booking.go
├── database/
│   └── database.go
├── middleware/
│   └── auth.go
├── model/
│   ├── user.go
│   ├── event.go
│   └── booking.go
├── route/
│   ├── routes.go
│   ├── user.go
│   ├── event.go
│   └── booking.go
├── util/
│   ├── jwt.go
│   └── cover.go
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```


