package main

import (
	"fmt"
)

/*
ในภาษา Go เราสามารถประกาศอาร์เรย์โดยใช้วิธี pre-initialization ซึ่งเป็นการกำหนดค่าเริ่มต้นให้กับอาร์เรย์โดยไม่ต้องระบุขนาดของอาร์เรย์ แต่ใช้จุดสามแทน (...) โดยขนาดของอาร์เรย์จะถูกกำหนดโดยอัตโนมัติตามจำนวนค่าที่กำหนดใน array literal
*/

func Arrays_PreInitializeExample() {
	// ประกาศอาร์เรย์ x โดยใช้ pre-initialization และกำหนดค่าเริ่มต้นเป็น [23, 4, 32, 5, 76]
	x := [...]int{23, 4, 32, 5, 76}
	// แสดงค่าของอาร์เรย์ x
	fmt.Println(x)
}

func main() {
	// เรียกใช้ฟังก์ชัน Arrays_PreInitializeExample
	Arrays_PreInitializeExample()
}
