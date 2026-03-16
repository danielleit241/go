# Chain of Responsibility

## Định nghĩa

**Chain of Responsibility (CoR)** là mẫu thiết kế cho phép đưa một request đi qua một chuỗi handler. Mỗi handler xử lý một phần việc rồi chuyển tiếp cho handler kế tiếp.

Ý tưởng chính:

- Tách quy trình lớn thành nhiều bước nhỏ độc lập
- Mỗi bước là một handler riêng
- Dễ thay đổi thứ tự, thêm hoặc bớt bước mà không sửa toàn bộ luồng

---

## Vấn đề

Trong [go-design-pattern/behavioral/chain-of-responsibility/problem/main.go](go-design-pattern/behavioral/chain-of-responsibility/problem/main.go):

- Toàn bộ các bước crawl được viết cứng trong một hàm `Crawl`
- Logic các bước bị dính chặt với nhau
- Khó tái sử dụng từng bước và khó mở rộng khi thêm bước mới

Cách này phù hợp demo đơn giản, nhưng sẽ khó maintain khi pipeline xử lý phức tạp hơn.

---

## Giải quyết

Trong [go-design-pattern/behavioral/chain-of-responsibility/solution/main.go](go-design-pattern/behavioral/chain-of-responsibility/solution/main.go):

- Mỗi bước được tách thành một `Handler` (`CheckingURL`, `FetchingContent`, `ExtractingLinks`, `SavingToDatabase`)
- Dữ liệu dùng chung đi qua các bước thông qua `Context`
- `HandlerNode` tạo cấu trúc linked-list để nối chuỗi xử lý
- `NewCrawler(...)` lắp chain theo thứ tự mong muốn

`WebCrawler.Crawl(url)` chỉ cần gọi chain một lần, từng node tự xử lý và chuyển tiếp. Nếu có lỗi, chain dừng và trả lỗi lên để log.

---

## Tóm tắt

CoR giúp biến một luồng xử lý dài thành pipeline gồm nhiều bước độc lập.

Trong ví dụ này:

- Mỗi handler làm đúng một nhiệm vụ
- Thứ tự xử lý được cấu hình khi dựng chain
- Dễ thêm/bớt bước crawl mà không đụng nhiều vào code cũ

Kết quả là code rõ trách nhiệm hơn, dễ mở rộng và dễ maintain hơn.
