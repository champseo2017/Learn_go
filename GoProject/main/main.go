package main

import (
	"fmt"
)

/*
แสดงการอ่านค่าจาก channel ที่ถูกปิดแล้ว เมื่ออ่านค่าจาก channel ที่ปิดแล้ว จะได้รับ zero value ของประเภทข้อมูลที่ channel เก็บ และตัวแปร ok จะเป็น false ซึ่งบ่งบอกว่า channel ถูกปิดแล้ว
*/

func main() {
	ch := make(chan int) // สร้าง channel ที่เก็บข้อมูลประเภท int
	close(ch)            // ปิด channel

	val, ok := <-ch // อ่านค่าจาก channel ที่ปิดแล้ว
	// val จะได้รับ zero value ของ int คือ 0
	// ok จะเป็น false เพราะ channel ถูกปิดแล้ว

	if !ok {
		fmt.Printf("Value %v is returned. the channel is closed\n", val)
		// แสดงข้อความและค่าที่ได้รับจาก channel เมื่อ channel ถูกปิด
	}
}

/*
โปรแกรมจะแสดงข้อความว่าได้รับค่า 0 (ซึ่งเป็น zero value ของ int) จาก channel และ channel ถูกปิดแล้ว เนื่องจากตัวแปร ok มีค่าเป็น false
*/
