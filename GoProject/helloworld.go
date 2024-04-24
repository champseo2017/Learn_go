package main

import "fmt"

/*
Go, Slice เป็นส่วนหนึ่งของ Array ที่สามารถปรับขนาดได้และมีความยืดหยุ่นมากกว่า Array Slice เป็นโครงสร้างข้อมูลเบาที่ห่อหุ้มและแสดงถึงส่วนหนึ่งของ Array โดย Slice ประกอบด้วย 3 สิ่ง ได้แก่ pointer, length และ capacity

ในการสร้าง Slice จาก Array เราสามารถระบุ index ต่ำสุดและสูงสุดโดยคั่นด้วยเครื่องหมายโคลอน (:) ค่าเริ่มต้นของ index ต่ำสุดคือ 0 และสูงสุดคือความยาวของ Slice

นอกจากนี้ เรายังสามารถสร้าง Slice โดยใช้ฟังก์ชัน make() ที่มีให้ใน Go library ซึ่งรับพารามิเตอร์ 3 ตัว ได้แก่ type, length และ capacity (capacity เป็นตัวเลือก)

Slice เป็น reference type ใน Go และอ้างอิงไปยัง underlying array ดังนั้น การแก้ไของค์ประกอบใน Slice จะแก้ไของค์ประกอบที่ตรงกันใน referenced array และ Slice อื่นๆ ที่อ้างอิงไปยัง array เดียวกันก็จะเห็นการเปลี่ยนแปลงนั้นด้วย

Nil Slice เป็น Slice ที่ไม่มีองค์ประกอบใดๆ ดังนั้น capacity และ length ของ Nil Slice จึงเป็น 0 และไม่มีการอ้างอิงไปยัง Array ใดๆ เราสามารถใช้ตัวดำเนินการ == เพื่อตรวจสอบว่า Slice เป็น Nil หรือไม่

Multi-dimensional Slice ใน Go เหมือนกับ Multi-dimensional Array ยกเว้นว่า Slice ไม่มีการระบุขนาด

เราสามารถวนลูปเข้าถึงองค์ประกอบใน Slice ได้โดยใช้ for loop หรือ range ใน for loop โดยการใช้ range ใน for loop จะทำให้เราได้ทั้ง index และค่าขององค์ประกอบด้วย

ในภาษา Go เราสามารถเรียงลำดับองค์ประกอบใน Slice ได้ โดย Go standard library มี package sort ที่มีเมธอดสำหรับการเรียงลำดับ Slice ของ int, float64, string และอื่นๆ
*/

func main() {
	// สร้าง Array
	arr := [5]int{1, 2, 3, 4, 5}
	// สร้าง Slice จาก Array
	slice1 := arr[1:4]
	fmt.Println("Slice 1:", slice1)
	/* 
	สร้าง Slice จาก Array เราสามารถระบุ index ต่ำสุดและสูงสุดโดยคั่นด้วยเครื่องหมายโคลอน (:) ค่าเริ่มต้นของ index ต่ำสุดคือ 0 และสูงสุดคือความยาวของ Slice
	*/
	// สร้าง Slice โดยใช้ฟังก์ชัน make()
	slice2 := make([]string, 3, 5)
	fmt.Println("Slice 2:", slice2)
	// การแก้ไของค์ประกอบใน Slice
	slice1[0] = 10
	fmt.Println("Array:", arr)
	fmt.Println("Slice 1:", slice1)
	// Nil Slice
	var slice3 []int
	fmt.Println("Slice 3 is nil:", slice3 == nil)
	// Multi-dimensional Slice
	multiSlice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Multi-dimensional Slice:", multiSlice)
	// วนลูปเข้าถึงองค์ประกอบใน Slice
	for i, v := range slice1 {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}

/* 
1. เราสร้าง Array `arr` ที่มี 5 องค์ประกอบ
2. เราสร้าง Slice `slice1` จาก Array `arr` โดยระบุ index ต่ำสุดเป็น 1 และสูงสุดเป็น 4 (ไม่รวม 4)
3. เราสร้าง Slice `slice2` โดยใช้ฟังก์ชัน `make()` โดยระบุ type เป็น string, length เป็น 3 และ capacity เป็น 5
4. เราแก้ไของค์ประกอบที่ index 0 ของ `slice1` เป็น 10 ซึ่งจะแก้ไขค่าใน `arr` ด้วย
5. เราประกาศ Slice `slice3` แบบ Nil และตรวจสอบว่าเป็น Nil หรือไม่โดยใช้ตัวดำเนินการ `==`
6. เราสร้าง Multi-dimensional Slice `multiSlice` ที่มี 3 แถวและแต่ละแถวมี 3 คอลัมน์
7. เราใช้ range ใน for loop เพื่อวนลูปและแสดงค่า index และค่าขององค์ประกอบใน `slice1`
*/