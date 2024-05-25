package main

import (
	"fmt"
	"time"
)

/*
เกี่ยวกับ Ticker ในแพ็คเกจ time ของภาษา Go:

- Ticker เป็นช่องสัญญาณที่ส่งเหตุการณ์ต่อเนื่องในช่วงเวลาที่กำหนด
- สร้าง Ticker ได้โดยใช้ฟังก์ชัน `time.NewTicker(duration)`
- อ่านเหตุการณ์จากช่องสัญญาณของ Ticker ได้โดยใช้ตัวดำเนินการ `<-`
- ปิดช่องสัญญาณของ Ticker ได้โดยเรียกใช้เมธอด `Stop()`
- ตัวอย่างโปรแกรมแสดงการใช้งาน Ticker ร่วมกับ goroutine และ select statement เพื่อจัดการเหตุการณ์จากหลาย Ticker พร้อมกัน

Ticker เป็นเครื่องมือที่มีประโยชน์สำหรับการสร้างเหตุการณ์ที่เกิดขึ้นเป็นระยะๆ ตามช่วงเวลาที่กำหนด ช่วยให้สามารถควบคุมการทำงานของโปรแกรมได้อย่างเป็นระบบ
*/

func receiver(ticker1 *time.Ticker, ticker2 *time.Ticker) {
	for {
		select {
		case msg := <-ticker1.C:
			// อ่านเหตุการณ์จาก ticker1 และพิมพ์ค่าที่ได้รับ
			fmt.Println("Message from ticker1", msg)
		case msg := <-ticker2.C:
			// อ่านเหตุการณ์จาก ticker2 และพิมพ์ค่าที่ได้รับ
			fmt.Println("Message from ticker2", msg)
		}
	}
}

func main() {
	// สร้าง Ticker สองตัวที่ส่งเหตุการณ์ทุก 1000 และ 1300 มิลลิวินาทีตามลำดับ
	ticker1 := time.NewTicker(time.Millisecond * 1000)
	ticker2 := time.NewTicker(time.Millisecond * 1300)

	// เรียกใช้ฟังก์ชัน receiver แบบ goroutine
	go receiver(ticker1, ticker2)

	// รอ 3000 มิลลิวินาที
	time.Sleep(time.Millisecond * 3000)
	// หยุดการทำงานของ ticker1
	ticker1.Stop()

	// รอ 10000 มิลลิวินาที
	time.Sleep(time.Millisecond * 10000)
}

/*

 */
