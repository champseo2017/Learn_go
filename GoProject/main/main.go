package main

import (
	"fmt"
)

/*

 */

func main() {
	// 1. ตัวแปรพอยน์เตอร์ a ที่ไม่ได้กำหนดค่าจะมีค่าเป็น nil
	var a *int
	fmt.Println(a == nil) // ผลลัพธ์เป็น true
}

/*

 */
