package main

import (
	"fmt"
)

/*
Go เราสามารถสร้าง slice ใหม่จาก slice ที่มีอยู่แล้วได้
*/

func Slice_from_Another_Slice() {
	// สร้าง slice ชื่อ z
	z := []int{10, 32, 42, 23, 12, 41, 5}

	// สร้าง slice ใหม่ x จาก slice z โดยใช้ส่วนของ z ตั้งแต่ index 0 ถึง 2
	x := z[0:3]

	// แก้ไขค่าใน slice x ที่ index 0
	x[0] = 30000

	// แสดงผล slice z
	// เนื่องจาก x เป็นการอ้างอิงไปยังข้อมูลใน z การแก้ไขใน x จึงมีผลกับ z ด้วย
	fmt.Println(z)
}

func main() {
	Slice_from_Another_Slice()
}

/*

 */
