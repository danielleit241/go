# Strategy Design Pattern

## Định nghĩa

**Strategy Design Pattern** là một mẫu thiết kế phần mềm cho phép:

- Định nghĩa **một tập hợp các thuật toán**
- **Đóng gói** từng thuật toán thành các thành phần riêng biệt
- Cho phép các thuật toán **có thể thay thế cho nhau**
- Giúp thuật toán **thay đổi độc lập** với phần code sử dụng nó (client)

---

## Hiểu đơn giản

Strategy Pattern giúp bạn:

- Định nghĩa **nhiều cách giải quyết khác nhau** cho cùng một bài toán
- **Thay đổi thuật toán linh hoạt** khi chương trình đang chạy
- Tránh viết nhiều `if - else` hoặc `switch` phức tạp
- Giữ code **dễ mở rộng – dễ bảo trì – clean hơn**

---

## Ý tưởng cốt lõi

Thay vì:

```
                +----------------+
Input Type ---> |   if / switch  |
                +----------------+
                 /      |       \
                /       |        \
         Algorithm A  Algorithm B  Algorithm C
```

Ta sẽ:

```
                    +-----------+
Input Type -------> |  Context  |
                    +-----------+
                           |
                           v
                   +----------------+
                   |   Strategy     | (Interface)
                   +----------------+
                     /      |      \
                    /       |       \
                   v        v        v
            +---------+ +---------+ +---------+
            |StrategyA| |StrategyB| |StrategyC|
            +---------+ +---------+ +---------+
```
