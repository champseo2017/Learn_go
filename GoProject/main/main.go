package main

import (
	"fmt"
)

/*
ในภาษา Go มีวิธีการส่งค่าเข้าไปยังฟังก์ชันได้ 2 วิธี คือ Pass by Value และ Pass by Reference

1. Pass by Value เป็นวิธีมาตรฐานที่ Go ใช้ในการส่งค่าเข้าฟังก์ชัน โดยจะทำการคัดลอกค่าจากตัวแปรต้นทางมาสร้างเป็นตัวแปรใหม่ แล้วจึงส่งค่าตัวแปรใหม่นี้เข้าไปยังฟังก์ชัน ทำให้ค่าเดิมในตัวแปรต้นทางไม่มีการเปลี่ยนแปลง
*/

// ฟังก์ชัน Add_One รับพารามิเตอร์ val ประเภท int
// และส่งคืนค่า val + 1
func Add_One(val int) int {
	return val + 1
}

func main() {
	// กำหนดค่าเริ่มต้นให้ตัวแปร res เท่ากับ 1
	res := 1

	// เรียกใช้ฟังก์ชัน Add_One พร้อมส่งค่า res เข้าไปเป็นพารามิเตอร์
	// และเก็บค่าที่ได้รับกลับมาไว้ในตัวแปร new_res
	new_res := Add_One(res)

	// แสดงค่าของ res ซึ่งยังคงมีค่าเดิมหลังจากเรียกใช้ฟังก์ชัน
	fmt.Println("The value of res is preserved after function execution", res)

	// แสดงค่าใหม่ที่ได้รับกลับมาจากฟังก์ชัน
	fmt.Println("The new value is returned", new_res)
}

/*

 */
