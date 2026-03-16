# Prototype

## Định nghĩa

**Prototype** là mẫu thiết kế cho phép **tạo bản sao đúng nghĩa** của một object (clone sâu), thay vì copy bằng `=` (copy nông, dễ dính bug).

---

## Vấn đề

Xem [problem/main.go](go-design-pattern/creational/prototype/problem/main.go):

- Có `Point` (X, Y) và `Node` (một Point + một slice các Point).
- Khi gán `n2 := n1`, Go chỉ copy **tham chiếu** của slice `Children`.
- Kết quả: sửa `n1.Children[0].X = 10` thì **`n2.Children[0].X` cũng đổi theo** vì hai struct đang dùng chung một mảng bên trong (shallow copy).
- Muốn bản sao độc lập thì phải tự duyệt từng phần tử, tạo slice mới, copy từng Point… → code rối, dễ quên, khó bảo trì (deep copy).

---

## Giải quyết

Xem [solution/main.go](go-design-pattern/creational/prototype/solution/main.go):

1. **Định nghĩa interface clone**
   - Interface `Prototype` có method `Clone() Prototype`.
   - Mọi kiểu muốn “clone được” thì implement method này.

2. **Clone từng kiểu**
   - **Point**: `Clone()` tạo `Point` mới, copy X, Y.
   - **Node**: `Clone()` tạo slice mới, clone từng `Point` trong `Children` và clone cả `Value` → bản sao **hoàn toàn tách biệt** với bản gốc.

Sau khi dùng `n2 := n1.Clone().(*Node)`, sửa `n1.Children[0].X` **không** ảnh hưởng `n2`.

---

## Tóm tắt

Prototype giúp bạn tạo bản sao object độc lập (deep copy), không dính reference chung.

Trong ví dụ này:

- `Point` và `Node` clone qua interface `Prototype`.
- `Color` clone qua `CloneWithRed(r)` cũng là một prototype.
