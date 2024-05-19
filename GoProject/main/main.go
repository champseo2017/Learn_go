package main

import (
	"fmt"
	"sync"
)

/*
นำ sync.WaitGroup มาใช้ในการจัดการ Goroutine ให้รันจนเสร็จสิ้นก่อนที่จะออกจากโปรแกรม
*/

// จัดการ Goroutine ด้วย sync.WaitGroup
var wg sync.WaitGroup

func hello() {
	defer wg.Done() // ถ้ากระบวนการในฟังก์ชันเสร็จแล้วให้ลด counter ของ WaitGroup
	/*
		เมื่อออกจากฟังก์ชัน hello() แล้ว ระบบจะมาทำงานคำสั่งที่ถูก defer ไว้คือ wg.Done() ซึ่งจะลดค่า counter ของ sync.WaitGroup ลง 1

		การใช้ defer ในที่นี้ช่วยให้เราสามารถลด counter ของ sync.WaitGroup ได้หลังจากฟังก์ชัน hello() ทำงานเสร็จสิ้นแล้ว ซึ่งเป็นสิ่งสำคัญในการควบคุมการทำงานของ Goroutine ให้ทำงานจนเสร็จสมบูรณ์ก่อนจะจบโปรแกรม
	*/
	fmt.Println("Hello Gophers!")
}

func main() {
	wg.Add(2) // เพิ่มค่า counter ของ WaitGroup เป็น 2

	go hello() // สร้าง Goroutine แรก
	go hello() // สร้าง Goroutine ที่สอง

	wg.Wait() // รอจนกว่า counter ของ WaitGroup จะเป็น 0 ค่อยจบการทำงาน

	fmt.Println("All Goes completed.")
}

/*
1. `var wg sync.WaitGroup` สร้างตัวแปร `wg` เป็น struct `sync.WaitGroup` ใช้ในการจัดการ Goroutine
2. ในฟังก์ชัน `hello()` มีการเรียกใช้ `wg.Done()` ผ่าน `defer` เพื่อลด counter ของ WaitGroup หลังจากทำงานในฟังก์ชันเสร็จแล้ว
3. ในฟังก์ชัน `main()` เรียกใช้ `wg.Add(2)` เพื่อตั้งค่า counter ของ WaitGroup เป็น 2 เนื่องจากมี 2 Goroutine ที่จะสร้างขึ้น
4. สร้าง Goroutine 2 ตัว ด้วยคำสั่ง `go hello()`
5. `wg.Wait()` จะรอจนกว่า counter ของ WaitGroup จะลดลงเป็น 0 ค่อยจบการทำงาน ซึ่งจะถูกลดลงเมื่อฟังก์ชัน `hello()` เสร็จสิ้นและเรียก `wg.Done()`

การใช้ sync.WaitGroup ทำให้สามารถรอจนกว่า Goroutine ทั้งหมดจะทำงานเสร็จก่อนจบโปรแกรม หลีกเลี่ยงปัญหา race condition ได้
*/
