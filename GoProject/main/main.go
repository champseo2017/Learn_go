package main

import (
	oops "GoProject/package"
	"fmt"
)

/*
Interface ในภาษา Go คือชุดของเมธอดหรือฟังก์ชัน โดยประกอบด้วยชื่อ พารามิเตอร์ และค่าที่ส่งกลับ ใช้สำหรับแสดงความคล้ายคลึงกันในแนวคิดของประเภทข้อมูล (types) ต่างๆ
*/

func main() {
	square_area := oops.Square{Side: 1}
	rectangle_area := oops.Rectangle{Width: 3, Height: 4}

	shapes := []oops.Shape{square_area, rectangle_area}

	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}

/*
1. `Shape` คือ interface ที่มีเมธอด `Area()` เพื่อคำนวณพื้นที่
2. `Square` และ `Rectangle` เป็น struct ที่มีการนิยามเมธอด `Area()` ตามข้อกำหนดของ interface `Shape`
3. ในฟังก์ชัน `main()` สร้างตัวแปร `square_area` และ `rectangle_area` ซึ่งเป็นออบเจกต์ของ `Square` และ `Rectangle` ตามลำดับ
4. จัดเก็บ `square_area` และ `rectangle_area` ในอาร์เรย์ของ `Shape`
5. วนลูปเรียกใช้เมธอด `Area()` ของแต่ละออบเจกต์ในอาร์เรย์

เมื่อเรียกใช้ `shape.Area()` ภายในลูป Go จะเรียกใช้เมธอด `Area()` ของออบเจกต์นั้นๆ โดยอัตโนมัติ เช่น สำหรับ `square_area` จะเรียกใช้เมธอด `Area()` ของ `Square` เป็นต้น

ด้วยการใช้ interface ทำให้เราสามารถเรียกใช้ฟังก์ชันเดียวกัน (เช่น `Area()`) กับออบเจกต์ที่แตกต่างกันได้ ซึ่งเป็นการสร้างพอลิมอร์ฟิซึมในการเขียนโปรแกรม
*/
