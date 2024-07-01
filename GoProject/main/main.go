package main

import (
	// oops "GoProject/package"
	"fmt"
)

/*
สรุปสั้นๆ เกี่ยวกับการใช้งาน Channel ในภาษา Go

- สร้าง Channel ด้วย `make(chan data_type)`
- ปิด Channel ด้วย `close(channel_name)` มักใช้คู่กับ `defer` เพื่อปิดเมื่อฟังก์ชันทำงานเสร็จ
- ใส่ข้อมูลลงใน Channel ด้วย `channel_name <- data`
- อ่านข้อมูลจาก Channel ด้วย `variable := <-channel_name`

Channel ช่วยให้เราสื่อสารและส่งข้อมูลระหว่าง Goroutine (ฟังก์ชันที่ทำงานพร้อมกัน) ได้อย่างมีประสิทธิภาพ ทำให้การเขียนโปรแกรมแบบ Concurrent ในภาษา Go เป็นเรื่องง่ายและสะดวกยิ่งขึ้นครับ
*/

func main() {
	// สร้าง Channel ชื่อ numChannel ที่เก็บข้อมูลประเภท int
	numChannel := make(chan int)

	go func(numChannel chan int) {
		// ปิด Channel เมื่อฟังก์ชันทำงานเสร็จ
		defer close(numChannel)

		// ใส่ข้อมูล 100 ลงใน numChannel
		numChannel <- 100
	}(numChannel)

	// อ่านข้อมูลจาก numChannel มาเก็บไว้ในตัวแปร num
	num := <-numChannel

	// แสดงค่าของ num
	fmt.Println(num)
}

/*

 */
