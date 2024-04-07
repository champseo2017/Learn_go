package main

import "fmt"

/*
Golang รองรับการส่งอาร์กิวเมนต์เข้าฟังก์ชันแบบ call by value เท่านั้น ไม่มีแบบ call by reference เนื่องจากใน Golang ไม่มีตัวแปรแบบ reference อย่างไรก็ตาม เรายังสามารถส่งที่อยู่ของตัวแปร (memory address) เข้าไปในฟังก์ชันได้ ซึ่งโดยปกติแล้ว Golang จะใช้วิธีการ call by value ในการส่งอาร์กิวเมนต์เข้าฟังก์ชัน

จะเห็นว่าการเรียกฟังก์ชัน add โดยส่งค่า x และ y เข้าไปนั้น เป็นการ copy ค่า 10 และ 20 ไปใส่ในพารามิเตอร์ a และ b ดังนั้นถึงแม้ในฟังก์ชันเราจะเปลี่ยนแปลงค่า a และ b ก็จะไม่กระทบค่าของ x และ y ในฟังก์ชัน main
*/

// ตัวอย่างการใช้ call by value ใน Golang
func add(a, b int) int {
	return a + b
}

func main() {
	x, y := 10, 20
	result := add(x, y)
	fmt.Println("Result:", result)
}