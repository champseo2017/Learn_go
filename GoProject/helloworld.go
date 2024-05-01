package main

import "fmt"

/*
เขียนโปรแกรม Go เพื่อสร้าง map โดยใช้ map literal
*/

func main() {
	// สร้าง map ด้วย map literal
	capitals := map[string]string {
		"Thailand": "Bangkok",
        "Japan":    "Tokyo",
        "USA":      "Washington D.C.",
        "Germany":  "Berlin",
	}
	fmt.Println("Capitals:", capitals)
}

/* 
เริ่มต้นด้วยการ import package "fmt" เพื่อใช้ฟังก์ชันพื้นฐานในการแสดงผลข้อความ
ในฟังก์ชัน main() เราสร้าง map ชื่อ capitals โดยกำหนดให้ key เป็นชนิด string และ value ก็เป็น string เช่นกัน
เราสามารถกำหนดค่าเริ่มต้นของ map ด้วย map literal คือการระบุคู่ key-value ไว้ในวงเล็บ {} ในรูปแบบ key:value
หลังจากนั้นเราก็แสดงค่าของ map ด้วย fmt.Println()
*/