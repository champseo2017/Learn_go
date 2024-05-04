package main

import "fmt"

/*
ในการประกาศ Method สามารถละชื่อ Receiver ได้ โดยระบุเฉพาะ Receiver Type ก็เพียงพอ หากMethod ไม่ได้อ่านหรือแก้ไขฟิลด์ใดๆ ของ Receiver สามารถละชื่อ Receiver ได้ (ถึงแม้ว่าจะไม่ใช่กรณีที่ถูกต้อง เพราะถ้า Method ไม่ได้อ่านหรือแก้ไขฟิลด์ใดๆ ก็ไม่มีความจำเป็นในการประกาศ Method นั้น)
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