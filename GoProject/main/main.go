package main

import "fmt"

/*
เราสามารถส่งองค์ประกอบที่อยู่ใน slice เข้าไปในฟังก์ชันแบบ variadic ได้ โดยใช้เทคนิคที่เรียกว่า "Slice Destructuring"
*/

func Variable_Arguments(s ...int) {
	fmt.Println(s)                // พิมพ์ค่าของอาร์กิวเมนต์ที่รับเข้ามาทั้งหมด
	fmt.Println(s[0], s[1], s[2]) // พิมพ์ค่าของอาร์กิวเมนต์ตัวที่ 1, 2, และ 3
	fmt.Println(len(s))           // พิมพ์จำนวนอาร์กิวเมนต์ที่รับเข้ามา
}

func main() {
	d := []int{10, 20, 30, 40, 50, 60} // สร้าง slice ชื่อ d ที่เก็บค่า int
	Variable_Arguments(d...)           // เรียกใช้ฟังก์ชัน Variable_Arguments โดยส่ง slice d เข้าไปแบบ destructuring
}

/*
1. เราประกาศฟังก์ชัน `Variable_Arguments` ที่รับอาร์กิวเมนต์แบบ variadic โดยใช้ `...int` เป็นพารามิเตอร์ (เหมือนตัวอย่างก่อนหน้า)
2. ในฟังก์ชัน `main` เราสร้าง slice ชื่อ `d` ที่เก็บค่า int โดยมีค่าเป็น 10, 20, 30, 40, 50, และ 60
3. เราเรียกใช้ฟังก์ชัน `Variable_Arguments` โดยส่ง slice `d` เข้าไป แต่เราใช้ `...` ต่อท้าย `d` เพื่อทำ slice destructuring
   - การใช้ `d...` จะเป็นการแยกองค์ประกอบใน slice `d` ออกเป็นอาร์กิวเมนต์แยกกัน และส่งเข้าไปในฟังก์ชัน `Variable_Arguments`
   - ดังนั้น `Variable_Arguments(d...)` จะเทียบเท่ากับ `Variable_Arguments(10, 20, 30, 40, 50, 60)`
*/
