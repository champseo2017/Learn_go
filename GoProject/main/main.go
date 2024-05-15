package main

import (
	"fmt"
	"time"
)

/*
ในการเริ่มต้นฟังก์ชันเป็น goroutine ใหม่ เราต้องเพิ่มคีย์เวิร์ด go ไว้ด้านหน้าการเรียกฟังก์ชัน ซึ่งจะทำให้ goroutine ใหม่มี call stack แยกต่างหาก และจะเป็น child ของ main goroutine เมื่อ main goroutine จบการทำงาน Go runtime จะทำการ terminate goroutines ลูกทั้งหมด
*/

func f1(name string) {
	for index := 0; index < 10; index++ {
		fmt.Printf("%v: index %d\n", name, index) // แสดงผลค่า name และ index
		time.Sleep(1 * time.Second)               // หน่วงเวลา goroutine เป็นเวลา 1 วินาที
	}
}

func main() {

	/*
		จะเห็นว่ามีการเรียกฟังก์ชัน f1 สองครั้งโดยใช้ goroutines ทำให้มีการทำงานแบบ concurrent ซึ่งผลลัพธ์ที่ได้จะไม่สามารถคาดเดาได้ และอาจแตกต่างกันในแต่ละครั้งที่รันโปรแกรม
	*/

	go f1("F1") // เริ่มต้น goroutine ใหม่เพื่อเรียกฟังก์ชัน f1 โดยส่งค่า "F1"
	go f1("F2") // เริ่มต้น goroutine ใหม่เพื่อเรียกฟังก์ชัน f1 โดยส่งค่า "F2"

	fmt.Println("Sleeping for 5 second")
	time.Sleep(5 * time.Second) // หน่วงเวลา main goroutine เป็นเวลา 5 วินาที
	fmt.Println("Main completed")

}

/*

 */
