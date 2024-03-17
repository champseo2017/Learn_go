package main

import "fmt"

// Arithmetic

func main() {
	var a int = 10
	var b int = 3

	fmt.Println("a + b = ", a+b) // บวก
	fmt.Println("a - b = ", a-b) // ลบ
	fmt.Println("a * b = ", a*b) // คูณ
	fmt.Println("a / b = ", a/b) // หาร
	fmt.Println("a % b =", a%b)  // หารเอาเศษ

	a++ // เพิ่มค่า a ขึ้นหนึ่ง
	fmt.Println("a++ = ", a)

	b-- // ลดค่า b ลงหนึ่ง
	fmt.Println("b-- = ", b)
}
