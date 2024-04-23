package main

import (
	"fmt"
	"sort"
)

/*
ในภาษา Go เราสามารถเรียงลำดับองค์ประกอบใน slice ได้ โดยใช้ package sort ที่มีอยู่ในไลบรารีมาตรฐานของ Go ซึ่งมีเมธอดสำหรับการเรียงลำดับ slice ของ int, float64 และ string
*/

func main() {
	// สร้าง Slice
	slice1 := []string{"India", "Japan", "China", "Russia", "Singapore"}
	slice2 := []int{200, 500, 700, 400, 800, 300, 600, 900}

	fmt.Println("***** Before sorting *****")
	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	// เรียงลำดับ slice
	sort.Strings(slice1)
	sort.Ints(slice2)

	fmt.Println("\n***** After sorting *****")
    fmt.Println("Slice 1:", slice1)
    fmt.Println("Slice 2:", slice2)
}

/* 
1. เราประกาศ package main เพื่อระบุว่านี่คือโปรแกรมหลัก

2. เรานำเข้า package "fmt" เพื่อใช้ในการแสดงผลลัพธ์ และ package "sort" เพื่อใช้ในการเรียงลำดับ slice

3. เรากำหนดฟังก์ชัน main() ซึ่งเป็นจุดเริ่มต้นของโปรแกรม

4. เราสร้าง slice1 ด้วยค่าเริ่มต้น {"India", "Japan", "China", "Russia", "Singapore"} และ slice2 ด้วยค่าเริ่มต้น {200, 500, 700, 400, 800, 300, 600, 900}

5. เราใช้ fmt.Println() เพื่อแสดงค่าของ slice1 และ slice2 ก่อนการเรียงลำดับ

6. เราใช้ฟังก์ชัน sort.Strings(slice1) เพื่อเรียงลำดับ slice1 ตามลำดับตัวอักษร และใช้ sort.Ints(slice2) เพื่อเรียงลำดับ slice2 จากน้อยไปมาก

7. เราใช้ fmt.Println() เพื่อแสดงค่าของ slice1 และ slice2 หลังการเรียงลำดับ

เมื่อรันโปรแกรมนี้ ผลลัพธ์ที่ได้จะเป็น:

***** Before sorting *****
Slice 1: [India Japan China Russia Singapore]
Slice 2: [200 500 700 400 800 300 600 900]

***** After sorting *****
Slice 1: [China India Japan Russia Singapore]
Slice 2: [200 300 400 500 600 700 800 900]

ซึ่งแสดงให้เห็นว่า slice1 ถูกเรียงลำดับตามตัวอักษร และ slice2 ถูกเรียงลำดับจากน้อยไปมาก หลังจากใช้ฟังก์ชันเรียงลำดับจาก package sort
*/