package main

import (
	oops "GoProject/package"
	"fmt"
)

/*
เกี่ยวกับ Receiver Function ในภาษา Go

- มี 2 แบบ คือ Pointer Receiver และ Normal Receiver
- Pointer Receiver จะแก้ไขค่าของตัวแปรต้นฉบับได้โดยตรง
- Normal Receiver ทำงานกับสำเนาของตัวแปรต้นฉบับ ต้อง return ค่ากลับไปเพื่อให้เกิดการเปลี่ยนแปลง
- เรียกใช้งาน Receiver Function ด้วย Dot Notation
*/

func main() {
	var a oops.Int = 10 // ประกาศตัวแปร a เป็นชนิด oops.Int และกำหนดค่าเริ่มต้น 10

	a.Add_One()    // เรียกใช้ Add_One เพื่อเพิ่มค่า a ขึ้น 1 เป็น 11
	fmt.Println(a) // แสดงค่า a ซึ่งขณะนี้เท่ากับ 11

	fmt.Println(a.Double()) // เรียกใช้ Double เพื่อคูณ a ด้วย 2 และแสดงผลทันที โดยค่า a ยังคงเป็น 11 ไม่เปลี่ยนแปลง
}

/*
สรุปคือ Pointer Receiver เหมาะกับการแก้ไขค่าตัวแปรโดยตรง ส่วน Normal Receiver ใช้เมื่อต้องการคืนค่าใหม่โดยไม่เปลี่ยนตัวแปรต้นฉบับ
*/
