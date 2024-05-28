package main

import (
	"fmt"
)

/*
เมื่อโปรแกรมเกิด panic การทำงานจะหยุดและแสดง stack trace ฟังก์ชัน defer จะทำงานเมื่อสิ้นสุดการเรียกใช้ฟังก์ชัน ถ้าโปรแกรมมีฟังก์ชัน defer ใน call stack ฟังก์ชัน defer จะถูกเรียกก่อนที่โปรแกรมจะหยุดทำงาน
*/

func fn1() {
	// เริ่มฟังก์ชัน fn1
	fmt.Println("Start of Fn1 Func")

	// ฟังก์ชัน defer ที่จะถูกเรียกเมื่อสิ้นสุดฟังก์ชัน fn1
	defer func() {
		fmt.Println("Defer in Fn1 func")
	}()

	// เรียกฟังก์ชัน fn2
	fn2()

	// ข้อความนี้จะไม่ถูกแสดงเนื่องจากโปรแกรมหยุดที่ panic ใน fn3
	fmt.Println("End of Fn1 Func")
}

func fn2() {
	// เริ่มฟังก์ชัน fn2
	fmt.Println("Start of Fn2 Func")

	// ฟังก์ชัน defer ที่จะถูกเรียกเมื่อสิ้นสุดฟังก์ชัน fn2
	defer func() {
		fmt.Println("Defer in Fn2 func")
	}()

	// เรียกฟังก์ชัน fn3
	fn3()

	// ข้อความนี้จะไม่ถูกแสดงเนื่องจากโปรแกรมหยุดที่ panic ใน fn3
	fmt.Println("End of Fn2 Func")
}

func fn3() {
	// เริ่มฟังก์ชัน fn3
	fmt.Println("Start of Fn3 Func")

	// ฟังก์ชัน defer ที่จะถูกเรียกเมื่อสิ้นสุดฟังก์ชัน fn3
	defer func() {
		fmt.Println("Defer in Fn3 func")
	}()

	// เรียกใช้ panic ทำให้โปรแกรมหยุดทำงานทันที
	panic("Calling panic explicitly in Fn3")
}

func main() {
	// เริ่มฟังก์ชัน main
	fmt.Println("Start of Main Func")

	// ฟังก์ชัน defer ที่จะถูกเรียกเมื่อสิ้นสุดฟังก์ชัน main
	defer func() {
		fmt.Println("Defer in Main func")
	}()

	// เรียกฟังก์ชัน fn1
	fn1()

	// ข้อความนี้จะไม่ถูกแสดงเนื่องจากโปรแกรมหยุดที่ panic ใน fn3
	fmt.Println("End of Main Func")
}

/*
แสดงว่าฟังก์ชัน defer ในแต่ละฟังก์ชันถูกเรียกใช้งานในลำดับย้อนกลับก่อนที่โปรแกรมจะหยุดการทำงาน
*/
