# 🎬 Backend Sistem Informasi Bioskop 

## Archie Vian Nizam Efendi
## 📌 Deskripsi Proyek
Sistem informasi bioskop ini dikembangkan menggunakan Golang dengan arsitektur **Clean Architecture** dengan framework **Gin** dan ORM **GORM**. Sistem Backend ini menyediakan berbagai **API** memungkinkan admin bioskop untuk mengelola studio, film, jadwal, promo atau diskon, serta transaksi tiket, dan pengguna untuk melihat jadwal film serta membeli tiket secara online

### 🎯 Fitur Utama
1. **Manajemen Studio** (Admin)
   - Menambahkan, memperbarui, dan menghapus studio.
2. **Manajemen Film** (Admin)
   - Menambahkan, memperbarui, dan menghapus film.
3. **Manajemen Jadwal** (Admin)
   - Menetapkan jadwal tayang film di studio.
4. **Manajemen Tiket** (Pengguna)
   - Melihat jadwal film dan memilih kursi.
5. **Manajemen Transaksi** (Pengguna)
   - Pembelian tiket dan sistem pembayaran.
6. **Notifikasi**
   - Memberikan notifikasi kepada pengguna terkait jadwal film dan transaksi.
7. **Laporan Penjualan**
   - Menyediakan laporan penjualan tiket secara periodik harian dan bulanan.

## 📡 API Documentation

Berikut adalah daftar API yang telah diimplementasikan dalam sistem ini:

### 🔹 **Auth API**
| Method | Endpoint          | Deskripsi                         |
|--------|------------------|----------------------------------|
| POST   | `/admin/signup`   | Registrasi admin                     |
| POST   | `/user/signup`    | Registrasi user 
| POST   | `/admin/signin` | Admin masuk                    |
| POST   | `/user/signin` | User masuk        |
| GET   | `/admin/viewall` | Melihat semua admin        |
| GET   | `/user/viewall` | Melihat semua user        |

### 🔹 **Studio API**
| Method | Endpoint       | Deskripsi                        |
|--------|---------------|---------------------------------|
| POST   | `admin/studios`     | Admin menambahkan studio baru        |
| GET    | `admin/studios/viewall`     | Admin melihat daftar studio      |
| PUT    | `admin/studios/:id` | Admin memperbarui informasi studio  |
| DELETE | `admin/studios/:id` | Admin menghapus studio              |

### 🔹 **Film API**
| Method | Endpoint    | Deskripsi                     |
|--------|------------|------------------------------|
| POST   | `admin/films`    | Menambahkan film baru       |
| GET    | `admin/films/viewall`    | Admin mendapatkan daftar film      |
| PUT    | `admin/films/:id`| Admin memperbarui informasi film |
| DELETE | `admin/films/:id`| Admin menghapus film             |

### 🔹 **Schedule API**
| Method | Endpoint         | Deskripsi                              |
|--------|-----------------|--------------------------------------|
| POST   | `admin/schedules`     | Admin menambahkan jadwal tayang film     |
| GET    | `admin/schedules/viewall`     | Admin mendapatkan daftar jadwal film      |
| GET    | `user/schedules/viewall`     | User mendapatkan daftar jadwal film      |
| PUT    | `admin/schedules/:id` | Admin memperbarui jadwal film            |
| DELETE | `admin/schedules/:id` | Admin menghapus jadwal film              |
| POST   | `admin/schedules/promo/:id`     | Admin menambahkan promo jadwal tayang film     |

### 🔹 **Ticket API**
| Method | Endpoint          | Deskripsi                        |
|--------|------------------|---------------------------------|
| GET    | `user/tickets/:id`    | User melihat detail tiket           |
| POST   | `user/tickets/book`    | User membooking tiket                  |

### 🔹 **Transaction API**
| Method | Endpoint              | Deskripsi                            |
|--------|----------------------|-------------------------------------|
| POST   | `user/transactions/pay/:id`    | User melakukan pembayaran tiket         |
| GET    | `admin/transactions/viewall`        | Admin melihat riwayat transaksi pengguna |
| GET    | `admin/transactions/viewfilm/daily`        | Admin melihat penjualan film harian |
| GET    | `admin/transactions/viewfilm/monthly`        | Admin melihat penjualan film bulanan |
| GET    | `admin/transactions/viewstudio/daily`        | Admin melihat penjualan studio harian |
| GET    | `admin/transactions/viewstudio/monthly`        | Admin melihat penjualan studio bulanan |

