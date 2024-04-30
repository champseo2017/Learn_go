package main

import "fmt"

/*
Go หากเราพยายามเพิ่ม key ที่มีอยู่แล้วในแมป มันจะเป็นการเขียนทับหรืออัปเดตค่าของ key นั้นด้วยค่าใหม่แทน
*/

func main() {
	// สร้างและกำหนดค่าเริ่มต้นให้กับแมป
	var my_map = map[int]string {
		1:"US",
		91: "India",
		86: "China",
		44: "UK",
	}
	fmt.Println("Original Map\n", my_map)

	// อัปเดตค่าในแมป
	my_map[1] = "United States"
	my_map[44] = "United Kingdom"
	fmt.Println("\nAfter updating values of the map \n", my_map)
}

/* 
1. เราสร้างแมป `my_map` และกำหนดค่าเริ่มต้นให้กับแมปโดยใช้ `map[int]string` ซึ่งหมายความว่า key เป็นชนิด `int` และ value เป็นชนิด `string`
2. เราพิมพ์แมปเดิมออกมาโดยใช้ `fmt.Println("Original Map\n", my_map)`
3. เราอัปเดตค่าในแมปโดยใช้ `my_map[1] = "United States"` และ `my_map[44] = "United Kingdom"` ซึ่งจะเปลี่ยนค่าของ key `1` จาก `"US"` เป็น `"United States"` และเปลี่ยนค่าของ key `44` จาก `"UK"` เป็น `"United Kingdom"`
4. สุดท้ายเราพิมพ์แมปที่อัปเดตแล้วออกมาโดยใช้ `fmt.Println("\nAfter updating values of the map \n", my_map)`

จากผลลัพธ์จะเห็นได้ว่าเมื่อเราพยายามเพิ่ม key ที่มีอยู่แล้วในแมป (`1` และ `44`) ค่าของ key เหล่านั้นจะถูกอัปเดตเป็นค่าใหม่ (`"United States"` และ `"United Kingdom"` ตามลำดับ) แทนที่จะเพิ่ม key-value pairs ใหม่ลงในแมป

นี่เป็นวิธีการอัปเดตค่าใน key ที่มีอยู่แล้วในแมปอย่างง่ายดายโดยใช้รูปแบบ `map_name[existing_key] = new_value`
*/