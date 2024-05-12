package main

import (
	"fmt"
)

/*
เกี่ยวกับ Pointer ในภาษา Go:

1. Pointer คือตัวแปรที่เก็บ memory address ของตัวแปรอื่น
2. ใช้ & เพื่อเข้าถึง address ของตัวแปร และใช้ * เพื่อเข้าถึงค่าที่ address นั้น
3. ตัวแปร pointer ต้องได้รับการกำหนดค่า address ของตัวแปรอื่นก่อนเข้าถึงค่า
4. ชนิดของตัวแปร pointer ต้องตรงกับชนิดของตัวแปรที่มันชี้ไป
5. ตัวดำเนินการเปรียบเทียบเพียงตัวเดียวที่ใช้กับ pointer คือ equality
6. แนะนำให้ส่ง slice ของ array เข้าไปในฟังก์ชันแทนการใช้ pointer ของ array
7. หากต้องการเพิ่มค่าใน slice ที่เป็น argument ของฟังก์ชัน ให้กำหนด argument เป็น pointer ไปยัง slice
8. เมื่อเปลี่ยนแปลงค่าของตัวแปรผ่าน pointer ในฟังก์ชัน การเปลี่ยนแปลงจะสะท้อนกลับมาที่ตัวแปรต้นฉบับด้วย
*/

type person struct {
	id   int
	name string
}

func modify(p1 *person) {
	p1.name = "NewXYZ"
}

func main() {
	p1 := person{101, "XYZ"}
	modify(&p1)
	fmt.Println("Person p1:", p1)
}

/*

 */
