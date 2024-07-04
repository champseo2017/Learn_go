package main

import (
	"fmt"
	// oops "GoProject/package"
)

/*
1. ในโค้ดมีการใช้คำสั่ง defer เพื่อให้โค้ดทำงานเป็นอันดับสุดท้าย และใช้ panic() เพื่อหยุดโปรแกรมเมื่อเกิดข้อผิดพลาด
2. defer มักใช้เพื่อปิดไฟล์หรือจัดการ I/O ต่างๆ ส่วน panic() ใช้เพื่อหยุดโปรแกรมเมื่อเจอข้อผิดพลาดร้ายแรง
3. Go มีคำสั่ง defer และ panic() เพื่อช่วยให้โปรแกรมทำงานได้อย่างสมบูรณ์และจัดการกับข้อผิดพลาดได้
*/

func main() {
	defer func() {
		fmt.Println("Code before panic")
		if r := recover(); r != nil {
			fmt.Println("Panic attack!!")
		}
	}()

	fmt.Println("Code after panic")
	panic("panic has been raised")
}

/*

 */
