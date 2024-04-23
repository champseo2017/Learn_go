package main

import (
	"fmt"
)

/*
Go เราสามารถวนลูปเพื่อเข้าถึงองค์ประกอบใน slice ได้ทีละตัว โดยใช้ for loop แบบปกติ หรือใช้ for loop ร่วมกับ range keyword
*/

func main() {
	// ตัวอย่างการวนลูปเข้าถึงองค์ประกอบใน slice ด้วย for loop แบบปกติ:
	// สร้าง slice1
	slice1 := []int{10, 20, 30, 40, 50}
	// วนลูปเข้าถึงองค์ประกอบใน slice ด้วย for loop
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice1[%d] = %d\n", i, slice1[i])
	}
	/* 
		1. เราสร้าง slice ชื่อ slice ด้วยค่าเริ่มต้น {10, 20, 30, 40, 50}
		2. เราใช้ for loop เพื่อวนลูปเข้าถึงองค์ประกอบใน slice โดยใช้ตัวแปร i เป็นตัวนับ เริ่มจาก 0 ถึง len(slice) - 1
		3. ในแต่ละรอบของลูป เราใช้ fmt.Printf() เพื่อแสดงค่าของ slice[i] พร้อมกับตำแหน่งของมัน
	*/
	// สร้าง slice2
	slice2 := []string{"apple", "banana", "cherry", "durian"}
	// วนลูปเข้าถึงองค์ประกอบใน slice2 ด้วย for loop ร่วมกับ range
	for index, value := range slice2 {
		fmt.Printf("slice[%d] = %s\n", index, value)
	}
	/* 
		1. เราสร้าง slice ชื่อ slice ด้วยค่าเริ่มต้น {"apple", "banana", "cherry", "durian"}
		2. เราใช้ for loop ร่วมกับ range เพื่อวนลูปเข้าถึงองค์ประกอบใน slice โดยในแต่ละรอบของลูป ตัวแปร index จะเก็บตำแหน่งขององค์ประกอบ และตัวแปร value จะเก็บค่าขององค์ประกอบ
		3. ในแต่ละรอบของลูป เราใช้ fmt.Printf() เพื่อแสดงค่าของ value พร้อมกับตำแหน่ง index ของมัน

		การใช้ for loop ร่วมกับ range จะทำให้โค้ดอ่านง่ายและกระชับกว่าการใช้ for loop แบบปกติ เพราะเราไม่ต้องกังวลเรื่องตัวนับและการเข้าถึงองค์ประกอบผ่านตำแหน่ง
	*/
}

