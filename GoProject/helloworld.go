package main

import "fmt"

/*
Go, เราสามารถใช้ตัวดำเนินการ == เพื่อตรวจสอบว่า slice เป็น nil หรือไม่ โดยการเปรียบเทียบ slice กับ nil ถ้า slice เป็น nil ผลลัพธ์จะเป็น true แต่ถ้า slice ไม่เป็น nil ผลลัพธ์จะเป็น false
*/

func main() {
	// สร้าง slice s1
	s1 := []int{1, 2, 3, 4, 5}
	// สร้าง slice s2 เป็น nil slice
	var s2 []int
	// ตรวจสอบว่า slice เป็น nil หรือไม่
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
}

/* 
1. เราประกาศ package main เพื่อระบุว่านี่คือโปรแกรมหลัก

2. เรานำเข้า package "fmt" เพื่อใช้ในการแสดงผลลัพธ์

3. เรากำหนดฟังก์ชัน main() ซึ่งเป็นจุดเริ่มต้นของโปรแกรม

4. เราสร้าง slice s1 ด้วยค่าเริ่มต้น [1, 2, 3, 4, 5] โดยใช้ s1 := []int{1, 2, 3, 4, 5}

5. เราสร้าง slice s2 เป็น nil slice โดยใช้ var s2 []int

6. เราใช้ fmt.Println(s1 == nil) เพื่อตรวจสอบว่า slice s1 เป็น nil หรือไม่ และแสดงผลลัพธ์

7. เราใช้ fmt.Println(s2 == nil) เพื่อตรวจสอบว่า slice s2 เป็น nil หรือไม่ และแสดงผลลัพธ์
*/