### 🔹 **Notification API**
| Method | Endpoint             | Deskripsi                              |
|--------|---------------------|--------------------------------------|
| GET    | `user/notifications/viewall`     | User mendapatkan notifikasi      |

## 🏛️ Project Pattern

Sistem ini menggunakan **Clean Architecture** untuk memastikan modularitas dan skalabilitas. Berikut adalah struktur proyek :

```
/trainingbackenddot  
  ├── config/                    # Konfigurasi database
  │   ├── database.go    
  ├── router/                    # Setup routing API
  │   ├── router.go           
  ├── domain/                   # Model data
  │   ├── admin.go               
  │   ├── user.go                
  │   ├── film.go
  │   ├── schedule.go
  │   ├── studio.go
  │   ├── ticket.go
  │   ├── transaction.go
  ├── infrastructure/           # Layer repository untuk komunikasi dengan database
  │   ├── db/
  │   │   ├── migration.go
  │   │   ├── admin_repository.go
  │   │   ├── user_repository.go
  │   │   ├── film_repository.go
  │   │   ├── schedule_repository.go
  │   │   ├── studio_repository.go
  │   │   ├── ticket_repository.go
  │   │   ├── transaction_repository.go  
  ├── usecase/                    # Layer business logic
  │   ├── admin_usecase.go        
  │   ├── user_usecase.go         
  │   ├── film_usecase.go
  │   ├── schedule_usecase.go
  │   ├── studio_usecase.go
  │   ├── ticket_usecase.go
  │   ├── transaction_usecase.go    
  ├── interface/                  # Layer controller / handler
  │   ├── http/                    
  │   │   ├── admin_handler.go      
  │   │   ├── user_handler.go       
  │   │   ├── film_handler.go
  │   │   ├── schedule_handler.go
  │   │   ├── studio_handler.go
  │   │   ├── ticket_handler.go
  │   │   ├── transaction_handler.go     
  ├── main.go    # Entry point aplikasi                        
  ├── go.mod
  ├── go.sum
  ├── .env
  ├── .gitignore  
```

## 🔥 Alasan Menggunakan Clean Architecture
1. **Modularitas**: Memisahkan logika bisnis, data, dan handler membuat kode lebih mudah dikelola.
2. **Scalability**: Struktur ini memungkinkan sistem berkembang tanpa mengubah banyak bagian kode.
3. **Maintainability**: Memudahkan debugging dan testing, karena setiap layer memiliki tanggung jawab sendiri.
4. **Reusability**: Logika bisnis dapat digunakan kembali tanpa ketergantungan pada implementasi lainnya.

## 📌 Kesimpulan
✅ **Fitur utama** sudah diimplementasikan:
✔ Manajemen studio, film, jadwal, tiket, transaksi
✔ Promo & diskon untuk tiket
✔ Laporan penjualan harian dan bulanan
✔ Notifikasi pengguna

🚀 **Saran pengembangan selanjutnya:**
- Implementasi **middleware** untuk logging & monitoring
- Integrasi **payment gateway** untuk metode pembayaran yang lebih luas
- Penambahan fitur **seat selection dengan tampilan real-time**
- **Testing & deployment automation** untuk CI/CD

### 🚀 Pengembangan Selanjutnya:
1. **Integrasi Payment Gateway** 💳  
   - Saat ini transaksi masih dummy, perlu integrasi pembayaran.
2. **Penggunaan Middleware untuk Keamanan Data** 🔒  
   - Implementasi rate limiting dan peningkatan keamanan API.


## 🚀 Cara Menjalankan Proyek
### 🔹 **Instalasi & Setup**
1. Clone repository ini:
   ```bash
   git clone https://github.com/archieVian8/trainingbackenddot.git
   cd trainingbackenddot
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Buat file `.env` untuk konfigurasi database:
   ```env
   DB_USER=youruser
   DB_PASSWORD=yourpassword
   DB_NAME=yourdbname
   DB_HOST=localhost
   DB_PORT=PORT
   ```
4. Jalankan aplikasi:
   ```bash
   go run main.go
   ```
5. Akses API di: `http://localhost:3000`