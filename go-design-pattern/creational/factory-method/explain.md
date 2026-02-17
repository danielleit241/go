# Factory Method

## Định nghĩa

**Factory Method (Function Constructor Style)** là cách áp dụng Factory Pattern bằng cách:

- Sử dụng **function type**
- Dùng **factory function** để tạo ra implementation phù hợp
- Tách logic tạo handler / object khỏi code client

---

## Hiểu đơn giản

Factory constructor function giúp bạn:

- Không cần dùng `new` hoặc struct cụ thể trong client
- Gom logic tạo object / handler vào 1 function duy nhất
- Dễ thêm loại implementation mới
- Giữ code **clean – ít boilerplate – dễ maintain**

---

## Ý tưởng cốt lõi

Thay vì:

```
                +----------------+
Input Type ---> |   if / switch  |
                +----------------+
                 /      |       \
                /       |        \
         New Email   New SMS   New Push
```

Ta sẽ:

```
                    +-------------------+
Input Type -------> | Factory Function  |
                    |  (NewNotifier)    |
                    +-------------------+
                               |
                               v
                       +----------------+
                       | NotifierFunc   | (Function Type)
                       +----------------+
                         /      |      \
                        /       |       \
                       v        v        v
                Email Func   SMS Func   Push Func
```
