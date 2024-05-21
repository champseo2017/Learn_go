package main

import (
	"fmt"
)

/*
แสดงการใช้ range clause ในการวนลูปอ่านค่าจาก channel โดยอัตโนมัติจนกว่า channel จะถูกปิด ซึ่งเป็นวิธีที่สะดวกและกระชับในการอ่านค่าจาก channel
*/

func main() {
	ch := make(chan int) // สร้าง channel ที่เก็บข้อมูลประเภท int

	go func() { // สร้าง goroutine เพื่อส่งข้อมูลเข้า channel
		for i := 0; i < 10; i++ {
			ch <- i // ส่งค่า i เข้า channel
		}
		close(ch) // ปิด channel เมื่อส่งข้อมูลครบแล้ว
	}()

	for val := range ch {
		// ใช้ range clause ในการวนลูปอ่านค่าจาก channel
		// ในแต่ละรอบของลูป:
		//   - อ่านค่าจาก channel และกำหนดให้กับตัวแปร val
		//   - เมื่อ channel ถูกปิด ลูปจะสิ้นสุดโดยอัตโนมัติ
		fmt.Println(val) // แสดงค่าที่อ่านได้จาก channel
	}
}

/*
เราใช้ range clause ในการวนลูปอ่านค่าจาก channel แทนการใช้ for-loop ธรรมดาหรือการประกาศตัวแปรในลูป
ในแต่ละรอบของลูป ค่าที่อ่านได้จาก channel จะถูกกำหนดให้กับตัวแปร val โดยอัตโนมัติ เมื่อ channel ถูกปิด ลูปจะสิ้นสุดโดยอัตโนมัติ โดยไม่ต้องเขียนโค้ดเพื่อตรวจสอบว่า channel ถูกปิดหรือไม่
การใช้ range clause กับ channel ช่วยให้โค้ดกระชับและอ่านง่าย เนื่องจากไม่ต้องเขียนโค้ดเพื่อจัดการการออกจากลูปเมื่อ channel ถูกปิด ทำให้โค้ดสะอาดและเข้าใจง่ายขึ้น
*/
