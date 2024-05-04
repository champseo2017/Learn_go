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

func (r rectangle) area() float64 {
	return r.length * r.width
}

func main() {
	// ตัวอย่างการเรียกใช้ Method ที่มี Value Receiver ด้วยตัวแปรที่เป็น Pointer
	rect := &rectangle{10, 20}
	fmt.Println(rect.area()) // Output: 200
	/* 
	นตัวอย่างนี้ area เป็น Method ที่มี Value Receiver (rectangle) แต่ในฟังก์ชัน main เราเรียกใช้ area ด้วยตัวแปร rect ซึ่งเป็น Pointer โดยเขียนเป็น rect.area() ทั้งๆ ที่ rect เป็น Pointer แต่ Go Compiler จะทำการ Dereferencing ให้โดยอัตโนมัติ กล่าวคือจะดึงค่า (Value) ที่ Pointer rect ชี้ไปให้กับ area ทำให้ area สามารถอ่านค่าจาก *rect ได้โดยตรง
	*/
}
/* 

*/