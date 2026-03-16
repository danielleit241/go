# Factory Method

## Định nghĩa

**Factory Method** là mẫu thiết kế tạo đối tượng, trong đó việc khởi tạo object được đưa vào một hàm factory thay vì để client khởi tạo trực tiếp.

Ý tưởng chính:

- Client chỉ làm việc với interface chung (ví dụ `Notifier`)
- Logic chọn implementation (`EmailNotifier`, `SMSNotifier`, ...) nằm trong factory method
- Việc tạo object được gom về một nơi duy nhất

---

## Vấn đề

Trong [go-design-pattern/creational/factory-method/problem/main.go](go-design-pattern/creational/factory-method/problem/main.go):

- `NotificationService` phụ thuộc trực tiếp vào concrete type (`&EmailNotifier{}`)
- User/client phải biết implementation cụ thể để khởi tạo
- Khi thêm loại notifier mới, code khởi tạo dễ bị sửa ở nhiều chỗ

Điều này làm tăng coupling giữa client và concrete implementation.

---

## Giải quyết

Trong [go-design-pattern/creational/factory-method/solution/main.go](go-design-pattern/creational/factory-method/solution/main.go):

- Thêm factory method `CreateNotifier(notificationType string) Notifier`
- Factory dùng `switch` để chọn và trả về notifier phù hợp
- `NotificationService` chỉ nhận `Notifier`, không cần biết concrete type

Ví dụ sử dụng:

```go
s := NotificationService{
	notifier: CreateNotifier("email"),
}
```

Như vậy client gọi factory để lấy `Notifier` thay vì tự `new` concrete struct.

---

## Tóm tắt

Factory Method giúp bạn viết code theo hướng mở rộng tốt hơn:

- Client biết **dùng gì** (interface `Notifier`)
- Factory quyết định **tạo gì** (`EmailNotifier`, `SMSNotifier`, ...)

Trong ví dụ này, pattern giúp ẩn chi tiết khởi tạo, giảm coupling và dễ thêm loại notifier mới.
