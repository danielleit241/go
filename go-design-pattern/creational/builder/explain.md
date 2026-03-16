# Builder

## Định nghĩa

**Builder** là mẫu thiết kế cho phép **tạo object phức tạp từng bước**, tách phần khởi tạo ra khỏi struct đích. Client không cần biết chi tiết lắp ráp, chỉ cần dùng director hoặc builder để nhận object hoàn chỉnh.

Ý tưởng chính:

- **Builder**: interface (hoặc struct) với các bước build từng phần (tên, logger, notifier, …), cuối cùng trả về object hoàn chỉnh.
- **Director**: biết thứ tự gọi các bước build; client gọi director thay vì tự gọi từng bước.
- Kết quả: code khởi tạo gọn, dễ đổi cách “lắp ráp” mà không đụng vào struct đích.

---

## Vấn đề

Trong [problem/main.go](problem/main.go):

- Có struct `ComplexService` với nhiều field (name, logger, notifier, dataLayer, uploader).
- Khởi tạo trực tiếp buộc phải truyền đủ tham số hoặc gán từng field → code dài, dễ sai, khó bảo trì khi thêm/bớt thành phần.

---

## Giải quyết

Trong [solution/main.go](solution/main.go):

1. **Interface Builder**
   - Các bước: `reset()`, `buildName()`, `buildLogger()`, `buildNotifier()`, `buildDataLayer()`, `buildUploader()`, và `result() Service`.
   - `serviceBuilder` implement Builder, giữ một `complexService` bên trong và lấp đầy từng bước; `result()` trả về service hoàn chỉnh.

2. **Director**
   - `ServiceDirector` có method `BuildService(builder Builder) Service`.
   - `serviceBuilderDirector` gọi lần lượt các bước build trên builder rồi trả về `builder.result()`.

3. **Cách dùng**
   - Client tạo director và builder, gọi `director.BuildService(builder)` là nhận `Service` sẵn dùng, không cần biết chi tiết lắp ráp.

---

## Tóm tắt

Builder giúp tạo object phức tạp theo từng bước, dễ đọc và dễ thay đổi cách khởi tạo.

Trong ví dụ này:

- **Builder** (`serviceBuilder`) đảm nhiệm từng bước build và trả về `Service`.
- **Director** (`serviceBuilderDirector`) quyết định thứ tự gọi các bước.

Kết quả: client chỉ cần director + builder để có `ComplexService` hoàn chỉnh, code gọn và dễ mở rộng.
