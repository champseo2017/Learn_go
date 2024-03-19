package main

import "fmt"

/*
การกำหนดค่าใหม่ให้กับตัวแปรที่มีอยู่แล้วในบล็อกโค้ดเดียวกัน

*/

func main() {
	a := 10
	a, b := 20, 30 // การประกาศใหม่ที่ถูกต้อง เนื่องจากมีการแนะนำ b ซึ่งเป็นตัวแปรใหม่
	fmt.Println("value of a:", a)
	fmt.Println("value of b:", b)
}
