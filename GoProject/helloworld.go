package main

import "fmt"

/*
ในภาษา Go ฟังก์ชันสามารถคืนค่าหลายค่าได้โดยใช้คำสั่ง return ชนิดของค่าที่คืนจะขึ้นอยู่กับชนิดของพารามิเตอร์ที่กำหนดในรายการพารามิเตอร์ ดังที่แสดงด้านล่าง
*/

// ฟังก์ชัน testfunc คืนค่า 2 ค่าที่เป็นชนิด int
func testfunc(x, y int) (int, int) {
	return x + y, x - y
}

func main() {
	// กำหนดค่าที่คืนจากฟังก์ชันให้กับตัวแปรต่างกัน
	testVar1 , testVar2 := testfunc(10, 20)
	fmt.Printf("ผลบวก: %d", testVar1)
	fmt.Printf("\nผลลบ: %d", testVar2)
}