package main

import (
	"fmt"
	"sync"
)

/*
แก้ไขปัญหา race condition จากตัวอย่างก่อนหน้า โดยใช้ sync.Mutex เพื่อควบคุมการเข้าถึงตัวแปรร่วม amount ในช่วง critical section
*/

var amount = 1000 // ตัวแปรร่วม amount

func withDraw(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()     // ลด counter ของ WaitGroup เมื่อฟังก์ชันทำงานเสร็จ
	mu.Lock()           // ล็อก mutex ก่อนเข้าถึง critical section
	amount = amount - 1 // critical section ลบค่า amount
	mu.Unlock()         // ปลดล็อก mutex หลังจากทำงานใน critical section เสร็จ
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex // สร้างตัวแปร mutex
	wg.Add(100)       // สร้าง Goroutine 100 ตัว
	for index := 0; index < 100; index++ {
		go withDraw(&wg, &mu) // ส่งค่าอ้างอิงของ mutex ไปยัง withdraw
	}

	wg.Wait()           // รอให้ Goroutine ทั้งหมดทำงานเสร็จ
	fmt.Println(amount) // พิมพ์ค่า amount หลังจบการทำงาน
}

/*
1. สร้างตัวแปร `mu` ประเภท `sync.Mutex` ในฟังก์ชัน `main()`
2. ส่งค่าอ้างอิง (address) ของ `mu` ไปยังฟังก์ชัน `withdraw()` ด้วย
3. ในฟังก์ชัน `withdraw()` ใช้ `mu.Lock()` เพื่อล็อกการเข้าถึง critical section ก่อนจะลบค่า `amount`
4. หลังจากลบค่า `amount` เสร็จแล้ว ให้ปลดล็อกด้วย `mu.Unlock()`

เมื่อใช้ mutex lock แล้ว จะมี Goroutine เพียงตัวเดียวเท่านั้นที่สามารถผ่านเข้าไปทำงานใน critical section ได้ในขณะนั้น ทำให้ไม่เกิดปัญหา race condition ขึ้น และจะได้ผลลัพธ์ที่ถูกต้องคงที่คือ `amount = 900` หลังจากทุก Goroutine ทำงานเสร็จ

อย่างไรก็ตาม การใช้ mutex lock มีข้อเสียคือจะชะลอประสิทธิภาพการทำงานของโปรแกรมลง เนื่องจากมีการรอคิวในการเข้าถึง critical section ดังนั้นจึงควรใช้ให้น้อยที่สุดเท่าที่จำเป็น
*/
