package main

import "fmt"

/*
ในภาษา Go เราสามารถเก็บฟังก์ชันไว้ในโครงสร้างข้อมูลได้ เช่น slice หรือ map
*/

func Storing_Functions_In_Slice() []func() {
	f1 := func() {
		fmt.Println("Raj") // ฟังก์ชัน f1 จะพิมพ์ข้อความ "Raj"
	}
	f2 := func() {
		fmt.Println("Shyam") // ฟังก์ชัน f2 จะพิมพ์ข้อความ "Shyam"
	}
	arr := []func(){f1, f2} // สร้าง slice ของฟังก์ชันชื่อ arr และเก็บฟังก์ชัน f1 และ f2 ไว้ในนั้น
	return arr              // คืนค่า slice ของฟังก์ชัน arr
}

func main() {
	arr := Storing_Functions_In_Slice() // เรียกใช้ฟังก์ชัน Storing_Functions_In_Slice และเก็บผลลัพธ์ไว้ในตัวแปร arr
	arr[0]()                            // เรียกใช้ฟังก์ชันตัวแรกใน slice arr (f1)
	arr[1]()                            // เรียกใช้ฟังก์ชันตัวที่สองใน slice arr (f2)
}

/*
1. เราประกาศฟังก์ชัน `Storing_Functions_In_Slice` ซึ่งมีการคืนค่าเป็น slice ของฟังก์ชัน (`[]func()`)
2. ภายในฟังก์ชัน `Storing_Functions_In_Slice` เราประกาศตัวแปร `f1` และ `f2` เป็นฟังก์ชันแบบ anonymous (ไม่ระบุชื่อ)
   - ฟังก์ชัน `f1` จะพิมพ์ข้อความ "Raj" และฟังก์ชัน `f2` จะพิมพ์ข้อความ "Shyam" ออกมาทาง console
3. เราสร้าง slice ของฟังก์ชันชื่อ `arr` และเก็บฟังก์ชัน `f1` และ `f2` ไว้ในนั้น
4. ฟังก์ชัน `Storing_Functions_In_Slice` คืนค่า slice ของฟังก์ชัน `arr` ออกมา
5. ในฟังก์ชัน `main` เราเรียกใช้ฟังก์ชัน `Storing_Functions_In_Slice` และเก็บผลลัพธ์ไว้ในตัวแปร `arr`
6. เราเรียกใช้ฟังก์ชันตัวแรก (`f1`) และตัวที่สอง (`f2`) ใน slice `arr` โดยใช้ index `arr[0]()` และ `arr[1]()`
*/
