package main

import (
	"sync"
)

/*
เกี่ยวกับการปิด channel:

- เราสามารถปิด channel ได้โดยใช้ฟังก์ชัน close
- เมื่อ channel ถูกปิดแล้ว จะไม่สามารถส่งข้อมูลเข้าไปใน channel นั้นได้อีก
- หากพยายามส่งข้อมูลเข้าไปใน channel ที่ปิดแล้ว โค้ดจะเกิด panic
- การรับข้อมูลจาก channel ที่ปิดแล้วยังสามารถทำได้ โดยจะได้รับค่า zero value ของประเภทข้อมูลที่ channel เก็บ
- การปิด channel ที่ปิดไปแล้วจะไม่เกิดข้อผิดพลาดใดๆ
- การปิด channel เป็นวิธีการบอกผู้รับว่าจะไม่มีข้อมูลส่งมาอีกแล้ว

การปิด channel เป็นการสื่อสารระหว่างผู้ส่งและผู้รับว่าไม่มีข้อมูลที่จะส่งผ่าน channel นั้นอีกต่อไป ซึ่งเป็นประโยชน์ในการควบคุมการทำงานของโปรแกรมที่ใช้ goroutine และ channel ในการสื่อสารและประสานงานกัน
*/

func writeVal(ch chan int, wg *sync.WaitGroup) {
	ch <- 10  // พยายามส่งค่า 10 เข้าไปใน channel ที่ปิดแล้ว (เกิด panic)
	wg.Done() // ลดจำนวน goroutine ใน WaitGroup
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)            // เพิ่มจำนวน goroutine ใน WaitGroup
	ch := make(chan int) // สร้าง channel
	close(ch)            // ปิด channel
	go writeVal(ch, &wg) // สร้าง goroutine เพื่อเรียกฟังก์ชัน writeVal และส่ง channel และ pointer ของ WaitGroup เข้าไป
	wg.Wait()            // รอให้ทุก goroutine ใน WaitGroup ทำงานเสร็จ
}

/*

 */
