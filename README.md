# InternGo

A comprehensive project repository for InternGo.

---

## Mục lục
- [Giới thiệu](#giới-thiệu)
- [Yêu cầu kỹ thuật](#yêu-cầu-kỹ-thuật)
- [Hướng dẫn cài đặt](#hướng-dẫn-cài-đặt)
- [Cấu trúc dự án](#cấu-trúc-dự-án)
- [Xây dựng và chạy dự án](#xây-dựng-và-chạy-dự-án)
- [Các chức năng nổi bật](#các-chức-năng-nổi-bật)
- [Cấu hình môi trường & Debug](#cấu-hình-môi-trường--debug)
- [Database & Migration](#database--migration)
- [Docker hoá dự án](#docker-hoá-dự-án)

---

## Giới thiệu

**InternGo** là dự án xây dựng hệ thống backend API sử dụng ngôn ngữ Go. Dự án áp dụng các thực hành tốt nhất về RESTful API, dependency injection, middleware, websocket, logging, multi-environment config, migration, tích hợp với Docker, sử dụng GORM cho ORM/database và hỗ trợ debug mạnh mẽ với VSCode.

---

## Yêu cầu kỹ thuật

- Go 1.18+
- Gin Framework
- GORM (với PostgreSQL hoặc MySQL)
- Docker & Docker Compose (khuyến khích)
- Delve (debugger cho Go)
- VSCode (với Go extension)
- Hệ điều hành: Linux, Windows hoặc MacOS

---

## Hướng dẫn cài đặt

1. **Clone repository**
   ```sh
   git clone https://github.com/NEIHT1612/InternGo.git
   cd InternGo
   ```

2. **Cài đặt Go dependencies**
   ```sh
   go mod tidy
   ```

3. **Cài đặt Delve để debug (tùy chọn)**
   ```sh
   go install github.com/go-delve/delve/cmd/dlv@latest
   ```

4. **Cấu hình file môi trường**
   - Copy `.env.example` thành `.env` rồi chỉnh sửa các tham số kết nối DB, JWT, v.v.
   - Sử dụng `.env.prod` cho môi trường production.

---

## Cấu trúc dự án

- `main.go`: Entry point của dự án
- `build.sh`, `run.bat`: Script build/chạy
- `.env`, `.env.prod`, `.ginconfig`: File cấu hình môi trường và Gin
- `common/`, `db/`, `models/`, `service/`, `repository/`, `route/`, `middleware/`: Source chính
- `logger/`: Cấu hình & custom log
- `scripts/`: Script hỗ trợ
- `uploads/`: Lưu file từ API

---

## Xây dựng và chạy dự án

### Build

- Build bằng script:
  ```sh
  ./build.sh
  ```
- Build thủ công:
  ```sh
  go build -o build/intern-go main.go
  ```

### Chạy

- **Trực tiếp với Go:**
  ```sh
  go run main.go
  ```
- **Chạy với Gin CLI (hot reload):**
  ```sh
  gin --config .ginconfig run main.go
  ```
- **Windows:**  
  Sử dụng `run.bat`.

- **Docker Compose:**
  ```sh
  docker-compose up --build
  ```

---

## Các chức năng nổi bật

### 1. Dựng API Restful (Gin Framework)
- **GET**: Lấy dữ liệu dạng list/object
- **POST**: Gửi/lưu trữ dữ liệu, tạo token, upload 1 hoặc nhiều file
- **DELETE**: Xoá dữ liệu
- **PUT**: Cập nhật dữ liệu

### 2. Middleware
- Xây dựng custom middleware (auth, logging, recover, CORS...)

### 3. Websocket (Socket trong Golang)
- Tích hợp websocket với Gin, ví dụ gửi thông báo real-time.

### 4. Gin CLI
- Sử dụng Gin CLI để hỗ trợ auto-reload khi dev.

### 5. Dependency Injection
- Áp dụng DI bằng cách sử dụng interface trong Go (inject repository/service).

### 6. Load ENV file
- Sử dụng thư viện `github.com/joho/godotenv` để load `.env`.

### 7. Logging & Config
- Tuỳ chỉnh logger cho Gin (log file, log level, format...)
- Phân biệt cấu hình debug/dev và production bằng `.env`.

### 8. Database
- Dùng PostgreSQL (hoặc MySQL tuỳ chỉnh)
- Sử dụng GORM để thao tác ORM
  - Hỗ trợ cả ORM lẫn native SQL query
- Migration tự động bằng GORM AutoMigrate

### 9. Migration & Chuyển đổi dữ liệu
- Hỗ trợ migration giữa các phiên bản SQL (có thể dùng GORM hoặc tool migration như goose/migrate).
- Hướng dẫn dump/restore và chuyển đổi dữ liệu giữa các channel.

### 10. Build đa nền tảng (Linux, Windows)
- Build cho Linux:
  ```sh
  GOOS=linux GOARCH=amd64 go build -o build/intern-go-linux main.go
  ```
- Build cho Windows:
  ```sh
  GOOS=windows GOARCH=amd64 go build -o build/intern-go.exe main.go
  ```

### 11. Debug với VSCode và Delve
- Cấu hình launch.json để debug trực tiếp với Gin CLI và Delve.

### 12. Docker hóa
- Có Dockerfile, docker-compose.yml cho môi trường phát triển và production.
- Dễ dàng đẩy code + DB lên container và connect các service.

---

## Cấu hình môi trường & Debug

- **Config log trong Gin**: Tuỳ chỉnh format, output ra file/console.
- **Config debug/dev server và production server**: Dùng các biến trong `.env` để phân biệt.
- **Config debug/run bằng VSCode**: Sử dụng launch.json, ví dụ:
  ```json
  {
    "version": "0.2.0",
    "configurations": [
      {
        "name": "Launch Gin",
        "type": "go",
        "request": "launch",
        "mode": "auto",
        "program": "${workspaceFolder}/main.go",
        "envFile": "${workspaceFolder}/.env"
      }
    ]
  }
  ```

---

## Database & Migration

- **Tạo database cho project**: Sử dụng script SQL hoặc Docker Compose (Postgres).
- **Migration**: Sử dụng GORM AutoMigrate hoặc tool migration.
- **Chuyển data giữa các channel**: Hướng dẫn dump/restore bằng pg_dump, mysql_dump...

---

## Docker 

1. **Build Docker image**
   ```sh
   docker build -t intern-go .
   ```

2. **Chạy container**
   ```sh
   docker run --env-file .env -p 8080:8080 intern-go
   ```

3. **Docker Compose**
   - Sử dụng `docker-compose.yml` để khởi động cả app + database + các service khác nếu có.

---

## Liên hệ & Đóng góp

- Mọi đóng góp vui lòng tạo Pull Request hoặc Issue mới.
