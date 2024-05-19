package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
แสดงการใช้ฟังก์ชัน Goexit() เพื่อหยุดการทำงานของ goroutine โดยเมื่อ goroutine f1 ทำงานถึง index เท่ากับ 5 จะเรียกใช้ฟังก์ชัน Goexit() ทำให้ goroutine f1 หยุดทำงานทันที ในขณะที่ฟังก์ชัน main จะรอจนกว่า goroutine f1 จะทำงานเสร็จก่อนจึงจะสิ้นสุดการทำงาน
*/

func f1(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for index := 0; index < 10; index++ {
		if index == 5 {
			runtime.Goexit() // เรียกใช้ฟังก์ชัน Goexit() เมื่อ index เท่ากับ 5 เพื่อหยุดการทำงานของ goroutine f1
		}
		fmt.Printf("%v: index %d\n", name, index)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go f1("F1", &wg)
	fmt.Println("Main: Waiting for Goroutines to finish")
	wg.Wait()
	fmt.Println("Main completed")
}

/*
คำอธิบาย code:
- ฟังก์ชัน main สร้าง WaitGroup เพื่อใช้ในการรอ goroutine f1 ทำงานเสร็จ และเรียกใช้ goroutine f1
- goroutine f1 ทำงานโดยมี loop ตั้งแต่ index 0 ถึง 9 และเช็คเงื่อนไขว่าถ้า index เท่ากับ 5 ให้เรียกใช้ฟังก์ชัน Goexit() เพื่อหยุดการทำงานของ goroutine f1 ทันที
- เมื่อ goroutine f1 ทำงานเสร็จหรือถูกหยุดด้วยฟังก์ชัน Goexit() จะส่งสัญญาณผ่าน WaitGroup ให้ฟังก์ชัน main รู้ว่า goroutine f1 ทำงานเสร็จแล้ว
- ฟังก์ชัน main จะรอจนกว่า goroutine f1 จะทำงานเสร็จ โดยใช้ wg.Wait() ก่อนจะสิ้นสุดการทำงานของโปรแกรม

ผลลัพธ์ที่ได้จากการรันโปรแกรมคือ goroutine f1 จะทำงานและแสดงผลลัพธ์ตั้งแต่ index 0 ถึง 4 เท่านั้น เนื่องจากเมื่อ index เท่ากับ 5 จะเรียกใช้ฟังก์ชัน Goexit() ทำให้ goroutine f1 หยุดทำงานทันที จากนั้นฟังก์ชัน main จะแสดงข้อความ "Main completed" และสิ้นสุดการทำงานของโปรแกรม
*/
