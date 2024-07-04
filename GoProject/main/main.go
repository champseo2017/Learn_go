package main

import (
	"fmt"
	// oops "GoProject/package"
)

/*
1. ในโค้ดมีการใช้ฟังก์ชัน recover() ร่วมกับ defer เพื่อดักจับและจัดการกับ panic ที่เกิดขึ้น
2. Panic คือข้อผิดพลาดร้ายแรงที่ทำให้โปรแกรมหยุดทำงานทันที แต่เราสามารถใช้ recover() เพื่อกู้คืนและจัดการกับมันได้
3. การใช้ recover() ช่วยให้เราสามารถดักจับ panic, แสดงข้อความแจ้งเตือน, หรือทำความสะอาดบางอย่างก่อนที่โปรแกรมจะหยุดทำงานไป ทำให้เราจัดการกับ panic ได้อย่างมีประสิทธิภาพ
*/

func null() {
	// ฟังก์ชันนี้จะทำให้เกิด panic เสมอ
	panic("null func panic")
}

func main() {
	// ใช้ defer เพื่อให้ฟังก์ชันนี้ทำงานเป็นอันดับสุดท้าย
	defer func() {
		// ใช้ recover() เพื่อดักจับ panic ที่เกิดขึ้น
		if err := recover(); err != nil {
			// ถ้ามี panic เกิดขึ้น ให้แสดงข้อความและค่า error
			fmt.Println("panic occurred:", err)
		}
	}()

	fmt.Println("Although panicked, we recovered, we call null() func")
	// เรียกฟังก์ชัน null() ซึ่งจะเกิด panic
	null()
	// โค้ดข้างล่างนี้จะไม่ทำงาน เพราะเกิด panic ไปแล้ว
	fmt.Println("will not print this")
}

/*

 */
