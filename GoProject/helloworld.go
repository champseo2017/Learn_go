package main

import "fmt"

/*
Slice ใน Go เป็น reference type และอ้างอิงไปยัง array ที่อยู่เบื้องหลัง ดังนั้น การแก้ไของค์ประกอบของ slice จะส่งผลต่อการแก้ไของค์ประกอบที่สอดคล้องกันใน array ที่ถูกอ้างอิง และ slice อื่นๆ ที่อ้างอิงไปยัง array เดียวกันก็จะเห็นการแก้ไขนั้นด้วยเช่นกัน
*/

func main() {
	a := [7]string{"One", "Two", "Three", "Four", "Five", "Six", "Seven"}
	slice1 := a[1:]
	slice2 := a[3:]

	fmt.Println("***** Before Modifications *****")
	fmt.Println("Array =", a)
	fmt.Println("Slice1 =", slice1)
	fmt.Println("Slice2 =", slice2)

	slice1[0] = "TWO"
	slice1[2] = "FOUR"
	slice2[1] = "FIVE"

	fmt.Println("\n***** After Modifications *****")
	fmt.Println("Array =", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
}

/* 
จากผลลัพธ์ จะเห็นได้ว่าการแก้ไขค่าใน slice1 และ slice2 ส่งผลต่อการแก้ไขค่าใน array a ด้วย เนื่องจาก slice อ้างอิงไปยัง array ที่อยู่เบื้องหลัง และเมื่อมีการแก้ไขค่าใน slice ค่าใน array ที่ถูกอ้างอิงก็จะถูกแก้ไขด้วย นอกจากนี้ slice อื่นๆ ที่อ้างอิงไปยัง array เดียวกันก็จะเห็นการแก้ไขนั้นเช่นกัน
*/