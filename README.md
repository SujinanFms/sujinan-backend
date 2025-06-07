# sujinan-backend

<!-- Todo: 1. บน Terminal -->
<!-- mkdir sujinan-backend
cd sujinan-backend
go mod init github.com/SujinanFms/sujinan-backend -->

<!-- //? Structure -->

sujinan-backend/
├── go.mod
├── main.go
├── Dockerfile
├── .env
├── handlers/
│ ├── contact.go
│ └── recommendations.go
├── models/
│ ├── contact.go
│ └── recommendation.go
├── database/
│ └── db.go
└── routes/
└── router.go

<!--  Todo: 1.สร้างไฟล์ main.go -->

<!--  Todo: 2.สร้างไฟล์ routes/router.go -->

ในภาษา Go จัดโค้ดเป็น package เพื่อแยกความรับผิดชอบ เช่น:

- main.go สำหรับเริ่มโปรแกรม
- routes สำหรับกำหนด path (/api/contact, /api/recommendations)
- handlers สำหรับจัดการ "พฤติกรรม" ของแต่ละ endpoint

สรุป
สิ่งที่ทำแล้ว รายละเอียด
✅ เชื่อม PostgreSQL ใช้ Docker Compose และเชื่อมผ่าน database/db.go
✅ เขียน handler API /api/contact และ /api/recommendations
✅ สร้าง model struct สำหรับ Contact และ Recommendation
✅ ทำ migration สร้างตารางใน Migrate()
✅ ทดสอบ API ผ่าน Postman หรือ Frontend
