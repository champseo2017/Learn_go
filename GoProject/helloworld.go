package main

import "fmt"

/*
เราสามารถเพิ่ม key-value pairs ลงในแมปที่สร้างไว้แล้วได้ โดยใช้รูปแบบดังนี้:
map_name[key] = value
*/

func main() {
	// สร้างและกำหนดค่าเริ่มต้นให้กับแมป
	var my_map = map[int]string {
		1: "US",
		91: "India",
		86: "China",
		44: "UK",
	}
	fmt.Println("Original Map\n", my_map)
	// เพิ่ม key-value pairs ใหม่ลงในแมป
	my_map[39] = "Italy"
	my_map[81] = "Japan"
	fmt.Println("\nMap after adding new key-value pairs \n", my_map)
}

/* 
1. เราสร้างแมป `my_map` และกำหนดค่าเริ่มต้นให้กับแมปโดยใช้ `map[int]string` ซึ่งหมายความว่า key เป็นชนิด `int` และ value เป็นชนิด `string`
2. เราพิมพ์แมปเดิมออกมาโดยใช้ `fmt.Println("Original Map\n", my_map)`
3. เราเพิ่ม key-value pairs ใหม่ลงในแมปโดยใช้ `my_map[39] = "Italy"` และ `my_map[81] = "Japan"` ซึ่งจะเพิ่ม key `39` กับ value `"Italy"` และ key `81` กับ value `"Japan"` ลงในแมป `my_map`
4. สุดท้ายเราพิมพ์แมปที่อัปเดตแล้วออกมาโดยใช้ `fmt.Println("\nMap after adding new key-value pairs \n", my_map)`
*/