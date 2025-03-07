```markdown
# 🚀 Technical Test API - Go (Gin, GORM, PostgreSQL, JWT, Redis)

Repo ini digunakan untuk kebutuhan technical test RESTful API yang dibangun dengan **Go (Gin)** menggunakan **GORM** untuk ORM, **JWT untuk autentikasi**, dan **Redis untuk session management dan caching**.

---

## 📌 Fitur Utama
✅ **Implement CRUD Tasks menggunakan database PostgreSQL dengan GORM**
✅ **Implement validation**
✅ **Implement filter get all tasks by query parameters**
✅ **Logging error untuk debugging**  
✅ **Implement autentikasi & otorisasi dengan JWT**
✅ **Implement concurrency**  
✅ **Feature Test menggunakan `go test` untuk validasi API**  

---

## 📦 1. Instalasi & Persiapan, Asumsi sudah ada PostgreSQL dan Redis
### **Clone Repository**
```sh
git clone https://github.com/yasseryazid/technical-test.git
cd technical-test
```

#### **b) Persiapan Database**
```sh
Buat Database dengan nama technical_test
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
Rename file **`.env-example`** menjadi **`.env`** serta sesuaikan dengan environment Anda:
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
REDIS_PASSWORD=YOUR_REDIS_PASSWORD
```

---

## 🚀 2. Menjalankan API
### **Start API**
```sh
go run cmd/main.go
```
**Main.go akan melakukan migration database ketika pertama kali, serta menjalankan server by default di:** `http://localhost:3000`

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

#### **Query Parameters Untuk Get All Tasks**
| Parameter  | Tipe   | Deskripsi |
|------------|--------|-------------|
| `status`   | `string` | Filter tasks berdasarkan status (`pending` / `completed`) |
| `page`     | `int`    | Nomor halaman untuk pagination (default: `1`) |
| `limit`    | `int`    | Jumlah tasks per halaman (default: `10`) |
| `search`   | `string` | Cari task berdasarkan `title` atau `description` |

**Contoh request dengan JWT Token:**
```sh
curl -X GET http://localhost:3000/api/tasks \
     -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response get all tasks dengan JWT Token:**
```sh
{
  "tasks": [
    {
      "id": 1,
      "title": "Meeting with Client",
      "description": "Discuss project scope",
      "status": "pending",
      "due_date": "2025-04-01"
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 2,
    "total_tasks": 10
  }
}
```

**Response get all tasks jika tanpa token:**
```sh
{
  "error": "Authorization header required"
}
```

---

#### **Query Parameters Create Task**
#### **Endpoint**
```sh
POST /api/tasks

```
#### **Header**
```sh
{
  "Authorization": "Bearer {{YOUR_JWT_TOKEN}}",]
}
```


#### **Payload**
```sh
{
  "title": "New Task",
  "description": "Complete documentation",
  "status": "pending",
  "due_date": "2025-04-01"
}
```

#### **Response**
```sh
{
    "message": "Task created successfully",
    "task": {
        "title": "Create new Task",
        "description": "Create new description",
        "status": "completed",
        "due_date": "2025-05-01"
    }
}
```

---

### **Get Task by ID**
```http
GET /api/tasks/:id
```

#### **Header**
```sh
{
  "Authorization": "Bearer {{YOUR_JWT_TOKEN}}",]
}
```

**Response:**
```sh
{
  "id": 2,
  "title": "Task 2",
  "description": "Description for Task 2",
  "status": "completed",
  "due_date": "2025-03-12T00:00:00Z"
}
```

---

### **4️⃣ Update Task**
```http
PUT /api/tasks/:id
```
#### **Header**
```sh
{
  "Authorization": "Bearer {{YOUR_JWT_TOKEN}}",]
}
```

**Payload:**
```sh
{
  "title": "Update Task",
  "description": "Update description",
  "status": "completed",
  "due_date": "2025-05-01"
}
```
**Response:**
```sh
{
    "message": "Task updated successfully",
    "task": {
        "title": "Update Task",
        "description": "Update description",
        "status": "completed",
        "due_date": "2025-05-01"
    }
}
```

---

### **Delete Task**
```http
DELETE /api/tasks/:id
```
#### **Header**
```sh
{
  "Authorization": "Bearer {{YOUR_JWT_TOKEN}}",]
}
```

**Response:**
```sh
{
    "message": "Task deleted successfully"
}
```

---

## 🔍 5. Menjalankan Test  
Jalankan **unit test dan integration test** dengan perintah:
```sh
go test ./tests -v
```
✔️ **Expected Output:**  
```
=== RUN   TestCreateTask
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

2025/03/07 08:34:26 [V] .env file loaded successfully
2025/03/07 08:34:26 [V] Connected to the database
2025/03/07 08:34:26 [V] Task created successfully: ID 14
[GIN] 2025/03/07 - 08:34:26 | 201 |   11.914292ms |                 | POST     "/api/tasks"
--- PASS: TestCreateTask (0.06s)
=== RUN   TestGetTaskByID
2025/03/07 08:34:26 [V] .env file loaded successfully
2025/03/07 08:34:26 [V] Task retrieved: ID 14
[GIN] 2025/03/07 - 08:34:26 | 200 |    2.204958ms |                 | GET      "/api/tasks/14"
--- PASS: TestGetTaskByID (0.00s)
=== RUN   TestUpdateTask
2025/03/07 08:34:26 [V] .env file loaded successfully
2025/03/07 08:34:26 [V] Task updated successfully: ID 14
[GIN] 2025/03/07 - 08:34:26 | 200 |    2.656084ms |                 | PUT      "/api/tasks/14"
--- PASS: TestUpdateTask (0.00s)
=== RUN   TestDeleteTask
2025/03/07 08:34:26 [V] .env file loaded successfully
2025/03/07 08:34:26 [V] Task deleted successfully: ID 14
[GIN] 2025/03/07 - 08:34:26 | 200 |       905.5µs |                 | DELETE   "/api/tasks/14"
--- PASS: TestDeleteTask (0.00s)
=== RUN   Test_CreateTask
--- PASS: Test_CreateTask (0.00s)
=== RUN   Test_GetTaskByID
--- PASS: Test_GetTaskByID (0.00s)
=== RUN   Test_UpdateTask
--- PASS: Test_UpdateTask (0.00s)
=== RUN   Test_DeleteTask
--- PASS: Test_DeleteTask (0.00s)
=== RUN   TestGetTasks
--- PASS: TestGetTasks (0.00s)
=== RUN   TestGetTaskByID_NotFound
--- PASS: TestGetTaskByID_NotFound (0.00s)
PASS
ok      github.com/yasseryazid/technical-test/tests     0.440s
PASS
```

---

## 📊 6. Logging Setiap Error untuk Debugging
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

## ⚡ 7. Implementasi Concurrency
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
✅ **Diatas adalah sample implementasi concurrency!**

---

## 🎯 Summary
✅ **Clone repository & setup environment**
✅ **Menjalankan API dengan Redis & PostgreSQL**  
✅ **Menggunakan autentikasi JWT dengan Redis**  
✅ **Implement CRUD Tasks serta penerapan filternya dengan query param dan implementasi validation**  
✅ **Menjalankan feature test untuk memastikan API berjalan dengan baik**  
✅ **Logging setiap error untuk debugging lebih mudah**  
✅ **Menggunakan concurrency**  

```