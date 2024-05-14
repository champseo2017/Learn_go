package main

import (
	"fmt"
)

/*
การสร้าง slice ในภาษา Go และอธิบายว่า slice ถูกสร้างขึ้นมาพร้อมกับ underlying array ที่มีขนาดถูกกำหนดตอนรันไทม์ นอกจากนี้ยังแสดงการสร้าง slice ใหม่จาก slice ที่มีอยู่แล้ว
*/

func Slice_from_Another_Slice_With_More_Size() {
	// สร้าง slice z ที่มีสมาชิกเป็นตัวเลขจำนวนเต็ม
	z := []int{10, 32, 42, 23, 12, 41, 5}

	// สร้าง slice x จาก slice z โดยระบุ index เริ่มต้นและสิ้นสุด
	// ในที่นี้ x จะมีขนาดมากกว่า z เพราะระบุ index สิ้นสุดเกินขนาดของ z
	x := z[0:7]

	// เปลี่ยนค่าสมาชิกตัวแรกของ slice x เป็น 30000
	x[0] = 30000

	// แสดงค่าของ slice z
	fmt.Println(z)
}

func main() {
	Slice_from_Another_Slice_With_More_Size()
}

/*
เมื่อเปลี่ยนค่าสมาชิกใน x เช่น x[0] = 30000 มันจะส่งผลต่อค่าใน z ด้วย เพราะทั้งสอง slice ใช้ underlying array ร่วมกัน
*/
