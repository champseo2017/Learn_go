package main

import (
	"fmt"
)

/*
เมื่อเราส่ง pointer ไปยัง struct เข้าไปในฟังก์ชัน และมีการเปลี่ยนแปลงค่าของ pointer นั้นภายในฟังก์ชันให้ชี้ไปยัง struct ใหม่ การเปลี่ยนแปลงนี้จะไม่ส่งผลต่อ pointer ตัวเดิมที่อยู่นอกฟังก์ชัน เนื่องจาก pointer ในฟังก์ชันเป็นเพียงสำเนาของ pointer ที่ส่งเข้ามา
ดังนั้น ใน main function เมื่อเรา print ค่าของ struct ที่ pointer p1 ชี้ไป หลังจากเรียกใช้ฟังก์ชัน modify ผลลัพธ์จะยังคงเป็นค่าเดิมของ struct ที่เรากำหนดให้กับ p1 ตั้งแต่แรก เพราะการเปลี่ยนแปลงค่าของ p1 ในฟังก์ชัน modify ไม่ได้ส่งผลต่อ p1 ใน main
*/

type person struct {
	id   int
	name string
}

func modify(p1 *person) {
	p1 = &person{202, "ABC"} // การเปลี่ยนแปลงนี้ไม่ส่งผลต่อ p1 ใน main
}

func main() {
	p1 := &person{101, "XYX"}
	modify(p1)
	fmt.Println("Person p1:", *p1) // ผลลัพธ์: {101 XYX}
}

/*

 */
