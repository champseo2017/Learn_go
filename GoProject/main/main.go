package main

import "fmt"

/*
ในภาษา Go เราสามารถใช้เทคนิค Currying ได้ คือการเรียกใช้ฟังก์ชันที่คืนค่าเป็นฟังก์ชันอีกทีโดยตรง
*/

func Return_Func_From_Another_Func() func() {
	return func() {
		fmt.Println("hello world") // ฟังก์ชันที่คืนค่ากลับจะพิมพ์ข้อความ "hello world"
	}
}

func main() {
	Return_Func_From_Another_Func()() // เรียกใช้ฟังก์ชันที่คืนค่ากลับโดยตรง (Currying)
}

/*
1. เราประกาศฟังก์ชัน `Return_Func_From_Another_Func` ซึ่งมีการคืนค่าเป็นฟังก์ชัน (`func()`)
2. ภายในฟังก์ชัน `Return_Func_From_Another_Func` เรา return ฟังก์ชันแบบ anonymous (ไม่ระบุชื่อ) ออกมา
   - ฟังก์ชันที่คืนค่ากลับนี้จะพิมพ์ข้อความ "hello world" ออกมาทาง console
3. ในฟังก์ชัน `main` เราเรียกใช้ฟังก์ชัน `Return_Func_From_Another_Func` และเรียกใช้ฟังก์ชันที่คืนค่ากลับทันทีต่อท้าย (`()`)
   - นี่คือการใช้เทคนิค Currying โดยเราไม่ต้องเก็บฟังก์ชันที่คืนค่ากลับไว้ในตัวแปรก่อน แต่สามารถเรียกใช้งานโดยตรงได้เลย
*/
