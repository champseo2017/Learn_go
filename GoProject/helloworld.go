package main

import "fmt"

/*
โปรแกรมนี้แสดงให้เห็นถึงการอัปเดตค่าของสมาชิกในโครงสร้าง (struct) โดยใช้พอยน์เตอร์ในภาษา Go โดยมีขั้นตอนดังนี้
*/
// สร้างโครงสร้าง my_struct ที่มีสมาชิก city และ country เป็นสตริง
type my_struct struct {
	city string
	country string
}

func main() {
	// สร้างอินสแตนซ์ของ my_struct ชื่อ my_struct1 และกำหนดค่าเริ่มต้น
	my_struct1 := my_struct{"New Delhi", "India"}
	// สร้างพอยน์เตอร์ my_pointer ชี้ไปยัง my_struct1
	my_pointer := &my_struct1

	// แสดงค่าของ my_pointer ก่อนการอัปเดต
	fmt.Println("Before updating value of struct member:\n", my_pointer)
	// อัปเดตค่า city ของ my_struct1 ผ่านทาง my_pointer เป็น "Pune"
	my_pointer.city = "Pune"
	// แสดงค่าของ my_pointer หลังการอัปเดต
	fmt.Println("After updating value of struct member:\n", my_pointer)
}

/* 
จากตัวอย่างโค้ดข้างต้น สังเกตได้ว่าเราสามารถอัปเดตค่าของสมาชิกในโครงสร้างผ่านทางพอยน์เตอร์ได้โดยตรง โดยใช้ my_pointer.city = "Pune" ซึ่งจะเปลี่ยนแปลงค่าของ city ในอินสแตนซ์ my_struct1 จาก "New Delhi" เป็น "Pune"
การใช้พอยน์เตอร์ในการอัปเดตค่าของสมาชิกในโครงสร้างนี้ ช่วยให้เราสามารถแก้ไขค่าของอินสแตนซ์โครงสร้างได้โดยตรงโดยไม่ต้องสร้างอินสแตนซ์ใหม่ ซึ่งเป็นประโยชน์ในการจัดการและปรับปรุงข้อมูลในโปรแกรม
*/