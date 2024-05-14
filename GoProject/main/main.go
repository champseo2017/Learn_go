package main

import (
	"fmt"
)

/*

 */

func main() {

	// ไม่สามารถเปรียบเทียบค่าของตัวแปรกับพอยน์เตอร์ได้โดยตรง จะเกิด compilation error
	a := 10
	b := &a
	fmt.Println(a == b) // compilation error: invalid operation: a == b (mismatched types int and *int)
}

/*

 */
