package main

import "fmt"

/*
Structs ใน Go เป็น value types หมายความว่าเมื่อมีการสร้าง instance ของ struct จะมีการจองพื้นที่ในหน่วยความจำเพื่อเก็บค่าของ struct นั้น และเมื่อมีการกำหนดค่า struct หนึ่งให้กับอีก struct หนึ่ง หรือส่ง struct ไปยังฟังก์ชันอื่น จะมีการสร้างสำเนาของ struct ขึ้นมาใหม่
ตัวอย่างเช่น ในโปรแกรมที่ให้มา มีการสร้าง struct ชื่อ city และ instance ของ city ชื่อ c1 และ c2 เมื่อกำหนดค่า c1 ให้กับ c2 จะมีการสร้างสำเนาของ c1 และกำหนดให้กับ c2 ดังนั้น c1 และ c2 จะเป็น struct ที่แยกจากกันในหน่วยความจำ การแก้ไขค่าในฟิลด์ของ c2 จะไม่ส่งผลต่อค่าใน c1
การทำความเข้าใจหลักการของ value types มีความสำคัญในการออกแบบและพัฒนาโปรแกรม Go เพื่อให้สามารถจัดการกับข้อมูลได้อย่างถูกต้องและหลีกเลี่ยงข้อผิดพลาดที่อาจเกิดขึ้นจากการแชร์ข้อมูลระหว่างส่วนต่างๆ ของโปรแกรมโดยไม่ตั้งใจ
*/
// สร้าง struct ชื่อ city ที่มีฟิลด์ city1, city2, และ city3 เป็น string
type city struct {
	city1 string
	city2 string
	city3 string
}
func main() {
	// สร้าง instance ของ city ชื่อ c1 และกำหนดค่าเริ่มต้นให้กับฟิลด์
	c1 := city{"Pune", "Delhi", "Lucknow"}
	// กำหนดค่า c1 ให้กับ c2 (สร้างสำเนาของ c1 และกำหนดให้กับ c2)
	c2 := c1
	// แสดงค่าของ c1 และ c2
	fmt.Println("city list c1 =", c1)
	fmt.Println("city list c2 =", c2)
	// แก้ไขค่าของฟิลด์ city2 และ city3 ใน c2
	c2.city2 = "Bengaluru"
	c2.city3 = "Mumbai"
	// แสดงค่าของ c1 และ c2 หลังจากแก้ไข c2
	fmt.Println("\nAfter modifying c2:")
	fmt.Println("city list c1 =", c1)
	fmt.Println("city list c2 =", c2)
}
/* 

*/