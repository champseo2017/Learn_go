package main

import (
	"fmt"
	"sync"
)

/*
Goroutines สามารถใช้เขียนโปรแกรมแบบ concurrent ใน Golang ได้ เมื่อ goroutines หลายตัวทำงานพร้อมกัน ลำดับการทำงานของมันจะคาดเดาไม่ได้ goroutine ไหนก็ได้สามารถเริ่มก่อนหรือจบก่อนก็ได้ เมื่อ main goroutine จบการทำงาน goroutines ลูกทั้งหมดก็จะจบตามไปด้วย ถ้าต้องการให้ goroutine แม่รอ goroutine ลูกทำงานให้เสร็จก่อน สามารถใช้ WaitGroup ในการควบคุมได้

WaitGroup มี 3 method ที่ใช้บ่อยคือ Add, Done และ Wait นอกจากนี้ยังสามารถใช้ MutexLock ในการป้องกัน critical section เพื่อให้มีแค่ goroutine เดียวเข้าไปทำงานใน critical section ได้

Channel ใช้สำหรับการสื่อสารระหว่าง goroutines สองตัว โดยใช้ operator <- ในการส่งหรืออ่านข้อความเข้าหรือออกจาก channel เมื่อปิด channel แล้วจะไม่สามารถส่งข้อความเข้า channel ได้อีก และสามารถใช้ range clause อ่านข้อความจาก channel ได้ เมื่อ channel ถูกปิดและอ่านข้อความครบแล้ว for loop จะจบการทำงาน

Channel มีสองแบบคือ buffered และ unbuffered ในการกำหนด channel เป็น argument ของ function สามารถกำหนดเป็น read-only หรือ write-only channel ได้ และยังสามารถใช้ select statement เมื่อต้องการสื่อสารระหว่าง channel หลายตัว
*/

var wg sync.WaitGroup

func printNumbers(prefix string) {
	defer wg.Done() // แจ้ง WaitGroup ว่า goroutine นี้เสร็จแล้ว
	/*
		สรุปได้ว่า defer จะเป็นคำสั่งที่ทำให้ฟังก์ชันที่อยู่ในคำสั่ง defer ถูกเรียกทุกครั้งหลังจากฟังก์ชันที่มีคำสั่ง defer นั้นจบการทำงาน ไม่ว่าฟังก์ชันนั้นจะจบการทำงานปกติหรือจบด้วยการเกิด panic ก็ตาม
	*/

	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", prefix, i)
	}
}

func main() {
	wg.Add(2) // เพิ่มจำนวน goroutine ที่ต้องรอใน WaitGroup

	go printNumbers("Goroutine 1") // สร้าง goroutine แรก
	go printNumbers("Goroutine 2") // สร้าง goroutine ที่สอง

	wg.Wait() // รอจนกว่าทุก goroutine จะเสร็จ
}

/*

 */
