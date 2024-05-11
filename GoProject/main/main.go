package main

import (
	"fmt"
)

/*
code ที่แสดงการสร้าง slice จาก array และส่ง slice เข้าไปใน function แทนการส่ง address ของ array โดยตรง
*/

func modify(slice []int) { // function ที่รับ slice ของ int
	slice[0] = 100 // เปลี่ยนค่า element ที่ 0 ของ slice เป็น 100
}

func main() {
	arr := [5]int{10, 20, 30, 40, 50}   // ประกาศและกำหนดค่าเริ่มต้นให้กับ array ของ int ขนาด 5 ตัว
	slice := arr[:]                     // สร้าง slice จาก arr โดยใช้ index เริ่มต้นและสิ้นสุดเป็นค่าเริ่มต้น (0 ถึง length ของ array - 1)
	modify(slice)                       // ส่ง slice เข้าไปใน modify function
	fmt.Println("Array Elements:", arr) // แสดงค่า element ของ arr
}
