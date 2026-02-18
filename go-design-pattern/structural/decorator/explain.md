# Decorator

## 1) Định nghĩa

**Decorator** là mẫu thiết kế cho phép thêm hành vi vào object một cách linh hoạt bằng cách bọc (wrap) object đó, thay vì sửa trực tiếp class/struct gốc.

Ý tưởng chính:

- Giữ interface chung (ví dụ `Notifier`)
- Có object gốc (ví dụ `EmailNotifier`)
- Thêm nhiều lớp bọc để mở rộng hành vi (ví dụ gửi thêm qua SMS, Telegram)

Nhờ vậy có thể kết hợp tính năng theo từng nhu cầu mà không phải tạo quá nhiều struct “ghép cứng”.

---

## 2) Vấn đề trong ví dụ `problem`

Trong [structural/decorator/problem/main.go](structural/decorator/problem/main.go):

- Khi muốn gửi qua nhiều kênh, phải tạo struct mới như `EmailSMSNotifier`
- Nếu thêm tổ hợp mới (Email + SMS + Telegram), lại phải tạo thêm struct khác
- Số lượng struct tăng nhanh theo số cách kết hợp

Cách này dễ dẫn đến trùng lặp code và khó mở rộng.

---

## 3) Cách giải trong ví dụ `solution`

Trong [structural/decorator/solution/main.go](structural/decorator/solution/main.go):

- Tạo `NotifierDecorator` cũng implement `Notifier`
- `NotifierDecorator` giữ `notifier` hiện tại và con trỏ `core` để chain nhiều lớp bọc
- Hàm `Decorate(...)` thêm notifier mới vào chuỗi
- Khi gọi `Send(...)`, decorator gửi bằng notifier hiện tại rồi gọi tiếp xuống `core`

Ví dụ:

```go
notifier := NewNotifierDecorator(&EmailNotifier{}).
	Decorate(&SMSNotifier{}).
	Decorate(&TelegramNotifier{})
```

Kết quả: một lệnh gửi sẽ đi qua nhiều kênh mà không cần tạo struct tổ hợp riêng.

---

## 4) Tóm tắt

Decorator giúp mở rộng hành vi theo kiểu “xếp lớp”, không sửa code gốc và không tạo nhiều class/struct tổ hợp.

Trong ví dụ này:

- `EmailNotifier` là hành vi nền
- `SMSNotifier`, `TelegramNotifier` được thêm qua decorator chain

Kết quả là hệ thống notification linh hoạt hơn, dễ thêm/bớt kênh gửi và giảm coupling.
