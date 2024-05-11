package main

import (
	"fmt"
)

/*
ในภาษา Go เราสามารถกำหนดค่าเริ่มต้นให้กับอาร์เรย์ได้ตั้งแต่ตอนประกาศอาร์เรย์ โดยใช้การประกาศอาร์เรย์แบบ array literal ซึ่งเป็นการระบุค่าเริ่มต้นของอาร์เรย์ไว้ในวงเล็บปีกกาหลังชนิดข้อมูลของอาร์เรย์
*/

func Arrays_Initialize_Example() {
	// ประกาศอาร์เรย์ a ขนาด 5 ช่อง ชนิด int และกำหนดค่าเริ่มต้นเป็น [1, 2, 3, 4, 5]
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	// แสดงค่าของอาร์เรย์ a
	fmt.Println(a)
}

func main() {
	// เรียกใช้ฟังก์ชัน Arrays_Initialize_Example
	Arrays_Initialize_Example()
}
