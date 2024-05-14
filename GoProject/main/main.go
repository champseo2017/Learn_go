package main

import (
	"fmt"
)

/*

 */

func main() {
	// 2. ตัวแปรพอยน์เตอร์ a และ b ที่ไม่ได้กำหนดค่าจะมีค่าเป็น nil
	var a *int
	var b *int
	fmt.Println(a == b) // ผลลัพธ์เป็น true
}

/*

 */
