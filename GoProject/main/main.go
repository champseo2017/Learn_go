package main

import (
	"fmt"
)

/*
แก้ไขโค้ดให้โปรแกรมทำงานจนจบได้อย่างสมบูรณ์โดยการแสดง error โดยห้ามแก้ไขค่า index ของ array
*/

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	index := 5

	// ตรวจสอบว่า index อยู่ในช่วงของ array หรือไม่
	if index < 0 || index >= len(arr) {
		// ถ้า index ไม่อยู่ในช่วง ให้แสดง error และจบการทำงาน
		panic(fmt.Errorf("index out of range: %d", index))
	}

	// ถ้า index อยู่ในช่วง ให้แสดงค่าใน array ตาม index นั้น
	fmt.Println(arr[index])
}

/*

 */
