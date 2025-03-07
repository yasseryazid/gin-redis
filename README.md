Berikut adalah **README.md** yang telah diperbarui dengan tambahan **feature test, logging error, dan concurrency** pada kesimpulan. 🚀  

---

```markdown
# 🚀 Technical Test API - Go (Gin, GORM, PostgreSQL, JWT, Redis)

Technical Test API adalah **RESTful API** yang dibangun dengan **Go (Gin)** menggunakan **GORM** untuk ORM, **JWT untuk autentikasi**, dan **Redis untuk session management dan caching**.

---

## 📌 Fitur Utama
✅ **Autentikasi & Otorisasi dengan JWT + Redis**  
✅ **CRUD Tasks dengan caching Redis**  
✅ **Middleware Proteksi JWT**  
✅ **Rate Limiting dengan Redis**  
✅ **Database PostgreSQL dengan GORM**  
✅ **Logging setiap error untuk debugging**  
✅ **Concurrency pada query & proses parallel task handling**  
✅ **Feature Test menggunakan `go test` untuk validasi API**  

---

## 📦 1. Instalasi & Persiapan  
### **Clone Repository**
```sh
git clone https://github.com/yasseryazid/technical-test.git
cd technical-test
```

### **Setup Database (PostgreSQL)**
#### **a) Jalankan PostgreSQL**
```sh
sudo systemctl start postgresql  # Untuk Linux
brew services start postgresql   # Untuk macOS (Homebrew)
```

#### **b) Buat Database**
```sh
psql -U postgres
CREATE DATABASE technical_test;
```

---

### **Setup Redis**
#### **a) Jalankan Redis**
```sh
redis-server
```
#### **b) Cek Redis Berjalan**
```sh
redis-cli ping
```
Jika outputnya `PONG`, Redis berjalan dengan baik.

---

### **Konfigurasi `.env`**
Buat file **`.env`** dan tambahkan:
```ini
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=technical_test
DB_PORT=5432
DB_SSLMODE=disable

JWT_SECRET=mysecretkey

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD={{YOUR_REDIS_PASSWORD}}
```

---

## 🚀 2. Menjalankan API  
### **Jalankan Migrasi Database**
```sh
go run cmd/migrate.go
```
### **Start API**
```sh
go run cmd/main.go
```
**API berjalan di:** `http://localhost:3000`

---

## 🔑 3. Autentikasi (JWT)  
### **Register**
```sh
curl -X POST http://localhost:3000/api/register \
     -H "Content-Type: application/json" \
     -d '{"username":"testuser", "password":"password"}'
```

### **Login & Dapatkan Token**
```sh
curl -X POST http://localhost:3000/api/login \
     -H "Content-Type: application/json" \
     -d '{"username":"testuser", "password":"password"}'
```
✔️ **Response:**  
```json
{
    "token": "YOUR_JWT_TOKEN"
}
```

### **Logout**
```sh
curl -X POST http://localhost:3000/api/logout \
     -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## 📌 4. Endpoints API  
Gunakan **JWT Token** untuk mengakses **endpoint tasks**.

### **Tasks (Protected)**
| Method | Endpoint       | Deskripsi |
|--------|--------------|-------------|
| `GET`  | `/api/tasks`  | Get all tasks |
| `POST` | `/api/tasks`  | Create task |
| `GET`  | `/api/tasks/:id` | Get task by ID |
| `PUT`  | `/api/tasks/:id` | Update task |
| `DELETE` | `/api/tasks/:id` | Delete task |

**Contoh request dengan JWT Token:**
```sh
curl -X GET http://localhost:3000/api/tasks \
     -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## 📜 5. Dokumentasi API
Terdapat dokumentasi lengkap untuk **register, login, CRUD tasks, dan logout**, semua sudah menggunakan **JWT dan Redis**.

Lihat **[Dokumentasi API](#📌-4-endpoints-api)** di atas untuk detailnya.

---

## 🔍 6. Menjalankan Pengujian  
Jalankan **unit test dan integration test** dengan perintah:
```sh
go test ./tests -v
```
✔️ **Expected Output:**  
```
=== RUN   TestCreateTask
--- PASS: TestCreateTask (0.10s)
=== RUN   TestGetTaskByID
--- PASS: TestGetTaskByID (0.05s)
=== RUN   TestUpdateTask
--- PASS: TestUpdateTask (0.05s)
=== RUN   TestDeleteTask
--- PASS: TestDeleteTask (0.03s)
=== RUN   TestGetAllTasks
--- PASS: TestGetAllTasks (0.10s)
PASS
```

---

## 📊 7. Logging Setiap Error untuk Debugging
Setiap error dalam API akan **tercatat dalam log** menggunakan `log.Println(err)` untuk mempermudah debugging.  

Misalnya dalam middleware autentikasi:
```go
if err != nil {
    log.Println("[X] Error validating JWT:", err)
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
    c.Abort()
    return
}
```
---

## ⚡ 8. Implementasi Concurrency untuk Optimasi API
**Concurrency digunakan pada:**
- **Get Tasks** → Menggunakan `sync.WaitGroup` untuk query paralel (task list & total count)
- **Processing Asynchronous Task Handling**

Contoh di `repositories/task_repository.go`:
```go
var wg sync.WaitGroup
errChan := make(chan error, 2)

wg.Add(2)

// Fetch tasks asynchronously
go func() {
    defer wg.Done()
    err := query.Limit(limit).Offset(offset).Find(&tasks).Error
    if err != nil {
        errChan <- err
    }
}()

// Fetch total count asynchronously
go func() {
    defer wg.Done()
    err := query.Count(&total).Error
    if err != nil {
        errChan <- err
    }
}()

wg.Wait()
close(errChan)
```
✅ **Diatas adalah sample implementasi concurrency!** 🚀

---

## 🎯 Kesimpulan
✅ **Clone repository & setup environment**  
✅ **Menjalankan API dengan Redis & PostgreSQL**  
✅ **Menggunakan autentikasi JWT dengan Redis**  
✅ **CRUD Tasks dengan caching dan rate limiting**  
✅ **Menjalankan feature test untuk memastikan API berjalan dengan baik**  
✅ **Logging setiap error untuk debugging lebih mudah**  
✅ **Menggunakan concurrency untuk mempercepat eksekusi query dan proses API**  

```


README ini **komprehensif**, menjelaskan **instalasi, API endpoints, testing, logging, dan concurrency**, cocok untuk proyek production-ready. 🚀🔥