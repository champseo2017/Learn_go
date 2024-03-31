package main

import "fmt"

// การประกาศฟังก์ชัน
func add(x int, y int) int {
	return x + y
}

// การคืนค่าหลายค่า
func swap(x, y string) (string, string) {
	return y, x
}

// การคืนค่าที่มีชื่อ
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// เรียกใช้ฟังก์ชัน add ด้วยการส่งค่า 42 และ 13 
    // และแสดงผลลัพธ์ด้วย fmt.Println
	fmt.Println(add(42, 13))

	 // เรียกใช้ฟังก์ชัน swap ด้วยการส่งค่า "hello" และ "world"
    // และกำหนดผลลัพธ์ให้กับตัวแปร a และ b
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// เรียกใช้ฟังก์ชัน split ด้วยการส่งค่า 17
    // และแสดงผลลัพธ์ทั้งสองค่าที่ได้จากฟังก์ชัน split ด้วย fmt.Println
	x, y := split(17)
	fmt.Println(x, y)
}
