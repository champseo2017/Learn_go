package main

import "fmt"

/*
เกี่ยวกับ Pointer Receiver และ Value Receiver ในภาษา Go:

1. สามารถเรียกใช้ Method ที่มี Pointer Receiver ด้วยตัวแปรที่ไม่ใช่ Pointer ได้ โดย Go Compiler จะทำการ Referencing (ดึงที่อยู่) ให้โดยอัตโนมัติ
2. สามารถเรียกใช้ Method ที่มี Value Receiver ด้วยตัวแปรที่เป็น Pointer ได้ โดย Go Compiler จะทำการ Dereferencing (ดึงค่า) ให้โดยอัตโนมัติ
3. การ Referencing และ Dereferencing จะทำงานได้เฉพาะเมื่อมีการใช้ตัวแปรเท่านั้น ไม่สามารถทำกับ Composite Literal (ค่าที่ประกาศโดยตรง) ได้
4. Type สามารถมีทั้ง Pointer Receiver และ Value Receiver Method ผสมกันได้ แต่แนะนำให้ใช้อย่างใดอย่างหนึ่งเป็นหลักเพื่อความชัดเจน
5. Pointer Receiver มีข้อดีกว่า Value Receiver ในแง่ของประสิทธิภาพและความยืดหยุ่นในการปรับเปลี่ยนค่าของ Receiver ในอนาคต

เมื่อเลือกใช้ Receiver ใน Method ให้พิจารณาถึงวัตถุประสงค์ของ Method และความยืดหยุ่นในการพัฒนาโค้ดต่อไป หากไม่แน่ใจ แนะนำให้ใช้ Pointer Receiver เป็นหลัก และหลีกเลี่ยงการเรียกใช้ Method ที่มี Pointer Receiver ด้วย Composite Literal โดยตรง ควรประกาศตัวแปรเก็บค่าก่อนเพื่อให้โค้ดอ่านง่ายและลดโอกาสเกิดข้อผิดพลาด

*/

type rectangle struct {
	length float64
	width float64
}

func (r *rectangle) increaseLength(value float64) {
	r.length += value
}

func main() {
	// ตัวอย่างการเรียกใช้ Method ที่มี Pointer Receiver ด้วยตัวแปรที่ไม่ใช่ Pointer
	rect := rectangle{10, 20}
	rect.increaseLength(5)
	fmt.Println(rect) // Output: {15 20}
	/* 
	ในตัวอย่างนี้ increaseLength เป็น Method ที่มี Pointer Receiver (*rectangle) แต่ในฟังก์ชัน main เราเรียกใช้ increaseLength ด้วยตัวแปร rect ซึ่งเป็น Value (ไม่ใช่ Pointer) โดยเขียนเป็น rect.increaseLength(5) ทั้งๆ ที่ rect ไม่ใช่ Pointer แต่ Go Compiler จะทำการ Referencing ให้โดยอัตโนมัติ กล่าวคือจะส่งที่อยู่ (Memory Address) ของ rect ให้กับ increaseLength ทำให้ increaseLength สามารถแก้ไขค่าใน rect ได้โดยตรง
	*/
}
/* 

*/