package main

import (
	"fmt"
	// oops "GoProject/package"
)

/*
สรุปสั้นๆ เกี่ยวกับ Error Handling ในภาษา Go

- เมื่อเกิดข้อผิดพลาดขึ้นระหว่างการทำงานของโปรแกรม Go จะสร้าง Panic ขึ้นมา
- เราสามารถใช้ `defer` และ `recover()` เพื่อดักจับและจัดการกับ Panic ได้
- `defer` จะทำให้โค้ดภายในทำงานทันทีเมื่อฟังก์ชันเสร็จสิ้น ไม่ว่าจะเกิด Panic หรือไม่ก็ตาม
- `recover()` จะคืนค่า `nil` ถ้าไม่เกิด Panic แต่ถ้าเกิด Panic จะคืนค่า Error ที่เกิดขึ้น
- เราสามารถแสดงข้อความ Error ให้ผู้ใช้เห็นได้ เมื่อเกิด Panic ขึ้น

ด้วยวิธีนี้ โปรแกรมจะไม่หยุดทำงานกะทันหันเมื่อเจอข้อผิดพลาด และเราสามารถจัดการกับ Error ได้อย่างเหมาะสมครับ
*/

func calculate() {
	// ใช้ defer และ recover() เพื่อดักจับและจัดการกับ Panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("----demo error handling----")
			fmt.Println(err)
		}
	}()

	a := 10
	b := 0
	c := 0

	// คำนวณ c = a / b ซึ่งจะเกิด Panic เพราะหารด้วย 0 ไม่ได้
	c = a / b
	fmt.Println(c)
}

func main() {
	fmt.Println("Error handling")

	// เรียกใช้ฟังก์ชัน calculate() ภายใน Goroutine
	go calculate()
}

/*

 */
