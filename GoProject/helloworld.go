package main

import "fmt"

/*
	การเริ่มต้น map โดยใช้ function make() เป็นอีกวิธีหนึ่งในการสร้าง map ในภาษา Go เราเพียงแค่ส่งประเภทของ map ไปยัง make() และ function นี้จะคืน map ที่ถูกเริ่มต้นและพร้อมใช้งานกลับมา
*/

func main() {
	var m1 = make(map[string]int)
	fmt.Println(m1)

	if m1 == nil {
		fmt.Println("Map is empty")
	} else {
		fmt.Println("Map is not empty")
	}
	// make() function returns an initialized and ready to use map
    // you can add new keys
	m1["ten"] = 10
	fmt.Println(m1)
}
/* 
- ประกาศตัวแปร `m1` และใช้ `make()` เพื่อสร้าง map ที่มี key เป็น string และ value เป็น int
- แสดงค่าของ `m1` ซึ่งจะเป็น empty map (`map[]`)
- ตรวจสอบว่า `m1` เป็น `nil` หรือไม่ ในกรณีนี้ `m1` ไม่เป็น `nil` เพราะถูกสร้างด้วย `make()` แล้ว ดังนั้นจะข้ามไปทำคำสั่งใน `else` และแสดงข้อความ "Map is not empty"
- เพิ่มคู่ key-value ใหม่ลงใน `m1` โดยกำหนดให้ key "ten" มีค่าเป็น 10 (แสดงให้เห็นว่า map ที่สร้างด้วย `make()` พร้อมใช้งานทันที)
- แสดงค่าของ `m1` อีกครั้ง ซึ่งตอนนี้จะมีคู่ key-value ที่เพิ่มเข้าไป

จะเห็นได้ว่าตอนแรก `m1` เป็น empty map ที่ถูกสร้างด้วย `make()` จากนั้นเราสามารถเพิ่มคู่ key-value ลงไปได้ทันที และ `m1` ก็มีข้อมูลตามที่เพิ่มเข้าไป แสดงให้เห็นว่า map ที่สร้างด้วย `make()` นั้นถูกเริ่มต้นและพร้อมใช้งานทันที
*/