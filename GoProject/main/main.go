package main

import (
	"fmt"
)

/*
ตัวอย่างการกำหนด Method ให้กับ Type ต่างๆ
*/

// Method ของ Function
type MathFunc func(int, int) int

func (f MathFunc) Compute(a, b int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}

func main() {
	mathFunc := MathFunc(add)
	result := mathFunc.Compute(3, 5)
	fmt.Println(result) // Output: 8
}
/* 
จะเห็นได้ว่าเราสามารถกำหนด Method ให้กับ Type ต่างๆ ได้หลากหลาย ทั้ง Basic Types อย่าง int และ string, Composite Types เช่น Array, Slice และ Map รวมถึง Function เพื่อเพิ่มความสามารถและพฤติกรรมเฉพาะให้กับ Type เหล่านั้น การเรียกใช้ Method สามารถทำได้โดยใช้ตัวแปรหรือค่าของ Type ที่กำหนดและตามด้วยชื่อ Method และพารามิเตอร์ (ถ้ามี) ผลลัพธ์ที่ได้จะขึ้นอยู่กับการดำเนินการภายใน Method ที่กำหนดไว้
*/