package main

import (
	"fmt"
)

// ประเภทฟังก์ชัน (Function types)
// ประกาศฟังก์ชัน add ที่รับพารามิเตอร์เป็น x และ y และคืนค่าเป็นผลบวกของ x และ y
func add(x int, y int) int {
    return x + y
}

func main() {
  // เรียกใช้ฟังก์ชัน add ด้วยค่า 5 และ 7
  result := add(5, 7)
  // แสดงผลลัพธ์
  fmt.Println("The result of adding 5 and 7 is:", result)
  /* 
    เราได้ประกาศฟังก์ชัน add ที่รับพารามิเตอร์สองตัวคือ x และ y ทั้งคู่เป็นประเภท int และคืนค่าเป็นผลบวกของ x และ y
    ในฟังก์ชัน main, เราเรียกใช้ฟังก์ชัน add โดยส่งค่า 5 และ 7 เข้าไป และเก็บผลลัพธ์ที่ได้ในตัวแปร result
    จากนั้นเราใช้ fmt.Println เพื่อแสดงผลลัพธ์ออกมา ซึ่งในกรณีนี้คือ "The result of adding 5 and 7 is: 12"
  */
}