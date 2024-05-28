package main

import (
	"fmt"
)

/*
ฟังก์ชัน recover ที่มีอยู่ในตัว Go ช่วยให้เราสามารถกู้คืนจากการเกิด panic ได้ โปรแกรมจะดำเนินการต่อหลังจากเรียก recover ควรเรียก recover จากฟังก์ชันที่ถูก defer ไว้ มิฉะนั้นจะไม่ส่งผลใดๆ

ฟังก์ชัน recover สามารถถูกเรียกจากฟังก์ชันใดๆ ใน call stack ไม่จำเป็นต้องเรียกจากฟังก์ชันเดียวกับที่เกิด panic ในตัวอย่างโค้ดด้านล่าง ฟังก์ชัน recover ถูกเรียกในฟังก์ชัน fn2 หลังจากเรียก fn2 โปรแกรมจะไม่หยุดทำงาน แต่จะดำเนินการต่อจนจบโปรแกรมอย่างถูกต้อง
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

	// แสดงข้อความเมื่อสิ้นสุดฟังก์ชัน fn1
	fmt.Println("End of Fn1 Func")
}

func fn2() {
	// เริ่มฟังก์ชัน fn2
	fmt.Println("Start of Fn2 Func")

	// ฟังก์ชัน defer ที่จะถูกเรียกเมื่อสิ้นสุดฟังก์ชัน fn2
	defer func() {
		fmt.Println("Defer in Fn2 func")
		// กู้คืนจาก panic ถ้ามีการเรียก panic
		if rec := recover(); rec != nil {
			fmt.Println("Recovered:", rec)
		}
	}()

	// เรียกฟังก์ชัน fn3
	fn3()

	// ข้อความนี้จะไม่ถูกแสดงเนื่องจากฟังก์ชัน recover จะไม่ทำให้โปรแกรมหยุด
	fmt.Println("End of Fn2 Func")
}

func fn3() {
	// เริ่มฟังก์ชัน fn3
	fmt.Println("Start of Fn3 Func")

	// ฟังก์ชัน defer ที่จะถูกเรียกเมื่อสิ้นสุดฟังก์ชัน fn3
	defer func() {
		fmt.Println("Defer in Fn3 func")
	}()

	// เรียกใช้ panic ทำให้โปรแกรมหยุดทำงานทันทีด้วยข้อความ "Calling panic explicitly in Fn3"
	panic("Calling panic explicitly in Fn3")
}

func main() {
	// แสดงข้อความเมื่อเริ่มฟังก์ชัน main
	fmt.Println("Start of Main Func")

	// ฟังก์ชัน defer ที่จะถูกเรียกเมื่อสิ้นสุดฟังก์ชัน main
	defer func() {
		fmt.Println("Defer in Main func")
	}()

	// เรียกฟังก์ชัน fn1
	fn1()

	// แสดงข้อความเมื่อสิ้นสุดฟังก์ชัน main
	fmt.Println("End of Main Func")
}

/*
หลังจาก Recovered: Calling panic explicitly in Fn3 ข้อความ Defer in Fn3 func ถูกแสดงก่อนการกู้คืน (recover) เนื่องจาก defer ใน fn3 ถูกเรียกก่อนที่ panic จะถูกกู้คืนใน fn2 นั่นเอง
*/
