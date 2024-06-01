package main

import "fmt"

/*
การประกาศตัวแปรที่มีชนิดข้อมูลเป็นฟังก์ชันใน Go:

1. ในภาษา Go เราสามารถประกาศตัวแปรที่มีชนิดข้อมูลเป็นฟังก์ชันได้
2. ใช้รูปแบบ `var ชื่อตัวแปร func(ชนิดข้อมูลของ Argument)` เพื่อประกาศตัวแปรที่มีชนิดข้อมูลเป็นฟังก์ชัน
3. สามารถ Assign ฟังก์ชันให้กับตัวแปรที่มีชนิดข้อมูลเป็นฟังก์ชันได้
4. เรียกใช้ฟังก์ชันที่เก็บไว้ในตัวแปรโดยใช้ `ชื่อตัวแปร(Argument)`
*/

func DeclareVariableOfFunctionType(s string) {
	var f func(string) = func(s string) {
		fmt.Println(s)
	}
	f(s)
}

func main() {
	DeclareVariableOfFunctionType("Hello Udit")
}

/*
ในตัวอย่าง ประกาศตัวแปร `f` ที่มีชนิดข้อมูลเป็นฟังก์ชัน และ Assign ฟังก์ชันให้กับตัวแปร `f` จากนั้นเรียกใช้ฟังก์ชันผ่านตัวแปร `f` โดยส่ง Argument เข้าไป แสดงให้เห็นว่าเราสามารถประกาศและใช้งานตัวแปรที่มีชนิดข้อมูลเป็นฟังก์ชันได้ในภาษา Go
*/
