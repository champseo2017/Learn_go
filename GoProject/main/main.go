package main

import (
	"fmt"
)

/*
แสดงวิธีการลบข้อมูลออกจากแมป (map) ในภาษา Go โดยใช้คำสั่ง delete พร้อมระบุแมปและคีย์ที่ต้องการลบ
*/

// ฟังก์ชันสำหรับลบคีย์จากแมป
func Delete_Key_From_Map() {
	// สร้างแมปที่มีคีย์เป็นจำนวนเต็มและค่าเป็นสตริง
	m := map[int]string{1: "Udit", 2: "Raju"}

	// แสดงค่าในแมปก่อนการลบ
	fmt.Println("the map values without deletion is", m)

	// ลบคีย์ 1 ออกจากแมป
	delete(m, 1)

	// แสดงค่าในแมปหลังการลบ
	fmt.Println("the deleted map is", m)
}

func main() {
	Delete_Key_From_Map()
}

/*
โค้ดนี้ทำงานโดยการสร้างฟังก์ชัน Delete_Key_From_Map() ซึ่งภายในฟังก์ชันจะสร้างแมปที่มีคีย์เป็นจำนวนเต็มและค่าเป็นสตริง จากนั้นแสดงค่าในแมปก่อนการลบ ต่อมาใช้คำสั่ง delete(m, 1) เพื่อลบคีย์ 1 ออกจากแมป m และสุดท้ายแสดงค่าในแมปหลังการลบ เมื่อรันโค้ดจะได้ผลลัพธ์เป็นค่าในแมปก่อนการลบและหลังการลบคีย์ 1 ออกไป
*/
