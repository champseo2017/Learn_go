package main

import (
	"fmt"
)

/*
สามารถเข้าถึงค่า field ของ struct ผ่าน pointer ได้โดยใช้หรือไม่ใช้เครื่องหมาย asterisk และเมื่อเราเปลี่ยนแปลงค่า field ของ struct ผ่าน pointer การเปลี่ยนแปลงนั้นจะสะท้อนกลับไปยัง struct ต้นฉบับด้วย
*/

// กำหนดโครงสร้างข้อมูล person ที่ประกอบด้วย field id และ name
type person struct {
	id   int
	name string
}

func main() {
	p1 := person{101, "XYX"}
	// สร้าง struct p1 ด้วยค่าเริ่มต้น id เป็น 101 และ name เป็น "XYX"

	p2 := &p1
	// สร้าง pointer p2 ที่ชี้ไปยัง struct p1

	fmt.Println("Person *p2.id:", (*p2).id)
	// แสดงค่า id ของ struct ที่ pointer p2 ชี้ไป โดยใช้ *p2 เพื่อ dereference

	fmt.Println("Person p2.id:", p2.id)
	// แสดงค่า id ของ struct ที่ pointer p2 ชี้ไป โดยไม่ต้องใช้ *p2

	p2.name = "ABC"
	// เปลี่ยนค่า name ของ struct ที่ pointer p2 ชี้ไปเป็น "ABC"

	fmt.Println("Person p1:", p1)
	// แสดงค่าของ struct p1
}

/*

 */
