package main

import (
	"fmt"
)

/*
buffer ในการประกาศ channel ด้วยฟังก์ชัน make:

- ถ้าไม่กำหนดขนาด buffer หรือกำหนดเป็น 0 จะได้ unbuffered channel
  - การส่งและรับข้อมูลผ่าน unbuffered channel จะเกิดขึ้นแบบ synchronous
- ถ้ากำหนดขนาด buffer มากกว่า 0 จะได้ buffered channel
  - การส่งและรับข้อมูลผ่าน buffered channel สามารถทำได้โดยไม่ต้องรอกัน (ภายใต้ขีดจำกัดของ buffer)
- ค่าเริ่มต้นของขนาด buffer คือ 0 ดังนั้นหากต้องการประกาศ unbuffered channel สามารถละการกำหนดขนาด buffer ได้

การอ่านค่าจาก channel ในฟังก์ชัน main จะถูก block จนกว่าฟังก์ชัน writeVal จะเขียนค่าลงใน channel เสร็จ ดังนั้นฟังก์ชัน main จะรอจนกว่าจะมีค่าใน channel
เนื่องจากการอ่านและเขียนค่าใน channel เป็น blocking call หากฟังก์ชัน main กำลังรอการอ่านค่าจาก channel แต่ไม่มี goroutine ที่เขียนค่าลงใน channel Go runtime จะตรวจพบว่าเป็น deadlock
*/

func writeVal(ch chan int) {
	ch <- 10 // เขียนค่า 10 ลงใน channel
}

func main() {
	ch := make(chan int, 0) // ประกาศ channel ที่เก็บค่า int และมีขนาด buffer เป็น 0
	go writeVal(ch)         // สร้าง goroutine เพื่อเรียกฟังก์ชัน writeVal และส่ง channel เข้าไป

	a := <-ch      // อ่านค่าจาก channel และเก็บไว้ในตัวแปร a (จะถูก block จนกว่าจะมีค่าใน channel)
	fmt.Println(a) // แสดงค่าของตัวแปร a
}

/*

 */
