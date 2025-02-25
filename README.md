### **Bài Test Tuyển Dụng Intern Developer Golang + ReactJS**

#### 🎯 **Mục tiêu:**
Đánh giá kỹ năng lập trình backend (Golang), frontend (ReactJS/HTML+CSS+JS), và khả năng triển khai (deployment) trên cloud của ứng viên.

---

### 🔥 **Yêu cầu Bài Test**

#### ✅ **Yêu cầu 1: Lập trình Backend bằng Golang (To-do List API)**
Xây dựng một REST API quản lý danh sách công việc (To-do List) với các tính năng sau:

1. **Thêm một công việc mới**
    - API: `POST /tasks`
    - Request body:
      ```json
      {
        "title": "Task title",
        "description": "Task description"
      }
      ```  
    - Response: Trả về ID của công việc vừa tạo.

2. **Lấy danh sách công việc**
    - API: `GET /tasks`
    - Response:
      ```json
      [
        {
          "id": 1,
          "title": "Task title",
          "description": "Task description",
          "completed": false
        }
      ]
      ```  

3. **Cập nhật trạng thái hoàn thành công việc**
    - API: `PUT /tasks/{id}`
    - Request body:
      ```json
      {
        "completed": true
      }
      ```  

4. **Xóa một công việc**
    - API: `DELETE /tasks/{id}`

**Yêu cầu kỹ thuật:**
- Sử dụng framework `Gin` hoặc `Echo`.
- Lưu trữ dữ liệu trong bộ nhớ (in-memory, sử dụng `map` hoặc `slice`).
- Xử lý lỗi với mã lỗi HTTP tương ứng.
- Code phải được tổ chức rõ ràng theo chuẩn Golang.

---

#### ✅ **Yêu cầu 2: Lập trình Frontend (ReactJS hoặc HTML+CSS+JS)**
Xây dựng giao diện người dùng (UI) để tương tác với API từ **Yêu cầu 1**:

1. Hiển thị danh sách các công việc từ API (`GET /tasks`).
2. Giao diện thêm công việc mới (sử dụng API `POST /tasks`).
3. Chức năng đánh dấu hoàn thành và xóa công việc (`PUT /tasks/{id}` và `DELETE /tasks/{id}`).
4. Giao diện đẹp, dễ sử dụng, tương thích tốt trên thiết bị di động.

**Yêu cầu kỹ thuật:**
- Sử dụng ReactJS (hoặc HTML+CSS+JS thuần nếu không quen với React).
- Sử dụng `axios` hoặc `fetch` để gọi API.
- Code cần tổ chức rõ ràng, dễ hiểu.

---

#### ✅ **Yêu cầu 3: Triển khai Backend và Frontend trên Cloud**
1. Tìm một nền tảng cloud có **free tier** (như **Render**, **Railway**, **Vercel**, **Netlify**, hoặc **AWS Free Tier**).
2. Deploy:
    - **Backend (Golang)**: Triển khai API trên nền tảng cloud.
    - **Frontend**: Deploy giao diện người dùng, có thể là trên Vercel hoặc Netlify (nếu sử dụng ReactJS).
3. Cấu hình Cross-Origin Resource Sharing (**CORS**) để backend và frontend có thể giao tiếp với nhau.

**Yêu cầu kỹ thuật:**
- Tạo file `README.md` hướng dẫn cách triển khai và sử dụng project.
- Nếu có thể, sử dụng Docker để đóng gói ứng dụng backend.

---

### 📋 **Yêu cầu nộp bài**

1. **Hoàn thành Yêu cầu 1** ➡️ Gửi link **GitHub repository** chứa mã nguồn backend.
2. **Hoàn thành Yêu cầu 1 & 2** ➡️ Gửi link **GitHub repository** chứa mã nguồn backend và frontend.
3. **Hoàn thành Yêu cầu 1 & 3** ➡️ Gửi link **GitHub repository** backend + file **Postman collection (.json)** chứa các API có thể kiểm thử (request/response).
4. **Hoàn thành cả 3 yêu cầu** ➡️ Gửi:
    - Link GitHub repository đầy đủ backend và frontend.
    - Link deploy của backend và frontend hoạt động trên nền tảng cloud.
    - File Postman collection chứa đầy đủ các API.

---

### 📅 **Thời hạn hoàn thành:**
**7 ngày** kể từ ngày nhận bài test.

---

### 🔑 **Tiêu chí đánh giá**
- ✅ Code sạch, dễ đọc.
- ✅ Xử lý lỗi tốt.
- ✅ Khả năng triển khai và cấu hình trên nền tảng cloud.

--- 