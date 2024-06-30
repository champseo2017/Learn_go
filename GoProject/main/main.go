package main

import (
	// oops "GoProject/package"
	"fmt"
)

/*
Channel ซึ่งเป็นวิธีการส่งข้อมูลระหว่างส่วนต่างๆ ของแอปพลิเคชัน ไม่ว่าจะเป็นระหว่าง Package หรือ Function
*/

func addNumberToChannel(numChannel chan int) {
	// กำหนดให้ numChannel มีค่าเท่ากับ 100
	numChannel <- 100
}

func main() {
	// สร้าง Channel ชื่อ numChannel ที่เก็บข้อมูลประเภท int
	numChannel := make(chan int)

	// เรียกใช้ฟังก์ชัน addNumberToChannel และส่ง numChannel เข้าไป
	go addNumberToChannel(numChannel)

	// อ่านค่าจาก numChannel มาเก็บไว้ในตัวแปร num
	num := <-numChannel

	// แสดงข้อความ "the number is printed from channel" ตามด้วยค่าของ num
	fmt.Println("the number is printed from channel", num)
}

/*
ใน `func main()` เราสร้าง Channel ชื่อ `numChannel` ที่ใช้เก็บข้อมูลประเภท `int` ด้วยคำสั่ง `make(chan int)`

จากนั้นเรียกใช้ฟังก์ชัน `addNumberToChannel` โดยส่ง `numChannel` เข้าไป และใช้คีย์เวิร์ด `go` เพื่อให้ฟังก์ชันทำงานแบบ Goroutine (ทำงานพร้อมกันโดยไม่ต้องรอให้ฟังก์ชันอื่นทำงานเสร็จก่อน)

ภายในฟังก์ชัน `addNumberToChannel` เรากำหนดให้ `numChannel` มีค่าเท่ากับ 100 ด้วยคำสั่ง `numChannel <- 100`

กลับมาที่ `func main()` อีกครั้ง หลังจากที่ `addNumberToChannel` ทำงานเสร็จ เราจะอ่านค่าจาก `numChannel` ด้วยคำสั่ง `num := <-numChannel` ซึ่งค่าที่อ่านได้จะถูกเก็บไว้ในตัวแปร `num`

สุดท้ายเราแสดงข้อความ "the number is printed from channel" ตามด้วยค่าของ `num` ด้วย `fmt.Println()`
*/
