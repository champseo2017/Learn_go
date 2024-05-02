package main

import "fmt"

/*
ในภาษา Go, Pointer เป็นตัวแปรที่ใช้สำหรับเก็บที่อยู่ของตัวแปรอื่น โดยตัวแปรจะใช้สำหรับเก็บข้อมูลที่ที่อยู่ในหน่วยความจำของระบบ เราสามารถใช้ Pointer กับ Struct ได้โดยใช้เครื่องหมาย & (Address Operator) ภาษา Go อนุญาตให้โปรแกรมเมอร์เข้าถึง Field ของ Struct โดยใช้ Pointer โดยไม่จำเป็นต้องมีการ Dereference อย่างชัดเจน
*/

type my_struct struct {
	city string
	country string
}

func main() {
	// สร้าง Instance ของ my_struct
	my_struct1 := my_struct{"New Delhi", "India"}
	// Pointer ไปยัง Struct
	my_pointer := &my_struct1
	fmt.Println(my_pointer)

	// เข้าถึง Field ของ Struct โดยใช้ Pointer
	fmt.Println(my_pointer.country)

	// เข้าถึง Field ของ Struct โดยใช้ Concept ของ Dereferencing
	fmt.Println((*my_pointer).country)

}

/* 
1. เราประกาศ Struct ชื่อ `my_struct` ที่มี Field `city` และ `country` เป็น string
2. ในฟังก์ชัน `main()` เราสร้าง Instance ของ `my_struct` ชื่อ `my_struct1` และกำหนดค่าเริ่มต้นให้กับ Field `city` เป็น "New Delhi" และ `country` เป็น "India"
3. เราสร้าง Pointer ชื่อ `my_pointer` และกำหนดให้ชี้ไปยัง `my_struct1` โดยใช้เครื่องหมาย & (Address Operator)
4. เราแสดงค่าของ Pointer `my_pointer` ซึ่งจะเป็นที่อยู่ของ `my_struct1`
5. เราเข้าถึง Field `country` ของ `my_struct1` ผ่าน Pointer `my_pointer` โดยใช้เครื่องหมายจุด (Dot Notation) โดยไม่จำเป็นต้อง Dereference Pointer
6. เราเข้าถึง Field `country` ของ `my_struct1` โดยใช้ Concept ของ Dereferencing `(*my_pointer).country` ซึ่งจะให้ผลลัพธ์เหมือนกับการเข้าถึงโดยใช้ Dot Notation

นอกจากนี้ยังช่วยประหยัดหน่วยความจำเนื่องจาก Pointer จะอ้างอิงไปยังที่อยู่ของ Struct โดยตรง ไม่ต้องสร้างสำเนาของ Struct ใหม่
*/