package main

import "fmt"

/*
Maps ใน Golang เป็นการรวบรวมคู่ key-value ที่ไม่เรียงลำดับ ช่วยให้สามารถค้นหา ดึง อัปเดต หรือลบค่าต่างๆ ได้อย่างรวดเร็วด้วยการอ้างอิงผ่าน keys เมื่อมีคู่ key และ value สามารถเก็บ value ลงใน map object ได้ หลังจากเก็บ value แล้ว ก็สามารถเรียกค่าใน map ได้โดยอ้างอิงผ่าน keys ที่เกี่ยวข้อง

โครงสร้างของบทนี้ประกอบด้วย:
- Maps ใน Go คืออะไร
- การประกาศ Maps ใน Go
- การกำหนดค่าเริ่มต้นให้ Maps ใน Go
- การวนลูปเพื่ออ่านค่าใน Maps
- การเพิ่มคู่ key-value ใน Maps
- การอัปเดตคู่ key-value ใน Maps
- การดึงค่า value ของ key ใน Maps
- การตรวจสอบว่ามี key ใน Maps หรือไม่
- การลบ key ออกจาก Maps
- Maps เป็น reference types

วัตถุประสงค์ของบทนี้คือ อธิบายแนวคิดพื้นฐานของ map และวิธีใช้งานใน Go โดยครอบคลุมการประกาศ กำหนดค่าเริ่มต้น วนลูป ดึงค่า อัปเดต ลบ และตรวจสอบการมีอยู่ของ key ใน map Maps ใช้ในการค้นหา value ผ่าน key ที่เกี่ยวข้อง
*/

func main() {
	// ประกาศ map ว่างเปล่าชื่อ m1 ที่มี key เป็น string และ value เป็น int
	// var m1 map[string]int
	// ประกาศและกำหนดค่าเริ่มต้นให้ map m2
	m2 := map[string]int {
		"apple": 5,
		"banana": 10,
		"oraange": 2,
	}
	// เข้าถึงค่า value ของ key "apple"
	fmt.Println(m2["apple"]) // ผลลัพธ์: 5
	// เพิ่มคู่ key-value ใหม่
	m2["grape"] = 7
	// อัปเดตค่า value ของ key ที่มีอยู่
	m2["banana"] = 8
	// วนลูปอ่านค่าทั้งหมดใน map
	for key, value := range m2 {
		fmt.Printf("%s: %d\n", key, value)
	}

	// ตรวจสอบว่ามี key ใน map หรือไม่
	// comma ok idiom (comma ok pattern)
	/* 
	- `m2["mango"]` เป็นการพยายามเข้าถึงค่าใน map `m2` ด้วย key "mango"
	- `:=` เป็นการประกาศตัวแปรและกำหนดค่าให้กับตัวแปรนั้นในคำสั่งเดียว
	- `_` (underscore) เป็นตัวแปรที่ไม่ได้ใช้งาน ใช้เพื่อละเว้นค่าที่ส่งกลับมาตัวแรก (ในที่นี้คือค่าของ key "mango") เพราะเราสนใจแค่ตัวที่สอง
	- `ok` เป็นตัวแปรบูลีนที่ระบุว่า key มีอยู่ใน map หรือไม่ ถ้ามี `ok` จะเป็น `true` ถ้าไม่มี จะเป็น `false`

	value := m2["mango"]
	value จะได้รับค่าแรกที่ส่งกลับมาเท่านั้น ซึ่งอาจเป็นค่าจริงของ key "mango" หรือค่า zero value ก็ได้ ขึ้นอยู่กับว่า "mango" มีอยู่ใน m2 หรือไม่ แต่เราจะไม่รู้ว่า key มีอยู่จริงหรือเปล่า
	*/
	if _, ok := m2["mango"]; ok {
		fmt.Println("Mango exists")
	} else {
		fmt.Println("Mango does not exist")
	}
	// ลบ key ออกจาก map
	delete(m2, "orange")

}
/* 

*/