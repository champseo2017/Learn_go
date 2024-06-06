package main

import "fmt"

/*
ตัวอย่างการเก็บฟังก์ชันไว้ใน map ในภาษา Go
*/

func Storing_Functions_In_Map() map[string]func() {
	f1 := func() {
		fmt.Println("Raj") // ฟังก์ชัน f1 จะพิมพ์ข้อความ "Raj"
	}
	f2 := func() {
		fmt.Println("Shyam") // ฟังก์ชัน f2 จะพิมพ์ข้อความ "Shyam"
	}
	funcMap := map[string]func(){
		"Raj":   f1, // เก็บฟังก์ชัน f1 ไว้ในคีย์ "Raj"
		"Shyam": f2, // เก็บฟังก์ชัน f2 ไว้ในคีย์ "Shyam"
	}
	return funcMap // คืนค่า map ของฟังก์ชัน funcMap
}

func main() {
	funcMap := Storing_Functions_In_Map() // เรียกใช้ฟังก์ชัน Storing_Functions_In_Map และเก็บผลลัพธ์ไว้ในตัวแปร funcMap
	funcMap["Raj"]()                      // เรียกใช้ฟังก์ชันที่เก็บไว้ในคีย์ "Raj" (f1)
	funcMap["Shyam"]()                    // เรียกใช้ฟังก์ชันที่เก็บไว้ในคีย์ "Shyam" (f2)
}

/*
1. เราประกาศฟังก์ชัน `Storing_Functions_In_Map` ซึ่งมีการคืนค่าเป็น map ของฟังก์ชัน (`map[string]func()`)
2. ภายในฟังก์ชัน `Storing_Functions_In_Map` เราประกาศตัวแปร `f1` และ `f2` เป็นฟังก์ชันแบบ anonymous (ไม่ระบุชื่อ)
   - ฟังก์ชัน `f1` จะพิมพ์ข้อความ "Raj" และฟังก์ชัน `f2` จะพิมพ์ข้อความ "Shyam" ออกมาทาง console
3. เราสร้าง map ของฟังก์ชันชื่อ `funcMap` และเก็บฟังก์ชัน `f1` และ `f2` ไว้ในคีย์ "Raj" และ "Shyam" ตามลำดับ
4. ฟังก์ชัน `Storing_Functions_In_Map` คืนค่า map ของฟังก์ชัน `funcMap` ออกมา
5. ในฟังก์ชัน `main` เราเรียกใช้ฟังก์ชัน `Storing_Functions_In_Map` และเก็บผลลัพธ์ไว้ในตัวแปร `funcMap`
6. เราเรียกใช้ฟังก์ชันที่เก็บไว้ในคีย์ "Raj" (`f1`) และ "Shyam" (`f2`) ใน map `funcMap` โดยใช้ `funcMap["Raj"]()` และ `funcMap["Shyam"]()`
*/
