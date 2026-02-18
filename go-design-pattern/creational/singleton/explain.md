# Singleton

## 1) Định nghĩa

**Singleton** là mẫu thiết kế đảm bảo một kiểu dữ liệu chỉ có **một instance duy nhất** trong toàn bộ ứng dụng, và cung cấp điểm truy cập toàn cục tới instance đó.

Ý tưởng chính:

- Chỉ khởi tạo object một lần
- Mọi nơi dùng chung cùng một instance
- Thường kết hợp cơ chế thread-safe khi chạy concurrent

Trong Go, cách phổ biến là dùng `sync.Once` để đảm bảo khởi tạo đúng 1 lần.

---

## 2) Vấn đề trong ví dụ `problem`

Trong [creational/singleton/problem/main.go](creational/singleton/problem/main.go):

- Cấu hình được tạo ở `main` rồi truyền ngữ cảnh sử dụng thủ công
- Hàm `requestHandler` cần biết config để log, nhưng chữ ký hàm không nhận config
- Dữ liệu dùng chung (config) chưa có điểm truy cập thống nhất cho toàn app

Khi ứng dụng lớn hơn hoặc có nhiều goroutine, việc quản lý config sẽ khó và dễ rối.

---

## 3) Cách giải trong ví dụ `solution`

Trong [creational/singleton/solution/main.go](creational/singleton/solution/main.go):

- Tạo `application` làm singleton, giữ `cfg *config`
- Cung cấp `GetApplication()` để truy cập instance dùng chung
- Dùng `sync.Once` trong `GetConfig()` để chỉ load config một lần
- Mọi nơi (`main`, `requestHandler`) đều gọi `GetApplication().GetConfig()`

Kết quả: config được lazy-load, dùng chung toàn hệ thống, an toàn khi có nhiều request đồng thời.

---

## 4) Tóm tắt

Singleton phù hợp cho tài nguyên dùng chung toàn cục như config, logger, connection manager.

Trong ví dụ này:

- `application` là singleton holder
- `sync.Once` đảm bảo chỉ load config đúng 1 lần
- Các request dùng chung cùng một config instance

Kết quả là giảm trùng lặp khởi tạo, dễ truy cập và ổn định hơn trong môi trường concurrent.
