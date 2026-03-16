# Strategy Design Pattern

## Định nghĩa

**Strategy Design Pattern** là mẫu thiết kế cho phép định nghĩa một tập hợp các thuật toán, đóng gói từng thuật toán thành các thành phần riêng biệt và cho phép chúng thay thế cho nhau. Thuật toán có thể thay đổi độc lập với phần code sử dụng nó (client).

---

## Vấn đề

Trong [go-design-pattern/behavioral/strategy/problem/main.go](go-design-pattern/behavioral/strategy/problem/main.go):

- Struct `Notification` dùng `switch` theo `NotificationType` trong method `Send` để chọn cách gửi (email, sms).
- Khi thêm loại thông báo mới (ví dụ Telegram), phải sửa trực tiếp method `Send` — vi phạm nguyên tắc Open/Closed (mở để mở rộng, đóng với sửa đổi).
- Logic nhiều thuật toán bị dính chặt vào một struct, khó tái sử dụng và bảo trì.

---

## Giải quyết

Trong [go-design-pattern/behavioral/strategy/solution/main.go](go-design-pattern/behavioral/strategy/solution/main.go):

- Định nghĩa interface `Notifier` với method `Send(message string)`.
- Mỗi cách gửi là một strategy riêng: `EmailNotifier`, `SMSNotifier` (và dễ thêm `TelegramNotifier` sau này).
- `NotificationService` giữ `notifier Notifier` và gọi `notifier.Send(message)`; client có thể đổi strategy (email, sms) mà không sửa code service.

Thêm loại thông báo mới chỉ cần implement `Notifier` và truyền vào service, không cần sửa code có sẵn.

---

## Tóm tắt

Strategy giúp tách các thuật toán (cách gửi thông báo) thành các thành phần thay thế được, giảm `if/switch` và tuân thủ Open/Closed. Client làm việc với interface, logic cụ thể nằm trong từng strategy; code dễ mở rộng và bảo trì hơn.
