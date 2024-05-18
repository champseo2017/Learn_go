package main

import (
	"fmt"
)

/*
Map Literal และการใช้ฟังก์ชัน `make(map[string]int)` เป็นวิธีการสร้าง Map ใน Go ที่แตกต่างกันดังนี้

1. การกำหนดค่าเริ่มต้น:
   - Map Literal: สามารถกำหนดค่าเริ่มต้นให้กับ Map ได้ในขณะที่ประกาศ โดยระบุคู่ Key-Value ภายในเครื่องหมายปีกกา `{}`
   - `make(map[string]int)`: ไม่สามารถกำหนดค่าเริ่มต้นให้กับ Map ได้ในขณะที่ประกาศ Map จะถูกสร้างเป็น Empty Map

2. ความสะดวกในการสร้าง:
   - Map Literal: เป็นวิธีที่สะดวกในการสร้างและกำหนดค่าเริ่มต้นให้กับ Map ในคำสั่งเดียว
   - `make(map[string]int)`: ต้องสร้าง Map ก่อน แล้วจึงกำหนดค่าให้กับ Map ในภายหลัง

3. การระบุขนาดเริ่มต้น:
   - Map Literal: ไม่สามารถระบุขนาดเริ่มต้นของ Map ได้ ขนาดของ Map จะถูกกำหนดโดยจำนวนคู่ Key-Value ที่กำหนดในขณะสร้าง
   - `make(map[string]int)`: สามารถระบุขนาดเริ่มต้นของ Map ได้โดยการส่งพารามิเตอร์เพิ่มเติม เช่น `make(map[string]int, 10)` จะสร้าง Map ที่มีขนาดเริ่มต้นเป็น 10

   ตัวอย่างการสร้าง Map ด้วย Map Literal:

   m := map[string]int{"A": 1, "B": 2, "C": 3}

   ตัวอย่างการสร้าง Map ด้วยฟังก์ชัน `make`:
    m := make(map[string]int)
	m["A"] = 1
	m["B"] = 2
	m["C"] = 3
*/

func Create_Map() {
	// สร้าง Map โดยใช้ฟังก์ชัน make โดยระบุประเภทของ Key เป็น string และประเภทของ Value เป็น int
	m := make(map[string]int)

	// แสดงค่าของ Map ที่สร้างขึ้น
	fmt.Println(m)
}

func Create_Map_Using_Map_Literals() {
	// สร้าง Map โดยใช้ Map Literal โดยระบุประเภทของ Key เป็น int และประเภทของ Value เป็น string
	// และกำหนดค่าเริ่มต้นให้กับ Map โดยใช้ Key เป็น 1 และ 2 และ Value เป็น "Udit" และ "Raju" ตามลำดับ
	m := map[int]string{1: "Udit", 2: "Raju"}

	// แสดงค่าของ Map ที่สร้างขึ้น
	fmt.Println(m)
}

func main() {
	Create_Map()
	Create_Map_Using_Map_Literals()
}

/*

 */
