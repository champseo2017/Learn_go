package main

import (
	"fmt"
	"time"
)

/*
เกี่ยวกับ Select Statement ในภาษา Go:
- Select Statement ใช้สำหรับจัดการ goroutine ที่ส่งข้อมูลผ่านหลาย channel พร้อมกัน
- มีไวยากรณ์คล้ายกับ Switch Statement โดยมีหลาย case ที่แต่ละ case จะมีการดำเนินการกับ channel
- เมื่อ case ใดสามารถรับข้อมูลจาก channel ได้ case นั้นจะถูกเลือกให้ทำงาน
- สามารถมี default case เมื่อไม่มี case ใดสามารถรับข้อมูลจาก channel ได้
- ตัวอย่างโปรแกรมประกอบด้วย:
  - ฟังก์ชัน `producer` ทำหน้าที่ส่งข้อมูลเข้า channel
  - ฟังก์ชัน `receiver` ทำหน้าที่รับข้อมูลจาก channel ด้วย select statement
- Select Statement ช่วยให้จัดการและประมวลผลข้อมูลจากหลาย channel ได้อย่างมีประสิทธิภาพ
*/

func receiver(ch1 chan string, ch2 chan string) {
	for {
		select {
		case msg, ok := <-ch1:
			// รับข้อมูลจาก ch1
			fmt.Println("Message from ch1", msg)
			if !ok {
				// ถ้า ch1 ถูกปิด ให้ออกจากฟังก์ชัน
				return
			}
		case msg, ok := <-ch2:
			// รับข้อมูลจาก ch2
			fmt.Println("Message from ch2", msg)
			if !ok {
				// ถ้า ch2 ถูกปิด ให้ออกจากฟังก์ชัน
				return
			}
		}
	}
}

func producer(ch chan string, name string, sleep time.Duration) {
	for index := 0; index < 10; index++ {
		// หน่วงเวลาตามค่า sleep ที่กำหนด
		time.Sleep(sleep)
		// ส่งข้อมูลเข้า channel
		ch <- fmt.Sprintf("%v: %v", name, index)
	}
	// ปิด channel เมื่อส่งข้อมูลครบ
	close(ch)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// สร้าง goroutine สำหรับ producer ที่ส่งข้อมูลผ่าน ch1 และ ch2 ตามลำดับ
	go producer(ch1, "P1", time.Millisecond*1000)
	go producer(ch2, "P2", time.Millisecond*1300)

	// เรียกใช้ฟังก์ชัน receiver เพื่อรับข้อมูลจาก ch1 และ ch2
	receiver(ch1, ch2)
}

/*

 */
