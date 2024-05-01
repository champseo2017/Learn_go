package main

import "fmt"

/*
Q9: เขียนโปรแกรม Go เพื่ออธิบายการลบ key ออกจาก map
*/

func main() {
	// สร้าง map
	numbers := map[string]int {
		"one":   1,
        "two":   2,
        "three": 3,
        "four":  4,
	}
	fmt.Println("Original map:", numbers)

	// ลบ key ออกจาก map
	delete(numbers, "two")
	fmt.Println("Map after deletion:", numbers)
	// ลบ key ที่ไม่มีอยู่ใน map
	delete(numbers, "five")
	fmt.Println("Map after deletion again:", numbers)
}

/* 
ในฟังก์ชัน main() เราสร้าง map ชื่อ numbers โดยใช้ map literal
แสดงค่าเริ่มต้นของ map ด้วย fmt.Println()
การลบ key ออกจาก map ทำได้โดยใช้ฟังก์ชันbuilin
*/