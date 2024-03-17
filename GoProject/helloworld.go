package main

import "fmt"

// Operator

func main() {
	// ตัวอย่างตัวดำเนินการคณิตศาสตร์
	fmt.Println("5 + 3 =", 5+3) // บวก
	fmt.Println("5 - 3 =", 5-3) // ลบ
	fmt.Println("5 * 3 =", 5*3) // คูณ
	fmt.Println("5 / 3 =", 5/3) // หาร
	fmt.Println("5 % 3 =", 5%3) // หารเอาเศษ

	// ตัวอย่างตัวดำเนินการเปรียบเทียบ
	fmt.Println("5 == 3:", 5 == 3) // เท่ากับ
	fmt.Println("5 != 3:", 5 != 3) // ไม่เท่ากับ
	fmt.Println("5 < 3:", 5 < 3)   // น้อยกว่า
	fmt.Println("5 > 3:", 5 > 3)   // มากกว่า

	// ตัวอย่างตัวดำเนินการตรรกะ
	fmt.Println("true && false =", true && false) // และ
	fmt.Println("true || false =", true || false) // หรือ
	fmt.Println("!true =", !true)                 // ไม่

	// ตัวอย่างตัวดำเนินการบิตไวส์
	fmt.Println("5 & 3 =", 5&3)   // AND
	fmt.Println("5 | 3 =", 5|3)   // OR
	fmt.Println("5 ^ 3 =", 5^3)   // XOR
	fmt.Println("5 &^ 3 =", 5&^3) // AND NOT

	// ตัวอย่างตัวดำเนินการกำหนดค่า
	var a int = 5
	a += 3 // a = a + 3
	fmt.Println("a += 3:", a)
}
