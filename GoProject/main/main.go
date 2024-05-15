package main

import (
	"fmt"
	"sync"
)

/*
เกี่ยวกับ Concurrency ใน Go:

- โปรแกรม concurrent ใน Go เรียกว่า goroutines ซึ่งสามารถรันขนานกันได้
- ทุกโปรแกรม Go มีอย่างน้อยหนึ่ง goroutine คือ main goroutine
- ใช้ WaitGroup เพื่อรอให้ goroutines ทำงานเสร็จ
- ใช้ channel สำหรับการสื่อสารระหว่าง goroutines
- มี goroutines สองแบบ: buffered และ unbuffered
- เมื่อรันหลาย threads บน single core เรียกว่าโปรแกรม concurrent
- เมื่อรันหลาย threads บนหลาย cores เรียกว่าโปรแกรม parallelized
*/

func main() {

	var wg sync.WaitGroup // สร้าง WaitGroup เพื่อรอให้ goroutines ทำงานเสร็จ

	// สร้าง goroutine แบบ unbuffered
	wg.Add(1) // เพิ่มจำนวน goroutine ที่ต้องรอ
	go func() {
		defer wg.Done() // ลดจำนวน goroutine เมื่อทำงานเสร็จ
		fmt.Println("Hello from unbuffered goroutine")
	}()

	// สร้าง goroutine แบบ buffered ด้วย channel
	ch := make(chan string, 1) // สร้าง channel แบบ buffered ขนาด 1
	go func() {
		ch <- "Hello from buffered goroutine" // ส่งข้อความผ่าน channel
	}()

	wg.Wait() // รอให้ goroutines ทั้งหมดทำงานเสร็จ

	msg := <-ch // รับข้อความจาก channel
	fmt.Println(msg)

}

/*

 */
