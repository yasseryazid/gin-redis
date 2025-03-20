# 🚀 **Go (Gin, GORM, PostgreSQL, JWT, Redis)**  

This repository is built using **Go** with **Gin Framework**, **GORM**, **PostgreSQL**, **JWT** for authentication, and **Redis** for session management & caching.  

---

## 📌 **Key Features**  

✔️ **CRUD Tasks** – Uses **PostgreSQL** with **GORM** for data management  
✔️ **Data Validation** – Prevents invalid input to maintain data integrity  
✔️ **Filter Get All Tasks** – Supports **query parameters** for searching & filtering data (see Query Parameters for Get All Tasks)  
✔️ **Error Logging** – Simplifies debugging by automatically recording errors  
✔️ **Authentication & Authorization** – Uses **JWT (JSON Web Token)** for secure access control  
✔️ **Concurrency** – Optimizes API performance with **parallel processing**  
✔️ **Feature Testing** – Uses **`go test`** to ensure API reliability  

---

## 📌 **1. Installation**  

To install this project, make sure you have Go (at least the latest stable version), PostgreSQL, Redis, and Git installed on your system.

### **a) Project Setup**
#### **Clone Repository**
```sh
git clone https://github.com/yasseryazid/gin-redis.git
cd gin-redis
```

#### **Database Preparation**
```sh
Create a Database named technical_test
```

---

### **b) Setup Redis**
#### **Start Redis**
```sh
redis-server
```
#### **Check if Redis is Running**
```sh
redis-cli ping
```
If the output is `PONG`, Redis is running properly.

---

### **c) Configure `.env`**
Rename **`.env-example`** to **`.env`** and adjust it according to your environment:
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

## 🚀 2. Running the API
### **Start API**
```sh
go run cmd/main.go
```
**Main.go will perform database migration on the first run and start the server by default at:** `http://localhost:3000`

---

## 🔑 3. Authentication (JWT)**  

### **🔐 JWT as Authentication & Authorization**  
To access the **Tasks** endpoint, authentication using **JWT (JSON Web Token)** is required. A token is **generated during registration or login** and used for authorization in every request to protected endpoints.  

---

### **📌 JWT Mechanism in This API**  

✅ **JWT Token is Used for Authentication & Authorization**  
✅ **Token Has an Expiration (`exp`) of 24 Hours**  
✅ **JWT Payload Contains:**
   - **`user_id`** → User identifier  
   - **`username`** → User's name  
   - **`exp`** → Token expiration time  

✅ **JWT is Stored in Redis with a TTL (`Time-To-Live`) of 24 Hours**  
✅ **On Logout, the Token is Removed from Redis**  

---

### **📌 JWT Usage Flow in the API**
1. **User Registers/Login**  
   - The system generates a JWT valid for **24 hours**  
   - The token is **stored in Redis** with a **TTL of 24 hours**  

2. **User Uses JWT to Access Protected Endpoints**  
   - Every request to the **Tasks API** must include a **Bearer Token** in the **Authorization Header**  
   - The system **validates the token** before granting access  

3. **On Logout, Token is Removed from Redis**  
   - The stored token in Redis will be deleted, requiring **users to log in again** to get a new token  

---

## 📌 4. API Endpoints  

### **Auth**
| Method | Endpoint       | Description |
|--------|--------------|-------------|
| `POST`  | `/api/register`  | Create a new account |
| `POST` | `/api/login`  | Login with an existing account |
| `POST` | `/api/logout`  | Logout |

Use **JWT Token** to access the **tasks endpoints**.
### **Tasks (Protected)**
| Method | Endpoint       | Description |
|--------|--------------|-------------|
| `GET`  | `/api/tasks`  | Get all tasks |
| `POST` | `/api/tasks`  | Create a task |
| `GET`  | `/api/tasks/:id` | Get task by ID |
| `PUT`  | `/api/tasks/:id` | Update task |
| `DELETE` | `/api/tasks/:id` | Delete task |

#### **Query Parameters for Get All Tasks**
| Parameter  | Type   | Description |
|------------|--------|-------------|
| `status`   | `string` | Filter tasks by status (`pending` / `completed`) |
| `page`     | `int`    | Page number for pagination (default: `1`) |
| `limit`    | `int`    | Number of tasks per page (default: `10`) |
| `search`   | `string` | Search tasks by `title` or `description` |

---

## 🔍 5. Running Tests  
Run **unit tests and integration tests** using:
```sh
go test ./tests -v
```

---

## 📊 6. Logging Every Error for Debugging
All API errors will **be logged** using `log.Println(err)` to simplify debugging.  

For example, in the authentication middleware:
```go
if err != nil {
    log.Println("[X] Error validating JWT:", err)
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
    c.Abort()
    return
}
```

---

## ⚡ 7. Implementing Concurrency
**Concurrency is used in:**
- **Get Tasks** → Uses `sync.WaitGroup` for parallel queries (task list & total count)
- **Processing Asynchronous Task Handling**

An example can be found in `repositories/task_repository.go`:
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

---

## 🎯 Summary
✅ **Clone the repository & setup environment**  
✅ **Run the API with Redis & PostgreSQL**  
✅ **Use JWT authentication with Redis**  
✅ **Implement CRUD Tasks with query filtering and validation**  
✅ **Run feature tests to ensure API functionality**  
✅ **Log errors for easier debugging**  
✅ **Use concurrency**  

---

This translated README preserves the structure, formatting, and technical details while making it clear and professional in English. Let me know if you need any refinements! 🚀
