package main

import (
	"fmt"
)

/*
สรุป:
- Channel ใน Golang ใช้สำหรับการสื่อสารระหว่าง goroutine
- ประกาศ channel ด้วยฟังก์ชัน make และกำหนดประเภทข้อมูลที่เก็บใน channel
- สามารถกำหนดขนาด buffer ให้ channel ได้ (ถ้าไม่กำหนดจะเป็น unbuffered channel)
- ใช้เครื่องหมาย <- สำหรับเขียนและอ่านค่าจาก channel
- การเขียนค่าลงใน channel ใช้รูปแบบ ch <- value
- การอ่านค่าจาก channel ใช้รูปแบบ variable := <-ch
*/

func main() {
	// ประกาศ channel ที่เก็บค่า integer
	ch := make(chan int)
	// สร้าง goroutine เพื่อเขียนค่าลงใน channel
	go func() {
		ch <- 10
	}()

	// อ่านค่าจาก channel และเก็บไว้ในตัวแปร a
	a := <-ch
	// แสดงค่าของตัวแปร a
	fmt.Println(a)
}

/*

 */
