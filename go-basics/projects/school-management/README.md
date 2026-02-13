# School Management (Go Console)

## 1) Mục tiêu project

Project này là ứng dụng quản lý trường học chạy trên console, gồm 2 nhóm dữ liệu:

- Student
- Lecturer

Ứng dụng được tổ chức theo kiến trúc nhiều lớp để dễ bảo trì:

- UI layer: nhận input, hiển thị menu/kết quả
- Service layer: business rule + validate nghiệp vụ
- Repository layer: lưu trữ và truy vấn dữ liệu
- Entity layer: mô hình dữ liệu domain

Luồng gọi chính:
UI -> Service -> Repository -> in-memory slice

## 2) Cách chạy

Yêu cầu: Go 1.25+

Chạy từ thư mục gốc module `projects/school-management`:

```bash
go run ./cmd/app
```

Build:

```bash
go build ./...
```

Lưu ý: nếu chạy lệnh từ thư mục cha module sẽ gặp lỗi "cannot find main module".

## 3) Giải thích kiến trúc và lý do dùng struct/interface

### Vì sao dùng struct?

- Struct chứa state và hành vi theo domain, ví dụ `Student`, `Lecturer`, `Menu`, `studentService`.
- Struct giúp gom dữ liệu liên quan trong một kiểu rõ ràng, dễ mở rộng.
- Với service/repository, struct giữ dependency đã được inject để dùng xuyên suốt vòng đời object.

### Vì sao dùng interface?

- Interface tách contract khỏi implementation cụ thể.
- Service phụ thuộc vào interface repository thay vì phụ thuộc concrete type, giúp thay thế implementation (memory, database, mock) mà không sửa service.
- Interface nhỏ theo use-case (`Reader`, `Writer`) giúp inject đúng nhu cầu, giảm coupling.

### Vì sao dùng constructor injection?

- Dependency được tạo ở `main` (composition root), rồi truyền vào service/UI.
- Service không tự `new` repository nên dễ unit test và dễ thay đổi hạ tầng.

## 4) Chi tiết từng thư mục và file

### cmd/app

#### cmd/app/main.go

Vai trò:

- Điểm vào chương trình.
- Khởi tạo repository, service, menu và chạy app.

Lý do đặt ở đây:

- Đây là nơi phù hợp để làm composition root cho DI.
- Tách rõ phần "wiring" khỏi business logic.

### internal/entity

#### internal/entity/person.go

Vai trò:

- Định nghĩa thông tin chung của một người: `ID`, `FirstName`, `LastName`, `Age`.

Lý do dùng struct:

- Là kiểu dữ liệu nền, được nhúng (embedded) vào Student/Lecturer để tránh lặp field.

#### internal/entity/student.go

Vai trò:

- Domain model cho học sinh: nhúng `Person`, thêm `Grade`, `GPA`.
- Implement method `GetID()`.

Lý do có `GetID()`:

- Để thỏa interface `HasID` ở repository generic.
- Nhờ đó generic repository có thể thao tác bằng khóa ID cho mọi entity cùng contract.

#### internal/entity/lecturer.go

Vai trò:

- Domain model cho giảng viên: nhúng `Person`, thêm `Department`, `Salary`.
- Implement method `GetID()` để dùng chung generic repository.

### internal/repository

#### internal/repository/generic_repo.go

Vai trò:

- Chứa phần repository dùng chung bằng generic:
  - `HasID`
  - `BaseReader`, `BaseWriter`, `BaseRepository`
  - `baseRepository[T]` với CRUD cơ bản
  - Error dùng chung (`ErrNotFound`, `ErrDuplicateID`, `ErrInvalidInput`)

Điểm kỹ thuật quan trọng:

- Constraint `T HasID` đảm bảo compiler biết `T` có `GetID()`.
- `GetAll()` dùng `slices.Clone` để tránh lộ trực tiếp slice nội bộ.

Lý do dùng struct + interface tại đây:

- Interface định nghĩa contract thao tác dữ liệu.
- Struct generic tái sử dụng logic CRUD, giảm duplicate giữa Student/Lecturer.

#### internal/repository/student_repo.go

Vai trò:

- Contract riêng cho Student:
  - `StudentReader`
  - `StudentWriter`
  - `StudentRepository`
- Implementation `studentRepository` (in-memory) bằng cách embed `baseRepository[entity.Student]`.
- Bổ sung behavior riêng: `SearchByName`.

Lý do tách file riêng:

- Tránh file repository quá to.
- Dễ theo dõi phần nghiệp vụ đặc thù của Student.

#### internal/repository/lecturer_repo.go

Vai trò:

- Contract riêng cho Lecturer:
  - `LecturerReader`
  - `LecturerWriter`
  - `LecturerRepository`
- Implementation `lecturerRepository` và logic đặc thù `SearchByDepartment`.

Lý do tách file riêng:

- Tăng tính module hóa, dễ mở rộng khi có thêm entity mới.

#### internal/repository/repository.go

Vai trò hiện tại:

- File dùng để giữ phần khai báo chung của package repository khi cần mở rộng thêm.

Gợi ý:

- Có thể giữ file này tối giản hoặc dồn toàn bộ phần dùng chung vào `generic_repo.go` để tránh trùng vai trò.

### internal/service

#### internal/service/student_service.go

Vai trò:

- Định nghĩa `StudentService` interface.
- Implementation `studentService` nhận dependency qua DI.
- Validate input nghiệp vụ trước khi gọi repository.
- Constructor:
  - `NewStudentService(repo StudentRepository)` cho production
  - `NewStudentServiceWithPorts(reader, writer)` cho test/mocking

Lý do struct + interface:

- Interface là public contract của layer service.
- Struct giữ state dependency (`reader`, `writer`) để thực thi use-case.

#### internal/service/lecturer_service.go

Vai trò tương tự student service:

- Contract `LecturerService`.
- Validate data lecturer.
- Dùng constructor injection để tách dependency.

### internal/ui

#### internal/ui/menu.go

Vai trò:

- Điều hướng menu chính và menu con.
- Thực hiện các thao tác CRUD/Search qua service.
- Chuyển đổi dữ liệu nhập thành entity trước khi gọi service.

Lý do dùng struct `Menu`:

- Giữ dependency `studentService`, `lecturerService`.
- Tách rõ phần điều hướng UI khỏi `main`.

#### internal/ui/inputer.go

Vai trò:

- Chứa helper nhập liệu: đọc input, parse số, và validate cơ bản.
- Là điểm gom logic nhập từ console để tránh lặp code ở menu.

Ghi chú:

- Tên file hiện tại là `inputer.go` (không phải `input.go`).

### pkg/utils

#### pkg/utils/validator.go

Vai trò:

- Chứa các hàm validation dùng lại ở UI:
  - ValidateID
  - ValidateName
  - ValidateGPA
  - ValidateSalary
  - ValidateDepartment
  - ValidateAge

Lý do để ở `pkg`:

- Đây là utility độc lập, không phụ thuộc layer domain cụ thể.

## 5) Nguyên tắc thiết kế đang áp dụng

- Separation of concerns giữa UI / Service / Repository / Entity.
- Dependency Injection tại `main`.
- Interface nhỏ theo use-case (reader/writer).
- Generic repository để giảm duplicate CRUD.
- Validation nhiều lớp:
  - UI: validate format nhập
  - Service: validate nghiệp vụ

## 6) Hướng mở rộng

- Thay repository memory bằng database (PostgreSQL/MySQL) mà không đổi service contract.
- Thêm unit test cho service và repository với mock/fake.
- Tách thêm application layer nếu cần use-case phức tạp hơn.
