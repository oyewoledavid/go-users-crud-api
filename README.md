# 🧩 Go User Management API

This is a simple RESTful API built with **Go (Golang)** using the **Gin web framework** and **GORM ORM**. It connects to a MySQL database and allows you to perform CRUD operations on a list of users.

---

## 🚀 Features

- Get all users  
- Get user by ID  
- Get user by name  
- Add new user  
- Update user details  
- Delete user  
- Uses `.env` file for configuration  
- Auto-seeds default data if DB is empty

---

## 📂 Project Structure

```
.
├── README.md              # Project documentation
├── main.go                # Entry point of the application
├── db                 # Configuration files and database connection
│   └── db.go
├── models                 # Data models
│   └── user.go
├── handlers                 # API route logics definitions
│   └── helper.go
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
└── .env                   # Environment variables
```
---

## 🛠️ Technologies Used

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [MySQL](https://www.mysql.com/)
- [godotenv](https://github.com/joho/godotenv)

---

## ⚙️ Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/go-user-api.git
cd go-user-api  
```
###  2. Create a .env File  
```dotenv
DB_USER=your_mysql_user
DB_PASS=your_mysql_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=your_database_name
```  
⚠️ Ensure the database exists in your MySQL server. The table will be created automatically by GORM.  
### 3. Install Dependencies
```bash
go run main.go
```
### 4. Run the Application
```bash
go run main.go
```

### 🧪 API Endpoints
| Method | Endpoint                | Description       |
| ------ | ----------------------- | ----------------- |
| GET    | `/users`                | Get all users     |
| GET    | `/users/id/:id`         | Get user by ID    |
| GET    | `/users/name/:fullname` | Get user by name  |
| POST   | `/users`                | Add new user      |
| PUT    | `/users/:id`            | Update user by ID |
| DELETE | `/users/:id`            | Delete user by ID |

