package main

import "fmt"

/*
ในภาษา Go, Zero-value slice หรือ Nil slice คือ slice ที่ไม่มีองค์ประกอบใดๆ อยู่ภายใน ดังนั้น ความยาว (length) และความจุ (capacity) ของ slice นี้จึงเป็น 0 นอกจากนี้ Nil slice ยังไม่มีการอ้างอิงถึง array ใดๆ อีกด้วย
*/

func main() {
	// สร้าง zero value slice
	var my_slice []string
	// แสดงความยาวของ slice
	fmt.Printf("Length of Slice = %d\n", len(my_slice))
	// แสดงความจุของ slice
	fmt.Printf("Capacity of Slice = %d", cap(my_slice))
}

/* 
1. เราประกาศ package main เพื่อระบุว่านี่คือโปรแกรมหลัก

2. เรานำเข้า package "fmt" เพื่อใช้ในการแสดงผลลัพธ์

3. เรากำหนดฟังก์ชัน main() ซึ่งเป็นจุดเริ่มต้นของโปรแกรม

4. เราประกาศตัวแปร my_slice เป็น zero value slice ชนิด string โดยใช้ var my_slice []string

5. เราใช้ฟังก์ชัน len() เพื่อแสดงความยาวของ slice และใช้ fmt.Printf() เพื่อแสดงผลลัพธ์

6. เราใช้ฟังก์ชัน cap() เพื่อแสดงความจุของ slice และใช้ fmt.Printf() เพื่อแสดงผลลัพธ์
*/