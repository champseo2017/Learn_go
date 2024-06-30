package main

import (
	oops "GoProject/package"
	"fmt"
)

/*
ในภาษา Go เราสามารถใช้ Type Assertion เพื่อตรวจสอบว่า Interface ที่เรากำลังทำงานอยู่นั้น มีค่าเป็น Type ที่เราต้องการหรือไม่
*/

func Find_The_Type(s oops.Shape) {
	// ตรวจสอบว่า s เป็น Rectangle หรือไม่
	// ถ้าใช่ ให้กำหนดค่าให้กับตัวแปร rect และตัวแปร ok จะมีค่าเป็น true
	// ถ้าไม่ใช่ ตัวแปร ok จะมีค่าเป็น false
	rect, ok := s.(oops.Rectangle)
	if ok {
		fmt.Println("นี่คือ Rectangle", rect)
	}

	// ตรวจสอบว่า s เป็น Square หรือไม่
	// ถ้าใช่ ให้กำหนดค่าให้กับตัวแปร square และตัวแปร ok จะมีค่าเป็น true
	// ถ้าไม่ใช่ ตัวแปร ok จะมีค่าเป็น false
	square, ok := s.(oops.Square)
	if ok {
		fmt.Println("นี่คือ Square", square)
	}
}

func main() {
	// สร้างตัวแปร s เป็น Interface ชื่อ Shape โดยมีค่าเป็น Square ที่มี Length เท่ากับ 10
	var s oops.Shape = oops.Square{Side: 10}

	// เรียกใช้ฟังก์ชัน Find_The_Type เพื่อตรวจสอบว่า s เป็น Type ใด
	Find_The_Type(s)
}

/*
ถ้า s เป็น Rectangle โปรแกรมจะแสดงข้อความ "นี่คือ Rectangle" ตามด้วยค่าของตัวแปร rect
แต่ถ้า s เป็น Square โปรแกรมจะแสดงข้อความ "นี่คือ Square" ตามด้วยค่าของตัวแปร square แทน
*/
