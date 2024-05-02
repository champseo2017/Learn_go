package main

import "fmt"

/*
เมื่อมีการกำหนด struct แต่ไม่ได้กำหนดค่าเริ่มต้นให้อย่างชัดเจน ฟิลด์ของ struct จะได้รับการกำหนดค่า zero โดยค่าเริ่มต้น (zero values) ของแต่ละประเภทข้อมูล เช่น สำหรับ string จะเป็น "" (สตริงว่าง) และสำหรับ int จะเป็น 0 นอกจากนี้ เรายังสามารถกำหนดค่าให้กับบางฟิลด์และละเว้นฟิลด์ที่เหลือได้ ซึ่งฟิลด์ที่ละเว้นจะได้รับการกำหนดค่า zero โดยอัตโนมัติ
*/

type employee struct {
	name string
	age int
	city string
	location string
}

func main() {
	// var emp1 employee // struct ที่มีค่า zero
	// fmt.Println("Employee1 details:", emp1)
	/* 
	- เรากำหนด struct ชื่อ `employee` ที่มีฟิลด์ `name` เป็น string, `age` เป็น int และ `city` เป็น string
	- ในฟังก์ชัน `main` เราประกาศตัวแปร `emp1` ของ struct `employee` แต่ไม่ได้กำหนดค่าเริ่มต้น
	- ดังนั้น ฟิลด์ `name` และ `city` จะได้รับค่า zero ของ string ซึ่งเป็น `""` (สตริงว่าง) และฟิลด์ `age` จะได้รับค่า zero ของ int ซึ่งเป็น `0`
	- ผลลัพธ์ที่ได้จะเป็น `Employee1 details: {0}`
	*/
	emp1 := employee {
		age: 30,
		location: "Pune",
	}
	fmt.Println("Employee1 details:", emp1)
}

/* 
- เรากำหนด struct ชื่อ `employee` ที่มีฟิลด์ `name` เป็น string, `age` เป็น int และ `location` เป็น string
- ในฟังก์ชัน `main` เราประกาศตัวแปร `emp1` ของ struct `employee` และกำหนดค่าเริ่มต้นโดยระบุเฉพาะฟิลด์ `age` และ `location`
- ฟิลด์ `name` ที่ไม่ได้ระบุค่าจะได้รับค่า zero ของ string ซึ่งเป็น `""` (สตริงว่าง)
- ผลลัพธ์ที่ได้จะเป็น `Employee1 details: {30 Pune}`

การทำความเข้าใจเกี่ยวกับ zero values ของ struct เป็นสิ่งสำคัญ เนื่องจากเมื่อเราประกาศตัวแปรของ struct โดยไม่ได้กำหนดค่าเริ่มต้น ฟิลด์ของ struct จะได้รับค่า zero โดยอัตโนมัติตามประเภทข้อมูลของแต่ละฟิลด์ นอกจากนี้ เรายังสามารถกำหนดค่าให้กับบางฟิลด์และละเว้นฟิลด์ที่เหลือได้ ซึ่งฟิลด์ที่ละเว้นจะได้รับค่า zero อยู่เสมอ
*/