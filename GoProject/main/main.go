package main

import (
	"fmt"
)

/*
ตัวอย่างของการใช้ unbuffered channel ในการสื่อสารระหว่าง goroutine ของ producer และ consumer โดย producer จะทำการส่งข้อมูลเข้าไปใน channel และ consumer จะทำการอ่านข้อมูลจาก channel ซึ่งการอ่านและเขียนข้อมูลผ่าน unbuffered channel จะเกิดขึ้นแบบ synchronous กล่าวคือ เมื่อ producer ส่งข้อมูล จะต้องรอให้ consumer อ่านข้อมูลก่อน จึงจะสามารถส่งข้อมูลต่อไปได้ และในทำนองเดียวกัน consumer จะต้องรอให้ producer ส่งข้อมูลเข้ามาก่อน จึงจะสามารถอ่านข้อมูลได้
*/

func producer(ch chan int, done chan bool) {
	for index := 0; index < 10; index++ {
		fmt.Printf("PRODUCER: sending %v\n", index)
		ch <- index // ส่งค่า index เข้าไปใน channel
	}
	close(ch)    // ปิด channel เมื่อส่งข้อมูลครบแล้ว
	done <- true // ส่งสัญญาณว่า producer ทำงานเสร็จแล้ว
}

func consumer(ch chan int) {
	for val := range ch { // อ่านข้อมูลจาก channel จนกว่าจะปิด
		fmt.Printf("CONSUMER: read %v\n", val)
	}
}

func main() {
	ch := make(chan int)    // สร้าง unbuffered channel สำหรับส่งข้อมูลชนิด int
	done := make(chan bool) // สร้าง channel สำหรับส่งสัญญาณเมื่อ producer ทำงานเสร็จ

	go producer(ch, done) // สร้าง goroutine สำหรับ producer
	go consumer(ch)       // สร้าง goroutine สำหรับ consumer

	<-done // รอจนกว่า producer จะทำงานเสร็จ
	/*
		<-done ในฟังก์ชัน main() มีความสำคัญในการรอให้ goroutine ของ producer ทำงานเสร็จก่อนที่โปรแกรมจะจบการทำงาน
		เนื่องจากในภาษา Go เมื่อเรียกใช้ goroutine โปรแกรมจะไม่รอให้ goroutine นั้นทำงานเสร็จ แต่จะทำงานต่อไปเรื่อยๆ จนกว่าจะจบฟังก์ชัน main() ดังนั้น หากไม่มีบรรทัด <-done โปรแกรมอาจจบการทำงานก่อนที่ producer จะทำงานเสร็จ ทำให้ไม่ได้ผลลัพธ์ตามที่ต้องการ
		บรรทัด <-done เป็นการรอรับสัญญาณจาก done channel ซึ่ง producer จะส่งสัญญาณมาเมื่อทำงานเสร็จแล้ว การรอรับสัญญาณนี้จะทำให้ goroutine ของ main() หยุดรอจนกว่าจะได้รับสัญญาณ ซึ่งก็คือรอจนกว่า producer จะทำงานเสร็จนั่นเอง
	*/
}

/*

 */
