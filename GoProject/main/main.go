package main

import (
	"fmt"
)

/*
ในภาษา Go เราสามารถส่ง pointer array เป็น argument ให้กับ function ได้ เมื่อมีการแก้ไขค่าใน element ของ pointer array ภายใน function ที่ถูกเรียก การเปลี่ยนแปลงนั้นจะสะท้อนกลับไปยัง function ที่เรียกด้วย
ในโปรแกรม 10.13 มีการส่ง address ของ array เข้าไปใน modify function และทำการเปลี่ยนแปลงค่าของ element ที่ 0 ของ array เป็น 100 ภายใน modify function เมื่อมีการแสดงค่า array ใน main function จะเห็นว่า element ที่ 0 ของ array มีค่าเป็น 100 ซึ่งแสดงให้เห็นว่าการแก้ไขใน function ที่ถูกเรียกส่งผลต่อ function ที่เรียกด้วย
อย่างไรก็ตาม ไม่ถือเป็นแนวปฏิบัติที่ดีในการส่ง address ของ array ไปยัง function โดยตรง แต่ควรสร้าง slice จาก array ด้วย index เริ่มต้นและสิ้นสุดเป็นค่าเริ่มต้น (0 ถึง length ของ array - 1) และส่ง slice นั้นเข้าไปใน function แทน เนื่องจาก slice ที่สร้างจาก array ด้วย index เริ่มต้นจะมี underlying data structure เป็น array เดิม ดังนั้นการแก้ไขค่าใน slice จะส่งผลต่อ array ต้นฉบับด้วย
*/

func modify(arr *[5]int) { // function ที่รับ pointer ของ array ขนาด 5 ตัว
	arr[0] = 100 // เปลี่ยนค่า element ที่ 0 ของ array เป็น 100
}

func main() {
	arr := [5]int{10, 20, 30, 40, 50}  // ประกาศและกำหนดค่าเริ่มต้นให้กับ array ของ int ขนาด 5 ตัว
	modify(&arr)                       // ส่ง address ของ arr เข้าไปใน modify function
	fmt.Println("Array Elements", arr) // แสดงค่า element ของ arr
}
