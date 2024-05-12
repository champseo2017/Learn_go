package main

import (
	"fmt"
)

/*
วิธีแก้ไขปัญหานี้คือใช้ pointer ชี้ไปยัง slice แทนการส่ง slice เข้าไปใน function โดยตรง
*/

func modify(s1 *[]int) {
	(*s1)[0] = 100
	// แก้ไขค่าสมาชิกตำแหน่งที่ 0 ของ slice ที่ pointer s1 ชี้ไปเป็น 100

	*s1 = append(*s1, 60)
	// เพิ่มสมาชิก 60 เข้าไปใน slice ที่ pointer s1 ชี้ไป

	*s1 = append(*s1, 70)
	// เพิ่มสมาชิก 70 เข้าไปใน slice ที่ pointer s1 ชี้ไป
}

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	// สร้าง slice s1 ที่มีสมาชิกเป็น 10, 20, 30, 40, 50

	modify(&s1)
	// เรียกใช้ function modify โดยส่ง pointer ของ s1 เป็นอาร์กิวเมนต์

	fmt.Println("Slice Elements", s1)
	// แสดงผลสมาชิกของ s1 หลังจากเรียกใช้ function modify
}

/*
เราส่ง pointer ของ s1 เข้าไปใน function modify แทนที่จะส่ง s1 โดยตรง ดังนั้นเมื่อมีการแก้ไขค่าสมาชิกหรือเพิ่มสมาชิกใหม่ผ่าน pointer ใน function จะส่งผลโดยตรงต่อ slice s1 ใน main function ด้วย
*/
