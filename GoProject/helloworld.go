package main

import "fmt"

/*
ภาษา Go เราสามารถใช้ range ใน for loop เพื่อวนลูปและเข้าถึงทั้งตำแหน่งและค่าขององค์ประกอบแต่ละตัวใน slice ได้ โดยในแต่ละรอบของลูป ตัวแปรแรก (เช่น index) จะเก็บตำแหน่งขององค์ประกอบ และตัวแปรที่สอง (เช่น element) จะเก็บค่าขององค์ประกอบ
*/

func main() {
	// สร้าง slice
	my_slice := []string{"Golang", "program", "to", "iterate", "over", "slice", "using", "range", "in", "for", "loop"}
	// วนลูปโดยใช้ range ใน for loop
	for index, element := range my_slice {
		fmt.Printf("Index = %d and Element = %s\n", index, element)
	}
	/* 
	1. เราประกาศ package main เพื่อระบุว่านี่คือโปรแกรมหลัก

	2. เรานำเข้า package "fmt" เพื่อใช้ในการแสดงผลลัพธ์

	3. เรากำหนดฟังก์ชัน main() ซึ่งเป็นจุดเริ่มต้นของโปรแกรม

	4. เราสร้าง slice ชื่อ my_slice ด้วยค่าเริ่มต้น {"Golang", "program", "to", "iterate", "over", "slice", "using", "range", "in", "for", "loop"}

	5. เราใช้ range ใน for loop เพื่อวนลูปและเข้าถึงทั้งตำแหน่งและค่าขององค์ประกอบแต่ละตัวใน my_slice
	- ในแต่ละรอบของลูป ตัวแปร index จะเก็บตำแหน่งขององค์ประกอบ และตัวแปร element จะเก็บค่าขององค์ประกอบ
	- เราใช้ fmt.Printf() เพื่อแสดงค่าของ index และ element ในรูปแบบ "Index = %d and Element = %s\n"

	6. ลูปจะทำงานต่อไปเรื่อยๆ จนกว่าจะครบทุกองค์ประกอบใน my_slice
	
	*/
}