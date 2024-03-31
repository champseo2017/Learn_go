package main

import "fmt"

/*
เมื่อมีการเรียกใช้ฟังก์ชัน คุณสามารถส่งค่าให้กับพารามิเตอร์ของฟังก์ชันนั้นได้ พารามิเตอร์เปรียบเสมือนตัวยึดตำแหน่ง (placeholder) ค่าที่ส่งให้พารามิเตอร์เรียกว่า actual parameter หรือ argument ลิสต์ของพารามิเตอร์ (parameter list) ระบุชนิดข้อมูล (type), ลำดับ (order), และจำนวน (number) ของพารามิเตอร์ในฟังก์ชัน การใช้พารามิเตอร์ในฟังก์ชันไม่จำเป็นต้องมีเสมอไป ดังนั้นฟังก์ชันอาจไม่มีพารามิเตอร์เลยก็ได้
*/

// ฟังก์ชันที่ไม่มีพารามิเตอร์
func sayHello() {
	fmt.Println("Hello!")
}
// ฟังก์ชันที่มีพารามิเตอร์เดียว
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// ฟังก์ชันที่มีหลายพารามิเตอร์
func add(x, y int) int {
	return x + y
}

func main() {
	sayHello()
	greet("John")
	result := add(3, 5)
	fmt.Println(result)
}
