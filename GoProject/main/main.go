package main

import (
	"fmt"
)

/*
คำสั่ง defer ใน Go จะสร้างฟังก์ชันขึ้นมาแต่จะไม่เรียกใช้จนกว่าฟังก์ชันที่สร้าง defer จะจบลง โดยจะเรียกใช้ตามลำดับกลับกันจากการสร้าง
*/

// ฟังก์ชัน fn1()
func fn1() {
	fmt.Println("Start of Fn1 Func") // พิมพ์ "Start of Fn1 Func" ก่อน

	// defer ฟังก์ชันถูกสร้างขึ้นและจะถูกเรียกเมื่อ fn1() จบลง
	defer func() {
		fmt.Println("Defer in Fn1 func")     // พิมพ์ "Defer in Fn1 func" ก่อน
		fmt.Println("Recovered:", recover()) // พิมพ์ "Recovered: <nil>" เนื่องจากไม่มี panic เกิดขึ้น
	}()

	fmt.Println("End of Fn1 Func") // พิมพ์ "End of Fn1 Func" ออกมาก่อนจบ fn1()
}

func main() {
	fmt.Println("Start of Main Func") // พิมพ์ "Start of Main Func" ออกมาก่อน

	// defer ฟังก์ชันถูกสร้างขึ้นและจะถูกเรียกเมื่อ main() จบลง
	defer func() {
		fmt.Println("Defer in Main func")
	}()

	fn1() // เรียกใช้ fn1()

	fmt.Println("End of Main Func") // พิมพ์ "End of Main Func" ออกมาก่อนจบ main()
}

/*

 */
