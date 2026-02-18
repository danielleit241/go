# Composite

## 1) Định nghĩa

**Composite** là mẫu thiết kế cho phép biểu diễn cấu trúc cây (tree) giữa object đơn lẻ và object nhóm, đồng thời xử lý chúng theo cùng một interface.

Ý tưởng chính:

- Tạo một abstraction chung (ví dụ `Item` với hàm `Cost()`)
- **Leaf** đại diện cho phần tử đơn (ví dụ `RealItem`)
- **Composite** đại diện cho nhóm phần tử (ví dụ `Box` chứa nhiều `Item`)

Nhờ đó, client có thể gọi cùng một hành vi (`Cost`) cho cả item đơn và hộp chứa nhiều item con.

---

## 2) Vấn đề trong ví dụ `problem`

Trong [structural/composite/problem/main.go](structural/composite/problem/main.go):

- Dùng một struct `Item` cho cả vai trò item thật và box
- Box phải đặt `Price = 0` để phân biệt với item thường
- Struct `Item` vừa mang dữ liệu sản phẩm, vừa mang `children`, nên trách nhiệm bị trộn lẫn

Cách này vẫn chạy được nhưng khó mở rộng, vì type không tách bạch rõ giữa **leaf** và **composite**.

---

## 3) Cách giải trong ví dụ `solution`

Trong [structural/composite/solution/main.go](structural/composite/solution/main.go):

- Định nghĩa interface `Item` với `Cost() float64`
- Tách `RealItem` (leaf): trả về trực tiếp `Price`
- Tách `Box` (composite): duyệt `children []Item` và cộng dồn `Cost()` của từng child

Như vậy `CreatePackage()` có thể trả về `Item`, và client không cần biết đó là `RealItem` hay `Box`, chỉ cần gọi `Cost()`.

---

## 4) Tóm tắt

Composite giúp bạn xử lý object đơn và object nhóm theo cùng một cách.

Trong ví dụ này:

- `RealItem` là phần tử đơn có giá riêng
- `Box` là phần tử nhóm, có thể chứa item hoặc box con

Kết quả là cấu trúc dữ liệu rõ ràng hơn, dễ mở rộng hơn, và rất phù hợp cho bài toán dạng cây/lồng nhau.
