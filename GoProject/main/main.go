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
อาจได้ค่าระหว่าง 900 ถึง 1000 เนื่องจากเกิดปัญหา race condition ซึ่งเป็นสถานการณ์ที่หลาย Goroutine พยายามเข้าถึงและแก้ไขค่าของตัวแปรร่วม (shared variable) `amount` พร้อมกัน โดยไม่มีการควบคุมการเข้าถึงอย่างเหมาะสม

จากโค้ด `amount = amount - 1` ในฟังก์ชัน `withdraw()` ประกอบด้วยขั้นตอนดังนี้:

1. อ่านค่าปัจจุบันของ `amount`
2. ลบค่า 1 จากค่าที่อ่านมา
3. เขียนค่าใหม่กลับไปยัง `amount`

เนื่องจากมีหลาย Goroutine ทำงานพร้อมกัน บางครั้งขั้นตอนเหล่านี้อาจถูกแทรกสอดจากการทำงานของ Goroutine อื่น ทำให้การอัพเดทค่าไม่ถูกต้อง

ยกตัวอย่างเช่น:

1. Goroutine A อ่านค่า `amount = 1000`
2. Goroutine B อ่านค่า `amount = 1000` ก่อนที่ A จะอัพเดทค่าใหม่
3. Goroutine B ลบ 1 และเขียนค่า `amount = 999` กลับ
4. Goroutine A ลบ 1 จากค่าเดิม 1000 และเขียนค่า `amount = 999` กลับ (ซ้ำกับ B)

ในกรณีนี้จะได้ค่า `amount = 999` แทนที่จะได้ 900 ตามที่คาดหวัง เนื่องจากมีการอัพเดทค่าซ้ำกันจากหลาย Goroutine

ดังนั้น เมื่อรันโปรแกรมหลายๆครั้ง อาจได้ผลลัพธ์ที่แตกต่างกันไป ระหว่าง 900 ถึง 1000 ขึ้นอยู่กับลำดับการทำงานของ Goroutine แต่ละครั้ง
*/
