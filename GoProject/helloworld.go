package main

import (
	"fmt"
)

/*
ในภาษา Go, ข้อมูลชนิด string ถือว่าเป็น "immutable" หรือ ไม่สามารถเปลี่ยนแปลงได้

*/

func main() {
	// ประกาศตัวแปรเป็นชนิดข้อมูลสตริง
	str := "Hello, world!"
	// ลองเปลี่ยนแปลง character ใน string
	// บรรทัดต่อไปนี้จะทำให้เกิด error เนื่องจาก string ใน Go เป็น immutable
	// str[0] = 'h'  // ไม่สามารถทำได้

	// วิธีที่ถูกต้องในการเปลี่ยนแปลง string คือการสร้าง string ใหม่
	newStr := "h" + str[1:]

	x := "\"hello golang\"\n I Love"
	y := `"Hello golang"\n I Love`

	fmt.Println(x)
	fmt.Println()
	fmt.Println(y)
	fmt.Println(newStr) // จะแสดงผล "hello, world!"

}
