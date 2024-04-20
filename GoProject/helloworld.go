package main

import "fmt"

/*
ในการสร้าง slice ใน Go นั้น slice จะประกอบด้วย 3 ส่วนหลักๆ ได้แก่ pointer, length และ capacity โดย pointer จะเป็นตัวชี้ไปยัง underlying array ที่ slice นั้นๆ อ้างอิงถึง ส่วน length จะเป็นความยาวของ segment ของ array ที่ slice นั้นๆ ครอบคลุม และ capacity จะเป็นขนาดสูงสุดที่ segment นั้นสามารถขยายได้
*/

func main() {
	// สร้าง underlying array
	arr := [5]int{1, 2, 3, 4, 5}
	// สร้าง slice ที่อ้างอิงถึง arr โดยเริ่มต้นที่ index 1 และสิ้นสุดที่ index 3 (ไม่รวม index 3)
	slice := arr[1:3]
	// แสดงค่าของ slice, len(slice) และ cap(slice)
	fmt.Printf("slice = %v, len = %d, cap = %d\n", slice, len(slice), cap(slice))
	// เปลี่ยนแปลงค่าใน slice
	slice[0] = 10
	// ค่าใน underlying array จะเปลี่ยนแปลงตามไปด้วย
	fmt.Printf("arr = %v\n", arr)
}

/* 
การใช้ %v เป็นวิธีที่สะดวกในการแสดงค่าของตัวแปรโดยไม่ต้องกังวลเกี่ยวกับประเภทข้อมูล เพราะ Printf จะเลือกรูปแบบการแสดงผลที่เหมาะสมให้โดยอัตโนมัติตามประเภทของตัวแปรที่ถูกส่งเข้ามา

นอกจาก %v แล้ว ยังมี verbs อื่นๆ เช่น %d สำหรับแสดงค่าตัวเลขจำนวนเต็ม, %f สำหรับแสดงค่าตัวเลขทศนิยม, %s สำหรับแสดงค่า string เป็นต้น ซึ่งเราสามารถเลือกใช้ให้เหมาะสมกับความต้องการในการแสดงผลข้อมูลได้

อธิบาย:
- เริ่มต้นโดยสร้าง underlying array `arr` ที่มี 5 elements
- สร้าง slice `slice` โดยอ้างอิงถึง `arr` เริ่มต้นที่ index 1 และสิ้นสุดที่ index 3 (ไม่รวม index 3) ดังนั้น `slice` จะมีค่าเป็น `[2 3]`
- `len(slice)` จะเท่ากับ 2 เนื่องจาก slice มี 2 elements
- `cap(slice)` จะเท่ากับ 4 เนื่องจาก slice สามารถขยายได้ถึง index สุดท้ายของ underlying array
- เมื่อเปลี่ยนแปลงค่าใน `slice` เช่น `slice[0] = 10` ค่าใน underlying array `arr` ก็จะเปลี่ยนแปลงตามไปด้วย เนื่องจาก slice และ underlying array ใช้ memory เดียวกัน
*/