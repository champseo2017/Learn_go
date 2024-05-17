package main

import (
	"fmt"
	"sync"
	"time"
)

/*
ในส่วนก่อนหน้านี้ เราใช้ฟังก์ชัน Sleep เพื่อรอให้ goroutine ลูกทำงานจนเสร็จ แต่ฟังก์ชัน Sleep ไม่ใช่วิธีปฏิบัติที่ดีในการรอ goroutine ลูกให้เสร็จสิ้น เราสามารถใช้ WaitGroup ของแพ็คเกจ sync เพื่อรอให้ goroutine ลูกทำงานจนเสร็จ

sync.WaitGroup เป็นประเภท struct และมีเมธอดที่มีประโยชน์สามเมธอด:

1. Add: เมธอด Add เพิ่มจำนวนเต็มให้กับ WaitGroup จำนวนเต็มนี้เป็นตัวนับที่บอกว่าฟังก์ชันหลักจะรอจำนวน goroutine เท่านั้นให้เสร็จสมบูรณ์

2. Done: เมธอด Done ลดตัวนับที่เพิ่มในฟังก์ชัน Add เราควรเรียกเมธอดนี้ที่ส่วนท้ายของ goroutine

3. Wait: เมธอด Wait รอให้ goroutines ทั้งหมดทำงานจนเสร็จ เมื่อตัวนับเป็น 0 มันจะสิ้นสุดการรอ
*/

func f1(name string, wg *sync.WaitGroup) {
	// ฟังก์ชัน f1 รับพารามิเตอร์ name เป็นสตริง และ wg เป็น pointer ของ WaitGroup

	for index := 0; index < 10; index++ {
		// ลูป for วนซ้ำ 10 ครั้งตั้งแต่ 0 ถึง 9
		fmt.Printf("%v: index %d\n", name, index)
		// พิมพ์ค่า name และ index ของแต่ละรอบลูป
		time.Sleep(1 * time.Second)
		// หยุดพักเป็นเวลา 1 วินาที
	}

	wg.Done()
	// ลดค่า wg ลงหนึ่งเมื่อ goroutine นี้ทำงานเสร็จแล้ว
}

func main() {

	var wg sync.WaitGroup
	// สร้างตัวแปร wg เป็น WaitGroup เพื่อใช้ในการรอ goroutines

	wg.Add(2)
	// เพิ่มค่าให้ wg เป็น 2 เพื่อบอกว่าต้องรอ 2 goroutines

	go f1("F1", &wg)
	go f1("F2", &wg)
	// เรียกฟังก์ชัน f1 สองครั้งเป็น goroutines โดยส่ง "F1" และ "F2" เป็นอาร์กิวเมนต์ และส่ง pointer ของ wg เข้าไปด้วย

	fmt.Println("Main: Waiting for Goroutines to finish")
	// พิมพ์ข้อความ "Main: Waiting for Goroutines to finish"

	wg.Wait()
	// รอให้ goroutines ทั้งหมดทำงานจนเสร็จโดยใช้ wg.Wait()

	fmt.Println("Main completed")
	// พิมพ์ข้อความ "Main completed" เมื่อ goroutines ทำงานเสร็จแล้ว

}

/*
โปรแกรมนี้ใช้ WaitGroup เพื่อรอให้ goroutines ทั้งหมดทำงานจนเสร็จก่อนที่ฟังก์ชัน main จะจบการทำงาน มันเรียกฟังก์ชัน f1 สองครั้งเป็น goroutines และส่ง pointer ของ WaitGroup เข้าไปด้วย แต่ละ goroutine จะพิมพ์ชื่อ (F1 หรือ F2) และดัชนีของลูปจาก 0 ถึง 9 โดยหยุดพักเป็นเวลา 1 วินาทีในแต่ละรอบ เมื่อ goroutine ทำงานเสร็จ มันจะเรียก wg.Done() เพื่อลดค่า WaitGroup ลง

ฟังก์ชัน main จะรอให้ goroutines ทั้งหมดทำงานจนเสร็จโดยใช้ wg.Wait() ก่อนที่จะพิมพ์ข้อความ "Main completed" และจบโปรแกรม วิธีนี้ทำให้มั่นใจได้ว่า goroutines จะทำงานจนเสร็จสมบูรณ์ก่อนที่โปรแกรมจะจบลง

ผลลัพธ์ของโปรแกรมจะแสดงข้อความที่ goroutines พิมพ์ออกมา ตามด้วยข้อความ "Main completed" เมื่อ goroutines ทำงานเสร็จแล้ว
*/
