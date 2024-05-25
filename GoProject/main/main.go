package main

import (
	"fmt"
	"time"
)

/*
โค้ดที่ไม่ใช้ Select Statement

1. ฟังก์ชัน `receiver` ถูกเรียกใช้แยกกันสำหรับแต่ละ channel ทำให้ต้องสร้าง goroutine เพิ่มเติม และโค้ดยาวขึ้น

2. ในฟังก์ชัน `receiver` การรับข้อมูลจาก channel จะเป็นแบบ blocking ถ้า channel ยังไม่มีข้อมูลส่งมา goroutine จะหยุดรอ ทำให้ไม่สามารถตรวจสอบ channel อื่นๆ ได้ในขณะที่รอ

3. ไม่สามารถกำหนดลำดับความสำคัญในการรับข้อมูลจาก channel ได้ ขึ้นอยู่กับว่า channel ใดมีข้อมูลส่งมาก่อน

4. ในฟังก์ชัน `main` ต้องใช้ `time.Sleep` เพื่อรอให้ goroutine ทำงานเสร็จ ซึ่งไม่ใช่วิธีที่ดีนัก เพราะไม่สามารถกำหนดเวลาที่แน่นอนได้ และอาจทำให้โปรแกรมรอนานเกินไป
*/

func receiver(ch chan string) {
	for {
		msg, ok := <-ch
		if !ok {
			// ถ้า channel ถูกปิด ให้ออกจากลูป
			break
		}
		fmt.Println("Message from channel:", msg)
	}
}

func producer(ch chan string, name string, sleep time.Duration) {
	for index := 0; index < 10; index++ {
		time.Sleep(sleep)
		ch <- fmt.Sprintf("%v: %v", name, index)
	}
	close(ch)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go producer(ch1, "P1", time.Millisecond*1000)
	go producer(ch2, "P2", time.Millisecond*1300)

	go receiver(ch1)
	go receiver(ch2)

	// รอเพื่อให้ goroutine ทำงานเสร็จ
	time.Sleep(time.Second * 15)
}

/*

 */
