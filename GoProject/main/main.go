package main

import (
	"fmt"
	"time"
)

/*
แก้ไขปัญหา deadlock ที่เกิดขึ้นเมื่อส่งข้อมูลไปยัง unbuffered channel โดยไม่มี goroutine รับข้อมูล เราสามารถสร้าง goroutine ที่คอยรับข้อมูลจาก channel
*/

func main() {
	ch := make(chan int)
	// สร้าง goroutine ที่คอยรับข้อมูลจาก channel
	// แก้ไขปัญหา deadlock ที่เกิดขึ้นเมื่อส่งข้อมูลไปยัง unbuffered channel โดยไม่มี goroutine รับข้อมูล เราสามารถสร้าง goroutine ที่คอยรับข้อมูลจาก channel
	go func() {
		data := <-ch
		fmt.Println("Received data:", data)
	}()

	ch <- 10 // ส่งข้อมูลไปยัง unbuffered channel
	// โปรแกรมจะไม่เกิด deadlock เพราะมี goroutine คอยรับข้อมูล

	// รอให้ goroutine ทำงานเสร็จ
	// ในตัวอย่างนี้ใช้ time.Sleep เพื่อให้ง่ายต่อการทำความเข้าใจ
	// ในการใช้งานจริง ควรใช้ sync.WaitGroup หรือวิธีอื่นที่เหมาะสม
	time.Sleep(time.Second)
}

/*

 */
