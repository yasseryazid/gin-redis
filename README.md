# 🚀 Technical Test API - Go (Gin, GORM, PostgreSQL, JWT, Redis)

Repo ini digunakan untuk kebutuhan Technical Test.
Anda bisa clone dan install untuk mencoba.

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

### **Instalasi**
#### **a) Clone Repository**
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
**JWT sebagai Authentication and Authorization**
Saat ini endpoint tasks statusnya protected, sehingga untuk menggunkannya dibutuhkan token yang akan digenerate ketika proses register atau login.
Setiap generate, expiration time pada token adalah 24 jam.
Payload pada JWT saat ini terdiri dari user_id, username, dan exp sebagai expiration time.
JWT yang berhasil di generate akan disimpan ke Redis dengan TTL selama 24 jam.
Ketika logout, token pada redis akan dihapus juga.

---

## 📌 4. Endpoints API  

### **Auth**
| Method | Endpoint       | Deskripsi |
|--------|--------------|-------------|
| `POST`  | `/api/register`  | Buat akun baru |
| `POST` | `/api/login`  | Login dengan akun yang sudah ada |
| `POST` | `/api/logout`  | Logout akun |

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

### **Auth Endpoint**
<details>
  <summary>📌 1. Register</summary>

```sh
POST /api/register

```
#### **Payload**
```sh
{
    "username":"user", 
    "password":"password"
}
```

#### **Response**
```sh
{
    "message": "User registered successfully"
}
```
</details>

<details>
  <summary>📌 2. Login</summary>

```sh
POST /api/login

```
#### **Payload**
```sh
{
    "username":"user", 
    "password":"password"
}
```

#### **Response**
```sh
{
    "token": "{{TOKEN_GENERATED}}"
}

```
</details>

<details>
  <summary>📌 3. Logout</summary>

```sh
POST /api/logout

```
#### **Header**
```sh
{
  "Authorization": "Bearer {{YOUR_JWT_TOKEN}}",]
}
```

#### **Response**
```sh
{
    "message": "Logged out successfully"
}
```
</details>

---

### **📝 All Tasks Endpoint**

<details>
  <summary>📌 1. Get All Tasks</summary>

#### **Request**
```http
GET /api/tasks
```

#### **Header**
```json
{
  "Authorization": "Bearer {{YOUR_JWT_TOKEN}}"
}
```

#### **Response**
```json
{
  "tasks": [
    {
      "id": "1",
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

#### **Response jika tanpa token**
```json
{
  "error": "Authorization header required"
}
```
</details>

<details>
  <summary>📌 2. Create Task</summary>

#### **Request**
```http
POST /api/tasks
```

#### **Header**
```json
{
  "Authorization": "Bearer {{YOUR_JWT_TOKEN}}"
}
```

#### **Payload**
```json
{
  "title": "New Task",
  "description": "Complete documentation",
  "status": "pending",
  "due_date": "2025-04-01"
}
```

#### **Response**
```json
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
</details>


<details>
  <summary>📌 3. Get Task by ID</summary>

#### **Request**
```http
GET /api/tasks/:id
```

#### **Header**
```json
{
  "Authorization": "Bearer {{YOUR_JWT_TOKEN}}"
}
```

#### **Response**
```json
{
  "id": "2",
  "title": "Task 2",
  "description": "Description for Task 2",
  "status": "completed",
  "due_date": "2025-03-12T00:00:00Z"
}
```
</details>

<details>
  <summary>📌 4. Update Task</summary>

#### **Request**
```http
PUT /api/tasks/:id
```

#### **Header**
```json
{
  "Authorization": "Bearer {{YOUR_JWT_TOKEN}}"
}
```

#### **Payload**
```json
{
  "title": "Update Task",
  "description": "Update description",
  "status": "completed",
  "due_date": "2025-05-01"
}
```

#### **Response**
```json
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
</details>


<details>
  <summary>📌 5. Delete Task</summary>

#### **Request**
```http
DELETE /api/tasks/:id
```

#### **Header**
```json
{
  "Authorization": "Bearer {{YOUR_JWT_TOKEN}}"
}
```

#### **Response**
```json
{
    "message": "Task deleted successfully"
}
```
</details>

---



## 🔍 5. Run Test  
Jalankan **unit test dan integration test** dengan perintah:
```sh
go test ./tests -v
```

<details>
  <summary>📌 Expected Result</summary>

```sh
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

</details> ```
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