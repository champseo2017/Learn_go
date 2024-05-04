package main

/*
1. Methods คือฟังก์ชันที่เชื่อมโยงกับ Type (ชนิดข้อมูล) เฉพาะ
2. Methods ประกาศด้วยการเพิ่ม Receiver ระหว่างคีย์เวิร์ด `func` และชื่อ Method
3. Receiver ระบุ Type ที่ Method เชื่อมโยงด้วย และสามารถเข้าถึงฟิลด์ของ Type นั้นได้
4. Methods สามารถมีพารามิเตอร์และคืนค่าได้เหมือนฟังก์ชันทั่วไป
5. ชื่อ Receiver ควรสั้นและใช้ชื่อเดียวกันสำหรับทุก Method เพื่อความสม่ำเสมอ
6. Methods และ Type ของมันต้องถูกประกาศในแพ็คเกจเดียวกัน ไม่สามารถประกาศ Method ในแพ็คเกจที่แตกต่างจาก Type ได้

Methods ช่วยให้เราสามารถเขียนโค้ดแบบเชิงวัตถุ (Object-Oriented) ในภาษา Go ได้ โดยกำหนดพฤติกรรมเฉพาะให้กับ Type ต่างๆ และทำให้โค้ดมีความเป็นโมดูลาร์ ง่ายต่อการอ่านและบำรุงรักษา

func (receiverName receiverType) methodName(param1 paramType) (returnType1) {
    // ...
}
*/

type rectangle struct {
	length float32
	width float32
}

func (r rectangle) area() float32 {
	return r.length * r.width
}

func (rect rectangle) perimeter() float32 {
	return 2 * (rect.length + rect.width)
}

func main() {
	
}
/* 

*/