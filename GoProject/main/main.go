package main

import "fmt"

/*
Value Receiver และ Pointer Receiver ในภาษา Go:

1. Value Receiver คือการส่งค่าของ Receiver (ตัวแปรที่เป็น Type ของ Method) เข้าไปใน Method
   - การเปลี่ยนแปลงค่าใน Method ที่ใช้ Value Receiver จะไม่ส่งผลต่อค่าเดิมของ Receiver ในฟังก์ชันที่เรียกใช้
   - Value Receiver เหมาะสำหรับการอ่านค่าจาก Receiver เท่านั้น

2. Pointer Receiver คือการส่งพอยน์เตอร์ (Memory Address) ของ Receiver เข้าไปใน Method
   - การเปลี่ยนแปลงค่าใน Method ที่ใช้ Pointer Receiver จะส่งผลต่อค่าเดิมของ Receiver ในฟังก์ชันที่เรียกใช้
   - Pointer Receiver เหมาะสำหรับการแก้ไขค่าของ Receiver หรือเมื่อต้องการให้การเปลี่ยนแปลงใน Method ส่งผลต่อค่าเดิมของ Receiver

การเลือกใช้ Value Receiver หรือ Pointer Receiver ขึ้นอยู่กับวัตถุประสงค์ของ Method ว่าต้องการแค่อ่านค่าจาก Receiver หรือต้องการแก้ไขค่าของ Receiver ด้วย ซึ่งจะช่วยให้เขียนโค้ดได้อย่างมีประสิทธิภาพและถูกต้องตามวัตถุประสงค์
*/

type rectangle struct {
	length float32
	width float32
}

func (r rectangle) area() float32 {
	return r.length * r.width
}

func (rect rectangle) perimeter() float32 {
	return 2 * (rect.length + rect.width)
}

func (rectangle) doSomething() {
	fmt.Println("I don't know what to do")
}

/* 
ประกาศ Method doSomething() ที่มีเฉพาะ Receiver Type เป็น rectangle โดยไม่มีชื่อ Receiver
Method นี้ไม่ได้อ่านหรือแก้ไขฟิลด์ใดๆ ของ rectangle struct แต่เพียงแค่แสดงข้อความทางหน้าจอ
*/

func main() {
	rect := rectangle{12.3, 21.45}
	a := rect.area()
	fmt.Println(a)
	/* 
	สร้างตัวแปร rect เป็น rectangle struct ด้วยค่า length เป็น 12.3 และ width เป็น 21.45
	เรียกใช้ Method area() ด้วย rect.area() และเก็บผลลัพธ์ไว้ในตัวแปร a
	แสดงค่าของ a ทางหน้าจอ
	*/
}
/* 
สรุปได้ว่า การประกาศ Method ใน Go สามารถทำได้โดยระบุ Receiver Type และชื่อ Receiver (ถ้าจำเป็น) การเรียกใช้ Method ทำได้โดยใช้เครื่องหมายจุด (.) กับค่าของ Type ที่ Method นั้นเชื่อมโยงอยู่ และ Method สามารถอ่านหรือแก้ไขฟิลด์ของ Receiver Type ได้ โดย Method และ Type ของมันต้องอยู่ในแพ็คเกจเดียวกัน
*/