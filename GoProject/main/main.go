package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
การใช้ฟังก์ชัน runtime.GOMAXPROCS() ในการกำหนดจำนวน logical processor ที่ Go runtime จะใช้ในการประมวลผล Goroutine
*/

func f1(name string, wg *sync.WaitGroup) {
	defer wg.Done() // ลด counter ของ WaitGroup เมื่อฟังก์ชันทำงานเสร็จ
	for index := 0; index < 10; index++ {
		fmt.Printf("%v: index %d\n", name, index) // พิมพ์ข้อความ
	}
}

func main() {
	runtime.GOMAXPROCS(2) // กำหนดให้ใช้ logical processor 2 ตัว
	var wg sync.WaitGroup
	wg.Add(2)        // สร้าง Goroutine 2 ตัว
	go f1("F1", &wg) // Goroutine 1
	go f1("F2", &wg) // Goroutine 2

	fmt.Println("Main: Waiting for Goroutines to finish")
	wg.Wait() // รอจนกว่า Goroutine ทั้งหมดจะทำงานเสร็จ
	fmt.Println("Main completed")
}

/*
1. `runtime.GOMAXPROCS(2)` กำหนดให้ Go runtime ใช้ logical processor 2 ตัวในการประมวลผล Goroutine
2. สร้าง Goroutine 2 ตัว โดยแต่ละตัวเรียกใช้ฟังก์ชัน `f1()`
3. `wg.Wait()` ในฟังก์ชัน `main()` จะรอจนกว่า Goroutine ทั้งสองตัวจะทำงานเสร็จ

เนื่องจากมีการกำหนด `GOMAXPROCS` เป็น 2 ดังนั้นถ้าเครื่องมี CPU มากกว่า 1 core Goroutine ทั้งสองจะสามารถทำงานแบบขนานกัน (parallel) ได้ แต่ถ้าเครื่องมี CPU เพียง 1 core Goroutine จะทำงานแบบ concurrency โดยใช้ multi-threading

ผลลัพธ์ของโปรแกรมจะไม่แน่นอน เนื่องจากการประมวลผลของ Goroutine ขึ้นอยู่กับการจัดสรรทรัพยากรของระบบ แต่จะสังเกตได้ว่าข้อความจาก Goroutine ทั้งสองจะถูกพิมพ์ออกมาแบบสลับกันไปมา

การกำหนดค่า `GOMAXPROCS` เป็นวิธีหนึ่งในการควบคุมว่า Go runtime จะใช้ CPU กี่ตัวในการประมวลผล ซึ่งมีผลต่อประสิทธิภาพในการรันโปรแกรมแบบ parallel
*/
