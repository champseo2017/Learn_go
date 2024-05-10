package main

import (
	"fmt"
)

/*
ในตัวอย่างโค้ด จะแสดงการใช้โครงสร้างควบคุมแบบ if-else โดยมีการส่งค่า i เข้าไปในฟังก์ชัน Control_flow_example และทำการตรวจสอบเงื่อนไขของ i ถ้า i น้อยกว่า 100 จะแสดงข้อความว่า "conditionally executes a section of code in the If block" แต่ถ้า i มากกว่าหรือเท่ากับ 100 จะแสดงข้อความว่า "This code executes the else block"
*/

func Control_flow_example(i int) {
	// ตรวจสอบเงื่อนไขของ i
	if i < 100 {
		// ถ้า i น้อยกว่า 100 จะทำงานในบล็อก if
		fmt.Println("conditionally executes a section of code in the If block")
	} else {
		// ถ้า i มากกว่าหรือเท่ากับ 100 จะทำงานในบล็อก else
		fmt.Println("This code executes the else block")
	}
}

func main() {
	Control_flow_example(5)
}

/*

 */
