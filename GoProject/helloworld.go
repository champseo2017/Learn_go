package main

import "fmt"

/*
ในการส่งอาร์กิวเมนต์แบบ call by value ค่าของอาร์กิวเมนต์จริงจะถูก copy ไปยังพารามิเตอร์ของฟังก์ชัน โดยอาร์กิวเมนต์จริงและพารามิเตอร์จะถูกเก็บไว้ในหน่วยความจำคนละที่ การเปลี่ยนแปลงค่าของพารามิเตอร์ภายในฟังก์ชันจะไม่ส่งผลต่ออาร์กิวเมนต์จริง

ถึงแม้ว่า Go จะไม่มีการส่งแบบ call by reference เพราะไม่มีตัวแปรแบบ reference แต่เราสามารถส่งที่อยู่ของตัวแปรเข้าไปในฟังก์ชันได้ โดยใช้ตัวดำเนินการ & และ * ร่วมกับ pointer ซึ่งจะทำให้พารามิเตอร์และอาร์กิวเมนต์จริงอ้างอิงไปยังหน่วยความจำตำแหน่งเดียวกัน ทำให้การเปลี่ยนแปลงค่าภายในฟังก์ชันมีผลต่ออาร์กิวเมนต์จริงด้วย
*/

// ตัวอย่างการส่งอาร์กิวเมนต์แบบ call by value

func replace(x int){
	x = 20
}

func main() {
	var x int = 10
	fmt.Printf("value of x before function call = %d", x)

	replace(x)
	fmt.Printf("\nvalue of x after function call = %d", x)
	// จะเห็นว่าค่าของ x ไม่เปลี่ยนแปลง เพราะการเปลี่ยนค่าใน replace ไม่มีผลต่อ x ใน main
}