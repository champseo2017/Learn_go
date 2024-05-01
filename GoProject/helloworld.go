package main

import "fmt"

/*
อธิบายการสร้าง map โดยใช้ฟังก์ชัน make() ในภาษา Go
*/

func main() {
	// สร้าง map โดยใช้ make()
	ages := make(map[string]int)

	// เพิ่มค่าลงใน map
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35

	fmt.Println("Ages:", ages)
}

/* 
เริ่มต้นโดยการ import package "fmt"
ในฟังก์ชัน main() เราสร้าง map ชื่อ ages โดยใช้ฟังก์ชัน make(map[key_type]value_type)
ที่นี่เราระบุให้ key เป็นชนิด string และ value เป็นชนิด int
หลังจากนั้นเราก็สามารถเพิ่มค่าลงใน map ได้โดยการกำหนด ages[key] = value
สุดท้ายแสดงค่าของ map ด้วย fmt.Println()
*/