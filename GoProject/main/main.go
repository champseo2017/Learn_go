package main

import (
	"fmt"
	"sync"
	"time"
)

/*
โจทย์ให้เขียนโปรแกรม producer-consumer ที่มี producer 2 ตัว ผลิตข้อมูลทุก 300 และ 400 มิลลิวินาที และมี consumer 1 ตัว โดย producer ผลิตข้อมูลแค่ 100 ครั้ง
*/

func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // เรียก wg.Done() เมื่อ producer ทำงานเสร็จ

	interval := 300 // กำหนดช่วงเวลาเริ่มต้นเป็น 300 มิลลิวินาที
	if id == 2 {
		interval = 400 // ถ้า id เป็น 2 ให้กำหนดช่วงเวลาเป็น 400 มิลลิวินาที
	}

	for i := 1; i <= 100; i++ {
		ch <- i                                                // ส่งข้อมูล i ไปยัง channel ch
		fmt.Printf("Producer %d produced data: %d\n", id, i)   // แสดงข้อความว่า producer ผลิตข้อมูลอะไร
		time.Sleep(time.Duration(interval) * time.Millisecond) // หยุดพักเป็นเวลา interval มิลลิวินาที
	}
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range ch {
		fmt.Printf("Consumer consumed data: %d\n", data)
	}
}

func main() {
	ch := make(chan int)
	/*
		   ch := make(chan int, 100) ใช้ buffered channel เพื่อเพิ่มประสิทธิภาพ
			ทำงานแบบ asynchronous โดยที่ producer สามารถส่งข้อมูลเข้า channel ได้แม้ว่า consumer จะยังไม่ได้รับข้อมูลนั้นในทันที และจะทำงานได้อย่างราบรื่นขึ้นเพราะ buffered channel ที่ใช้ในการเก็บข้อมูลชั่วคราว
	*/
	var wg sync.WaitGroup

	// สร้าง goroutine ของ producer 2 ตัว
	wg.Add(2)
	go producer(1, ch, &wg)
	go producer(2, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	/*
		ใช้อีกหนึ่ง WaitGroup (consumerWg) เพื่อจัดการการรอคอยของ consumer หลังจากที่ producer ทั้งหมดทำงานเสร็จ เราจะปิด channel เพื่อให้ consumer หยุดทำงานอย่างถูกต้องและไม่มี deadlock เกิดขึ้น
	*/
	var consumerWg sync.WaitGroup
	consumerWg.Add(1)
	go consumer(ch, &consumerWg)

	// รอให้ consumer ทำงานเสร็จ
	consumerWg.Wait()
}

/*

 */
