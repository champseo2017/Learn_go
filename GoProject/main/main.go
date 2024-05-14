package main

import (
	"fmt"
)

/*

 */

func main() {

	// 3. x และ y เป็นพอยน์เตอร์ที่ชี้ไปยังตำแหน่งความจำที่แตกต่างกัน
	a := 10
	b := 10
	x := &a
	y := &b
	fmt.Println(x == y) // ผลลัพธ์เป็น false
}

/*

 */
