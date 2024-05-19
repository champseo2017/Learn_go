package main

import (
	"fmt"
	"sync"
)

/*
ปัญหา race condition ที่เกิดขึ้นเมื่อมีการเข้าถึงและแก้ไขตัวแปรร่วมกัน (shared variable) จากหลาย Goroutine พร้อมกัน ซึ่งอาจทำให้ผลลัพธ์ไม่ตรงตามที่คาดหวัง
*/

var amount = 1000 // ตัวแปรร่วม amount มีค่าเริ่มต้นเท่ากับ 1000

func withdraw(wg *sync.WaitGroup) {
	defer wg.Done()     // ลด counter ของ WaitGroup เมื่อฟังก์ชันทำงานเสร็จ
	amount = amount - 1 // ลด amount ลง 1
}

func main() {
	var wg sync.WaitGroup
	wg.Add(100) // สร้าง Goroutine 100 ตัว

	for index := 0; index < 100; index++ {
		go withdraw(&wg) // แต่ละ Goroutine เรียกใช้ฟังก์ชัน withdraw
	}

	wg.Wait()           // รอให้ Goroutine ทั้งหมดทำงานเสร็จ
	fmt.Println(amount) // พิมพ์ค่า amount หลังจากทุก Goroutine ทำงานเสร็จ
}

/*
1. ตัวแปร `amount` ถูกสร้างขึ้นมาเป็นตัวแปรร่วม (shared variable) ที่มีค่าเริ่มต้นเท่ากับ 1000
2. ในฟังก์ชัน `main()` จะสร้าง Goroutine 100 ตัว โดยแต่ละ Goroutine จะเรียกใช้ฟังก์ชัน `withdraw()`
3. ฟังก์ชัน `withdraw()` จะลบค่า `amount` ลง 1 ทีละ Goroutine
4. หลังจากทุก Goroutine ทำงานเสร็จ จะพิมพ์ค่า `amount` ออกมา

เนื่องจากมีการเข้าถึงและแก้ไขค่า `amount` พร้อมกันจากหลาย Goroutine โดยไม่มีการควบคุม ทำให้เกิดปัญหา race condition ค่า `amount` ที่ได้จึงไม่แน่นอน อาจได้ค่าระหว่าง 900 ถึง 1000 แทนที่จะได้ 900 ตามที่คาดหวังไว้
*/
