package main

import (
	"fmt"
)

/*
ในภาษา Go การส่ง slice เข้าไปใน function แทนการส่ง address ของ array โดยตรงถือเป็นแนวปฏิบัติที่ดีกว่า เพราะ slice มีประสิทธิภาพและความยืดหยุ่นมากกว่า array และการแก้ไขค่าใน slice ก็ยังส่งผลต่อ underlying array เดิมอยู่

การเปลี่ยนแปลงค่าใน slice ภายใน modify function ส่งผลต่อ arr ใน main function ด้วย เนื่องจาก slice ที่ส่งเข้าไปใน modify มี underlying data structure เป็น arr เดิม
*/

func modify(arr []int) { // function ที่รับ slice ของ int เป็น parameter
	arr[0] = 100 // เปลี่ยนค่า element ที่ 0 ของ slice เป็น 100
}

func main() {
	arr := [5]int{10, 20, 30, 40, 50}  // ประกาศและกำหนดค่าเริ่มต้นให้กับ array arr
	modify(arr[:])                     // สร้าง slice จาก arr ด้วย arr[:] และส่งเข้าไปใน modify function
	fmt.Println("Slice Elements", arr) // แสดงค่า element ของ arr
}
