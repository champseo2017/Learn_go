package main

import (
	"fmt"
)

/*
เรียกใช้ฟังก์ชัน recover โดยไม่ได้ใช้ร่วมกับ defer มันจะไม่มีผลอะไรเลย และจะ return ค่า nil เสมอ
*/

func fn1() {
	fmt.Println("Start of Fn1 Func")
	fmt.Println("Recovered", recover()) // recover จะ return nil เพราะไม่ได้ใช้ร่วมกับ defer
	defer func() {
		fmt.Println("Defer in Fn1 func")
	}()
	panic("Panicing from fn1") // เรียก panic ใน fn1
}

func main() {
	fmt.Println("Start of Main Func")
	defer func() {
		fmt.Println("Defer in Main func")
	}()
	fn1()
	fmt.Println("End of Main Func")
}

/*
- ฟังก์ชัน `recover` จะไม่มีผลใดๆ หากเรียกใช้โดยไม่ได้ใช้ร่วมกับ `defer` และจะ return ค่า `nil` เสมอ
- ใน `main` มีการเรียกใช้ `defer` เพื่อแสดงข้อความหลังจาก `main` จบการทำงาน
- เรียกใช้ `fn1` ซึ่งมีการเรียก `recover` แต่ไม่มีผล เพราะไม่ได้ใช้ร่วมกับ `defer`
- ใน `fn1` มีการเรียก `defer` เพื่อแสดงข้อความก่อนจบฟังก์ชัน และเรียก `panic` ทำให้โปรแกรมหยุดทำงานทันที
- `defer` ใน `main` จะทำงานก่อนโปรแกรมจบการทำงาน
*/
