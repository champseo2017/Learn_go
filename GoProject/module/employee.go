package module

// โครงสร้าง Employee ที่ส่งออก
type Employee struct {
	Id int // ฟิลด์ที่ส่งออก
	employee_name string // ฟิลด์ที่ไม่ได้ส่งออก (สามารถเข้าถึงได้เฉพาะภายในแพ็คเกจ module)
	City string // ฟิลด์ที่ส่งออก
}