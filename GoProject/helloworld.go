package main

import (
	"fmt"
)

/*

ในภาษา Go, ประเภทข้อมูลบูลีน (Boolean) ใช้สำหรับเก็บค่าความจริงที่มีเพียงสองค่าเท่านั้น คือ true (จริง) และ false (เท็จ)


*/

func main() {
   
	var b bool // ประกาศตัวแปร b ประเภทบูลีน
	fmt.Println(b) // ค่าเริ่มต้นของบูลีนคือ false

	b = true // กำหนดค่าให้ตัวแปร b เป็น true
	fmt.Println(b) // แสดงค่าของ b ซึ่งตอนนี้คือ true
	
}
