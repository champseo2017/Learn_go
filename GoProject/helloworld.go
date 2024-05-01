package main

import "fmt"

/*
Q6: เขียนโปรแกรม Go เพื่ออธิบายการอัปเดตค่า value ใน map
*/

func main() {
	// สร้าง map
	scores := map[string]int {
		"Alice": 80,
		"Bob": 75,
	}
	fmt.Println("Original scores:", scores)
	// อัปเดตค่า score
	scores["Alice"] = 85
	scores["Bob"] = 92
	scores["Charlie"] = 77 // เพิ่ม key-value ใหม่
	fmt.Println("Updated scores:", scores)
}

/* 
เริ่มต้นด้วยการ import package "fmt"
ในฟังก์ชัน main() เราสร้าง map ชื่อ scores โดยใช้ map literal
แสดงค่าเริ่มต้นของ map ด้วย fmt.Println()
การอัปเดตค่า value ในmapทำได้โดยการกำหนดค่าใหม่ให้กับ key นั้นๆ เช่น scores["Alice"] = 85
ถ้าเป็น key ที่ยังไม่มีอยู่ใน map เช่น scores["Charlie"] = 77 จะเป็นการเพิ่ม key-value คู่ใหม่ลงใน map
หลังจากอัปเดตค่าเรียบร้อยแล้ว เราก็แสดงค่าของ map ด้วย fmt.Println() อีกครั้ง
*/