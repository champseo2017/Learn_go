package main

import (
	"fmt"
)

/*
Value Receiver และ Pointer Receiver ในภาษา Go:

1. Value Receiver คือการส่งค่าของ Receiver (ตัวแปรที่เป็น Type ของ Method) เข้าไปใน Method
   - การเปลี่ยนแปลงค่าใน Method ที่ใช้ Value Receiver จะไม่ส่งผลต่อค่าเดิมของ Receiver ในฟังก์ชันที่เรียกใช้
   - Value Receiver เหมาะสำหรับการอ่านค่าจาก Receiver เท่านั้น

2. Pointer Receiver คือการส่งพอยน์เตอร์ (Memory Address) ของ Receiver เข้าไปใน Method
   - การเปลี่ยนแปลงค่าใน Method ที่ใช้ Pointer Receiver จะส่งผลต่อค่าเดิมของ Receiver ในฟังก์ชันที่เรียกใช้
   - Pointer Receiver เหมาะสำหรับการแก้ไขค่าของ Receiver หรือเมื่อต้องการให้การเปลี่ยนแปลงใน Method ส่งผลต่อค่าเดิมของ Receiver

การเลือกใช้ Value Receiver หรือ Pointer Receiver ขึ้นอยู่กับวัตถุประสงค์ของ Method ว่าต้องการแค่อ่านค่าจาก Receiver หรือต้องการแก้ไขค่าของ Receiver ด้วย ซึ่งจะช่วยให้เขียนโค้ดได้อย่างมีประสิทธิภาพและถูกต้องตามวัตถุประสงค์
*/

type rectangle struct {
	length float32
	width float32
}

func (r rectangle) increaseLength(a float32) {
	r.length = r.length + a
}

func main() {
	rect := rectangle{12.3, 21.45}
	rect.increaseLength(3.4)
	fmt.Println(rect)
}
/* 

*/