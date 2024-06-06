package main

import "fmt"

/*
ในภาษา Go เราสามารถคืนค่าฟังก์ชันจากฟังก์ชันอื่นได้
*/

func Return_Func_From_Another_Func() func() {
	return func() {
		fmt.Println("hello world") // ฟังก์ชันที่คืนค่ากลับจะพิมพ์ข้อความ "hello world"
	}
}

func main() {
	f := Return_Func_From_Another_Func() // กำหนดฟังก์ชันที่คืนค่ากลับให้กับตัวแปร f
	f()                                  // เรียกใช้ฟังก์ชันที่คืนค่ากลับผ่านตัวแปร f
}

/*
1. เราประกาศฟังก์ชัน `Return_Func_From_Another_Func` ซึ่งมีการคืนค่าเป็นฟังก์ชัน (`func()`)
2. ภายในฟังก์ชัน `Return_Func_From_Another_Func` เรา return ฟังก์ชันแบบ anonymous (ไม่ระบุชื่อ) ออกมา
   - ฟังก์ชันที่คืนค่ากลับนี้จะพิมพ์ข้อความ "hello world" ออกมาทาง console
3. ในฟังก์ชัน `main` เรากำหนดให้ตัวแปร `f` มีค่าเท่ากับฟังก์ชันที่คืนค่ากลับจาก `Return_Func_From_Another_Func`
4. เรียกใช้ฟังก์ชันที่คืนค่ากลับผ่านตัวแปร `f` โดยการเขียน `f()`
*/
