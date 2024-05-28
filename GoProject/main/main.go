package main

import (
	"fmt"
)

/*
การเรียกใช้ฟังก์ชัน panic ในฟังก์ชัน fn ทำให้โปรแกรมหยุดทำงานเมื่อฟังก์ชัน fn ถูกเรียกใช้ ฟังก์ชัน panic จะทำให้โปรแกรมหยุดทันทีหลังจากเกิดข้อผิดพลาด และข้อความหลังจากนั้นจะไม่ถูกแสดงในคอนโซล
*/

// ฟังก์ชัน fn ที่ถูกเรียกใช้ใน main
func fn() {
	// แสดงข้อความเมื่อเริ่มฟังก์ชัน fn
	fmt.Println("Start of fn Func")

	// เรียกใช้ panic ทำให้โปรแกรมหยุดทำงานทันที
	panic("Calling panic explicitly")

	// ข้อความนี้จะไม่ถูกแสดงเนื่องจากโปรแกรมหยุดที่ panic
	fmt.Println("End of fn Func")
}

func main() {
	// แสดงข้อความเมื่อเริ่มฟังก์ชัน main
	fmt.Println("Start of Main Func")

	// เรียกใช้ฟังก์ชัน fn
	fn()

	// แสดงข้อความเมื่อสิ้นสุดฟังก์ชัน main
	fmt.Println("End of Main Func")
}

/*

 */
