package main

import "fmt"

/*
- Golang ไม่มีไวยากรณ์ Class แบบภาษา OOP ทั่วไป แต่ใช้ Struct และ Receiver Function แทน
- Struct เปรียบเสมือน Class ใช้กำหนดโครงสร้างข้อมูล
- Receiver Function คือการประกาศฟังก์ชันให้ทำงานกับ Struct เฉพาะ เหมือน Method ใน Class
*/

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Printf("สวัสดี ฉันชื่อ %s อายุ %d ปี\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "ก้อง", Age: 30}
	p.SayHello()
}

/*
- ประกาศ Struct `Person` เก็บข้อมูล `Name` และ `Age`
- สร้างฟังก์ชัน `SayHello` ใช้ `(p Person)` เป็น Receiver เพื่อเข้าถึงข้อมูลใน `Person`
- สร้าง Object `p` จาก Struct `Person` และเรียกใช้ฟังก์ชัน `SayHello()`

สรุป Golang จำลองแนวคิด OOP ผ่าน Struct และ Receiver Function
*/
