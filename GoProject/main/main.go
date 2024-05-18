package main

import (
	"fmt"
)

/*
วิธีการเข้าถึงข้อมูลใน Map โดยใช้ Key ที่สอดคล้องกันในวงเล็บเหลี่ยม [] นอกจากนี้ เรายังสามารถตรวจสอบได้ว่า Key ที่ต้องการมีอยู่ใน Map หรือไม่ โดยการใช้ [] กับ Map จะคืนค่า 2 ค่า โดยค่าที่สองจะบอกว่า Key นั้นมีอยู่หรือไม่ เราสามารถตรวจสอบค่าตัวแปรนั้นก่อนที่จะใช้ Value
*/

func main() {
	// สร้าง Map โดยใช้ Map Literal
	m := map[int]string{1: "Udit", 2: "Raju"}

	// เข้าถึงและแสดงค่าใน Map โดยใช้ Key ในวงเล็บเหลี่ยม []
	fmt.Println(m[1], m[2])

	// เข้าถึงค่าใน Map โดยใช้ Key ในวงเล็บเหลี่ยม [] และกำหนดให้ตัวแปร data และ itExists
	data, itExists := m[3]

	// แสดงค่าของตัวแปร data และ itExists
	fmt.Println(data, itExists)

	// ตรวจสอบว่า Key มีอยู่ใน Map หรือไม่
	if itExists {
		// ถ้ามีอยู่ ให้แสดงค่า Value ที่สอดคล้องกับ Key
		fmt.Println(data)
	} else {
		// ถ้าไม่มีอยู่ ให้แสดงข้อความว่าข้อมูลไม่มีอยู่
		fmt.Println("the data does not exists")
	}
}

/*

 */
