package main

import "fmt"

/*
คำสั่งควบคุมใน Go ใช้เพื่อกำหนดทิศทางการทำงานของโปรแกรมตามเงื่อนไขที่กำหนด ประกอบด้วย:

1. if...else และ if...else if...else: ทำงานตามเงื่อนไขที่เป็นจริง
2. switch: เลือกทำงานใน case ที่ตรงกับเงื่อนไข
3. for loop: ทำซ้ำโค้ดตามเงื่อนไข เป็นลูปเดียวใน Go
4. break: หยุดลูปหรือคำสั่งปัจจุบัน
5. goto: ข้ามไปทำงานที่คำสั่งที่ระบุด้วยป้ายกำกับ
6. continue: ข้ามรอบปัจจุบันของลูป แล้วเริ่มรอบใหม่
7. infinite loop: ลูปที่ทำงานไม่รู้จบ
8. fallthrough: ทำงานใน case ถัดไปใน switch แม้ไม่ตรงเงื่อนไข

*/

func main() {
	switch number := 200; {
	case number < 300:
		fmt.Printf("%d is less than 300\n", number)
		fallthrough
	case number > 100:
		fmt.Printf("%d is greater than 100", number)
	}
}

/*

*/