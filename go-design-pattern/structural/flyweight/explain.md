# Flyweight

## 1) Định nghĩa

**Flyweight** là mẫu thiết kế giúp giảm sử dụng bộ nhớ bằng cách **chia sẻ trạng thái dùng chung** giữa nhiều object thay vì lưu trùng lặp.

Ý tưởng chính:

- Tách dữ liệu thành 2 phần:
  - **Intrinsic state**: dữ liệu cố định, có thể chia sẻ (ví dụ avatar, profile sender)
  - **Extrinsic state**: dữ liệu theo ngữ cảnh từng object (ví dụ nội dung message)
- Dùng một factory/cache để tái sử dụng phần intrinsic đã tạo trước đó.

---

## 2) Vấn đề trong ví dụ `problem`

Trong [structural/flyweight/problem/main.go](structural/flyweight/problem/main.go):

- Mỗi `ChatMessage` lưu trực tiếp `SenderName` và `SenderAvatar []byte`
- Avatar là dữ liệu nặng (hàng trăm KB)
- Nhiều message cùng sender sẽ bị **nhân bản avatar** nhiều lần trong RAM

Kết quả: bộ nhớ tăng nhanh theo số lượng message, dù số sender thực tế ít.

---

## 3) Cách giải trong ví dụ `solution`

Trong [structural/flyweight/solution/main.go](structural/flyweight/solution/main.go):

- `ChatMessage` chỉ giữ `Content` + con trỏ `Sender`
- `Sender` chứa dữ liệu nặng (`Avatar`)
- `SenderFactory` giữ `cacheSender map[string]*Sender`
- `GetSender(name)`:
  - Nếu đã có trong cache thì trả lại object cũ
  - Nếu chưa có thì tạo mới, lưu cache, rồi trả về

Như vậy nhiều message của cùng một người gửi sẽ dùng chung một `*Sender`.

---

## 4) Tóm tắt

Flyweight tối ưu bộ nhớ bằng cách **tách phần dùng chung và chia sẻ nó qua cache/factory**.

Trong ví dụ này:

- Message giữ dữ liệu thay đổi theo từng dòng chat
- Sender giữ dữ liệu nặng và được tái sử dụng

Kết quả là code tiết kiệm RAM hơn và dễ mở rộng khi số lượng message tăng lớn.
