package main

import (
	"fmt"
)

// While Loop
func For_Loop_Example_Form2(i int) {
	j := 0 // กำหนดค่าเริ่มต้นให้กับตัวแปร j เป็น 0

	// ใช้ For Loop เหมือน While Loop โดยมีเพียงส่วนเงื่อนไข (j < i)
	for j < i {
		fmt.Println("The value of j in this iteration is", j)
		j++ // เพิ่มค่าของ j ขึ้นทีละ 1 ในแต่ละรอบของการวนลูป
	}
}

// Infinite Loop
func For_Loop_Example_Form3(i int) {
	for { // ใช้ For Loop โดยไม่ระบุเงื่อนไข เพื่อสร้าง Infinite Loop
		fmt.Println("The value of j in this iteration is", i)
	}
}

func main() {
	// For_Loop_Example_Form2(10)
	// For_Loop_Example_Form3(1)
}
