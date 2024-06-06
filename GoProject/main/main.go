package main

import "fmt"

/*
Go มีคำสั่ง `defer` ที่ใช้สำหรับเลื่อนการทำงานของโค้ดบางส่วนไปทำหลังจากที่ฟังก์ชันทำงานเสร็จ
*/

func Defer_demo() {
	fmt.Println("this is 1 statement")
	defer fmt.Println("this is 2 statement") // เลื่อนการทำงานไปทำหลังจากฟังก์ชันเสร็จสิ้น
	defer fmt.Println("this is 3 statement") // เลื่อนการทำงานไปทำหลังจากฟังก์ชันเสร็จสิ้น
	fmt.Println("this is 4 statement")
	fmt.Println("this is 5 statement")
}

func main() {
	defer fmt.Println("Defer statement") // เลื่อนการทำงานไปทำหลังจากฟังก์ชัน main เสร็จสิ้น
	Defer_demo()
}

/*
1. ในฟังก์ชัน `Defer_demo` มีการใช้คำสั่ง `defer` กับ statement ที่ 2 และ 3
   - `defer fmt.Println("this is 2 statement")` จะเลื่อนการทำงานของ statement นี้ไปทำหลังจากที่ฟังก์ชัน `Defer_demo` ทำงานเสร็จแล้ว
   - `defer fmt.Println("this is 3 statement")` ก็เช่นเดียวกัน จะเลื่อนการทำงานไปทำหลังจากฟังก์ชันเสร็จสิ้น
2. ในฟังก์ชัน `main` มีการใช้คำสั่ง `defer` กับ statement `fmt.Println("Defer statement")`
   - statement นี้จะถูกเลื่อนการทำงานไปทำหลังจากที่ฟังก์ชัน `main` ทำงานเสร็จแล้ว
3. เมื่อรันโค้ด ลำดับการทำงานจะเป็นดังนี้
   - ฟังก์ชัน `main` เริ่มทำงาน
   - เจอคำสั่ง `defer` ใน `main` แต่ยังไม่ทำงานตอนนี้
   - เรียกใช้ฟังก์ชัน `Defer_demo`
   - ฟังก์ชัน `Defer_demo` ทำงาน โดยพิมพ์ "this is 1 statement", "this is 4 statement", "this is 5 statement" ตามลำดับ
   - เจอคำสั่ง `defer` ใน `Defer_demo` แต่ยังไม่ทำงานตอนนี้
   - ฟังก์ชัน `Defer_demo` ทำงานเสร็จ เริ่มทำงาน `defer` ใน `Defer_demo` โดยพิมพ์ "this is 3 statement" และ "this is 2 statement" ตามลำดับ (สังเกตว่าทำงานแบบ LIFO - Last In, First Out)
   - กลับมาทำงานใน `main` ต่อ
   - ฟังก์ชัน `main` ทำงานเสร็จ เริ่มทำงาน `defer` ใน `main` โดยพิมพ์ "Defer statement"
*/
