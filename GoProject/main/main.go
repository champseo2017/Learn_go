package main

import (
	"fmt"
)

/*
โจทย์ข้อที่ 4: กำหนด read-only channel เป็น argument ของฟังก์ชัน
*/

func read(ch <-chan int) {
	// ใช้ <-chan แสดงว่าเป็น receive-only channel
	// ไม่สามารถส่งข้อมูลเข้า channel ได้
	// วนลูปอ่านข้อมูลจาก channel จนกว่า channel จะถูกปิด
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				// channel ถูกปิดแล้ว ออกจากลูป
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Received data:", data)
		}
	}
}

func main() {
	ch := make(chan int)

	// สร้าง goroutine เพื่อส่งข้อมูลเข้า channel
	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
		close(ch)
	}()

	// เรียกใช้ฟังก์ชัน read เพื่ออ่านข้อมูลจาก channel
	read(ch)
}

/*

 */
