package main

import (
	"fmt"
)

/*
Tagless Switch เป็นประเภทของ switch ที่ไม่มีการตรวจสอบเงื่อนไขที่ switch statement แต่จะมีการตรวจสอบที่ case statement แต่ละอัน
*/

func Tagless_Switch_Case_Example(x string) {
	switch {
	case x == "Rahul":
		// ถ้า x เท่ากับ "Rahul" ให้แสดงข้อความ "This is Rahul"
		fmt.Println("This is Rahul")
	case x == "Arjun":
		// ถ้า x เท่ากับ "Arjun" ให้แสดงข้อความ "This is Arjun"
		fmt.Println("This is Arjun")
	case x == "Vikas":
		// ถ้า x เท่ากับ "Vikas" ให้แสดงข้อความ "This is Vikas"
		fmt.Println("This is Vikas")
	default:
		// ถ้าไม่ตรงกับเงื่อนไขใดเลย ให้แสดงข้อความ "He has no name"
		fmt.Println("He has no name")
	}
}

func main() {
	// เรียกใช้ฟังก์ชัน Tagless_Switch_Case_Example โดยส่งค่า "Arjun" เข้าไป
	Tagless_Switch_Case_Example("Arjun")
}
