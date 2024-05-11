package main

import (
	"fmt"
)

/*
ฟังก์ชัน Scan ใน Golang ช่วยในการอ่านข้อมูลจาก console และนำไปใช้ในโปรแกรมได้ โดยสามารถอ่านค่าจากผู้ใช้และเก็บไว้ในตัวแปรที่กำหนด
*/

func Scan_HelloMessage_Example() {
	var msg int
	fmt.Println("Enter a number")
	// อ่านค่าจากผู้ใช้และเก็บไว้ในตัวแปร msg
	// num คือจำนวนค่าที่อ่านได้, err คือ error ที่เกิดขึ้น (ถ้ามี)
	num, err := fmt.Scan(&msg)
	if err != nil {
		// ถ้ามี error ให้แสดงข้อความ error
		fmt.Println("the error is as follows", err)
	} else {
		// ถ้าไม่มี error ให้แสดงค่าที่อ่านได้ (msg) และจำนวนค่าที่อ่านได้ (num)
		fmt.Println(msg, num)
	}
}

/*
1. ฟังก์ชัน `Scan_HelloMessage_Example` ถูกเรียกใช้ในฟังก์ชัน `main`
2. กำหนดตัวแปร `msg` เป็นชนิด int เพื่อเก็บค่าที่อ่านจากผู้ใช้
3. แสดงข้อความ "Enter a number" เพื่อบอกให้ผู้ใช้ป้อนค่า
4. ใช้ฟังก์ชัน `fmt.Scan` เพื่ออ่านค่าจาก console และเก็บไว้ในตัวแปร `msg`
   - `&msg` หมายถึงการส่งตำแหน่งของตัวแปร `msg` เพื่อให้ฟังก์ชัน `Scan` สามารถเปลี่ยนแปลงค่าได้
   - `num` จะเก็บจำนวนค่าที่อ่านได้
   - `err` จะเก็บ error ที่เกิดขึ้น (ถ้ามี)
5. ตรวจสอบว่ามี error หรือไม่ โดยใช้ `if err != nil`
   - ถ้ามี error ให้แสดงข้อความ error ด้วย `fmt.Println("the error is as follows", err)`
   - ถ้าไม่มี error ให้แสดงค่าที่อ่านได้ (`msg`) และจำนวนค่าที่อ่านได้ (`num`) ด้วย `fmt.Println(msg, num)`
6. เมื่อฟังก์ชันทำงานเสร็จ โปรแกรมจะจบการทำงาน
*/

func main() {
	Scan_HelloMessage_Example()
}
