package main

import "fmt"

/*
Q4: เขียนโปรแกรม Go เพื่ออธิบายการวนลูปผ่าน map
*/

func main() {
	// สร้าง map
	scores := map[string]int {
		"Alice": 85,
		"Bob": 92,
		"Charlie": 77,
	}
	// วนลูปผ่าน map
	for name, score := range scores {
		fmt.Printf("%s got %d points\n", name, score)
	}
}

/* 
เริ่มต้นด้วยการ import package "fmt"
ในฟังก์ชัน main() เราสร้าง map ชื่อ scores โดยใช้ map literal
เราสามารถวนลูปผ่าน map ได้โดยใช้ for range loop ในรูปแบบ for key, value := range mapname
ในแต่ละรอบของลูป ตัวแปร name จะได้รับค่า key และ score จะได้รับค่า value ของแต่ละคู่ใน map
เราสามารถนำ key และ value มาใช้ได้ในบล็อกของลูป เช่น การแสดงผลด้วย fmt.Printf()
*/