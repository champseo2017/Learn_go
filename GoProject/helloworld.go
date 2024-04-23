package main

import "fmt"

/*
Go, multi-dimensional slice คือ slice ที่มีหลายมิติ คล้ายกับ multi-dimensional array แต่ slice ไม่ต้องระบุขนาด เราสามารถสร้างและเข้าถึง multi-dimensional slice ได้เหมือนกับ multi-dimensional array
*/

func main() {
	// สร้าง multi-dimensional slice s1
	s1 := [][]int {
		{10, 20},
		{40, 80},
		{50, 100},
		{200, 400},
	}
	// เข้าถึง multi-dimensional slice s1
	fmt.Println("Slice 1 :", s1)

	// สร้าง multi-dimensional slice s2
	s2 := [][]string {
		{"Rajasthan", "Jaipur"},
        {"Maharashtra", "Mumbai"},
       {"Karnataka", "Bengaluru"},
	}
	// เข้าถึง multi-dimensional slice s2
	fmt.Println("Slice 2: ", s2)
}

/* 
1. เราประกาศ package main เพื่อระบุว่านี่คือโปรแกรมหลัก

2. เรานำเข้า package "fmt" เพื่อใช้ในการแสดงผลลัพธ์

3. เรากำหนดฟังก์ชัน main() ซึ่งเป็นจุดเริ่มต้นของโปรแกรม

4. เราสร้าง multi-dimensional slice s1 ด้วยค่าเริ่มต้น {{10, 20}, {40, 80}, {50, 100}, {200, 400}}

5. เราใช้ fmt.Println("Slice 1 :", s1) เพื่อแสดงค่าของ slice s1

6. เราสร้าง multi-dimensional slice s2 ด้วยค่าเริ่มต้น {{"Rajasthan", "Jaipur"}, {"Maharashtra", "Mumbai"}, {"Karnataka", "Bengaluru"}}

7. เราใช้ fmt.Println("Slice 2 :", s2) เพื่อแสดงค่าของ slice s2

เมื่อรันโปรแกรมนี้ ผลลัพธ์ที่ได้จะเป็น:

Slice 1 : [[10 20] [40 80] [50 100] [200 400]]
Slice 2 : [[Rajasthan Jaipur] [Maharashtra Mumbai] [Karnataka Bengaluru]]
*/