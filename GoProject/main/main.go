package main

import (
	"fmt"
)

/*
สรุปสั้นๆ เกี่ยวกับการใช้ Pointer Receiver ใน Method ของภาษา Go:

1. Pointer Receiver ใช้เพื่อรับค่าเป็น Pointer (Memory Address) ของ Receiver Type
2. การประกาศ Method ด้วย Pointer Receiver ทำโดยเพิ่มเครื่องหมายดอกจัน (`*`) ไว้ด้านหน้า Receiver Type
3. เมื่อใช้ Pointer Receiver ใน Method เราสามารถเข้าถึงและแก้ไขค่าของ Receiver ได้โดยตรง
4. การเปลี่ยนแปลงค่าใน Method ที่ใช้ Pointer Receiver จะส่งผลต่อค่าเดิมของ Receiver ในฟังก์ชันที่เรียกใช้ Method
5. Pointer Receiver มีประโยชน์อย่างมากในการเขียนโค้ดที่ต้องการปรับเปลี่ยนค่าของ struct หรือ type ต่างๆ

ดังนั้น การเลือกใช้ Pointer Receiver หรือ Value Receiver ใน Method ขึ้นอยู่กับวัตถุประสงค์ของเรา หากต้องการให้การเปลี่ยนแปลงค่าใน Method ส่งผลต่อค่าเดิมของ Receiver ให้ใช้ Pointer Receiver แต่หากต้องการแค่อ่านค่าจาก Receiver โดยไม่เปลี่ยนแปลงค่าเดิม ให้ใช้ Value Receiver

การประกาศ Method ด้วย Pointer Receiver
func (receiverName *receiverType) methodName(param1 paramType) (returnType1) {
    // ...
}
*/

type rectangle struct {
	length float32
	width float32
}

func (r *rectangle) increaseLength(a float32) {
	r.length = r.length + a
}

func main() {
	rect := rectangle{12.3, 21.45}
	rect.increaseLength(3.4)
	fmt.Println(rect)
}
/* 

*/