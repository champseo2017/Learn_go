package main

import "fmt"

/*
Q5: การเพิ่มคู่ key-value ลงใน map ในภาษา Go
*/

func main() {
	// สร้าง map
	person := make(map[string]string)
	// เพิ่มคู่ key-value
	person["name"] = "Alice"
	person["age"] = "25"
	person["city"] = "Bangkok"
	fmt.Println("Detaile:", person)
}

/* 
เริ่มต้นด้วยการ import package "fmt"
ในฟังก์ชัน main() เราสร้าง map ชื่อ person ด้วย make(map[string]string) โดยกำหนดให้ทั้ง key และ value เป็นชนิด string
การเพิ่มคู่ key-value ลงใน map ทำได้โดยการกำหนดค่าในรูปแบบ mapname[key] = value
เช่น person["name"] = "Alice" หมายถึงการเพิ่มคู่ key "name" กับ value "Alice" ลงใน map
หลังจากเพิ่มข้อมูลแล้ว เราก็แสดงค่าของ map ด้วย fmt.Println()
*/