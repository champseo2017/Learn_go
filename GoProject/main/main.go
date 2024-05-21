package main

import (
	"fmt"
)

/*
เกี่ยวกับการวนลูปอ่านค่าจาก channel:

- มีสองวิธีหลักในการวนลูปอ่านค่าจาก channel:
  1. ใช้ค่าที่สองที่คืนกลับมาจาก channel เพื่อตรวจสอบว่า channel ถูกปิดหรือไม่ ถ้าเป็น false ให้ออกจากลูป
  2. ใช้ range clause ในการวนลูปอ่านค่าจาก channel โดยอัตโนมัติจนกว่า channel จะถูกปิด

- เมื่อ channel ถูกปิด การอ่านค่าจาก channel จะคืนค่า zero value ของประเภทข้อมูลที่ channel เก็บ และค่าที่สองจะเป็น false
- การวนลูปอ่านค่าจาก channel เป็นวิธีที่สะดวกในการรับข้อมูลทั้งหมดที่ส่งผ่าน channel จนกว่า channel จะถูกปิด
- ควรมีการปิด channel หลังจากส่งข้อมูลครบแล้ว เพื่อให้ goroutine ที่รอรับข้อมูลสามารถออกจากลูปได้อย่างถูกต้อง
*/

func main() {
	ch := make(chan int) // สร้าง channel ที่เก็บข้อมูลประเภท int
	go func() {          // สร้าง goroutine เพื่อส่งข้อมูลเข้า channel
		for i := 0; i < 10; i++ {
			ch <- i // ส่งค่า i เข้า channel
		}
		close(ch) // ปิด channel เมื่อส่งข้อมูลครบแล้ว
	}()

	// for { // วนลูปอ่านค่าจาก channel
	// 	val, ok := <-ch // อ่านค่าจาก channel
	// 	if !ok {        // ถ้า ok เป็น false แสดงว่า channel ถูกปิดแล้ว
	// 		break // ออกจากลูป
	// 	}
	// 	fmt.Println(val) // แสดงค่าที่อ่านได้จาก channel
	// }
	for val := range ch { // ใช้ range clause ในการวนลูปอ่านค่าจาก channel
		fmt.Println(val) // แสดงค่าที่อ่านได้จาก channel
	}
}

/*

 */
