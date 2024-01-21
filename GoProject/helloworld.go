package main

import (
	"fmt"
)

func main() {
	// ประกาศตัวแปรเป็นชนิดข้อมูลแบบ byte และ rune
	s1 := "abcd"
	s2 := "Hello"

	byte_s := []byte(s1) // แปลง s1 เป็น slice ของ byte
	rune_s := []rune(s2) // แปลง s2 เป็น slice ของ rune

	// ใช้ fmt.Printf สำหรับการจัดรูปแบบข้อความ
	// %c แสดงข้อมูลในรูปแบบของ character หรือ rune
	// %T แสดงประเภทของตัวแปร
	// %v แสดงค่าของตัวแปรในรูปแบบ default
	fmt.Printf("Type of %c is %T and value is %v\n", byte_s, byte_s, byte_s)
	fmt.Printf("Type of %c is %T and value is %v\n", rune_s, rune_s, rune_s)
}
