package main

import "fmt"

// Relational operators

func main() {
	var a int = 5
	var b int = 10

	fmt.Println("a == b:", a == b) // เท่ากับ
	fmt.Println("a != b:", a != b) // ไม่เท่ากับ
	fmt.Println("a > b:", a > b)   // มากกว่า
	fmt.Println("a < b:", a < b)   // น้อยกว่า
	fmt.Println("a >= b:", a >= b) // มากกว่าหรือเท่ากับ
	fmt.Println("a <= b:", a <= b) // น้อยกว่าหรือเท่ากับ
}
