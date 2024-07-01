package main

import (
	"fmt"
	// oops "GoProject/package"
)

/*
สรุปตัวอย่างการใช้งานฟังก์ชัน `deferExample()` และ `demoErrorHandling()` สั้นๆ ดังนี้

1. ฟังก์ชัน `deferExample()`:
   - แสดงลำดับการทำงานของคำสั่ง `defer`
   - คำสั่งที่อยู่หลัง `defer` จะถูกเลื่อนไปทำหลังสุด

2. ฟังก์ชัน `demoErrorHandling()`:
   - แสดงการจัดการ Error ด้วย `defer` และ `recover()`
   - เมื่อเกิด Panic (เช่น หารด้วย 0) `defer` จะทำงานทันทีและแสดงข้อความ Error

3. ฟังก์ชัน `main()`:
   - เรียกใช้ฟังก์ชัน `deferExample()` และ `demoErrorHandling()` เพื่อแสดงตัวอย่างการทำงาน

สรุปคือ `defer` ใช้เพื่อเลื่อนการทำงานของคำสั่งและใช้จัดการ Error ร่วมกับ `recover()` ส่วน `recover()` ใช้ดักจับ Panic ที่เกิดขึ้นระหว่างการทำงานของโปรแกรม

defer จะถูกจัดเก็บเป็น Stack และทำงานแบบ Last In First Out (LIFO) นั่นคือ defer ที่ถูกเรียกทีหลังสุดจะทำงานก่อน
*/

func deferExample() {
	fmt.Println("This is example 1")
	defer fmt.Println("This is example 2") // ถูกเรียกเป็นลำดับที่ 2
	fmt.Println("This is example 3")
}

func demoErrorHandling() {
	fmt.Println("----Demo Error Handling----")
	defer func() { // ถูกเรียกเป็นลำดับแรก
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	// เกิด Panic เนื่องจากหารด้วย 0
	valZero := 0
	result := 10 / valZero
	fmt.Println("Result:", result)
}

func main() {
	fmt.Println("==== defer Example ====")
	deferExample()

	fmt.Println("\n==== Error Handling Example ====")
	demoErrorHandling()
}

/*

 */
