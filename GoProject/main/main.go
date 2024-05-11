package main

import (
	"fmt"
)

/*
1. การเข้าถึง element ของ pointer array สามารถทำได้โดยไม่ต้องใช้ `*` นำหน้า เช่น `arr1[0]` แทนที่จะเป็น `(*arr1)[0]` ซึ่งทำให้โค้ดอ่านง่ายและชัดเจนขึ้น

2. เมื่อประกาศ pointer array ด้วย `var` แต่ไม่ได้กำหนดที่อยู่ของ array ให้ชี้ไป และพยายามเข้าถึง element ของ pointer array โปรแกรมจะเกิด panic เนื่องจากไม่มี array ที่อยู่เบื้องหลัง pointer

3. สามารถสร้าง pointer array โดยใช้ `:=` ได้ โดยประกาศ array ด้วย `var` และใช้ `:=` เพื่อให้ type ของ pointer array ถูก infer จากด้านขวาของ `:=`
*/

func main() {
	var arr1 *[5]int
	var arr2 = [5]int{1, 2, 3, 4, 5}
	arr1 = &arr2 // ใช้ := เพื่อให้ arr1 เป็น pointer ชี้ไปยัง arr2

	fmt.Println(arr1[0]) // เข้าถึง element ที่ 0 ของ arr1 โดยไม่ใช้ *
}
