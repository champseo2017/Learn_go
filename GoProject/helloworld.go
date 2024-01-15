package main

import (
	"fmt"
)

/*
 ประกาศตัวแปร name ด้วย var ไว้ภายในฟังก์ชั่น main และประกาศ
 ตัวแปร age ด้วย var ไว้ภายนอกฟังก์ชัน main
*/

var age int = 18

func main() {
	var name string = "Orapin"
	fmt.Println("Hello World", name)
	fmt.Println("Your age is ", age)
}
