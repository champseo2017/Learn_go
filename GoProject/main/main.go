package main

import (
	"GoProject/module" // นำเข้าชื่อของ module ที่มีสตรักเจอร์ Person และ Address
	"fmt"
)

/*

 */

func main() {
	// สร้างออบเจ็กต์ Person
	p := module.Person{
		Name: "Udit",
		Age:  30,
		Address: module.Address{
			HouseNo: "A8",
			City:    "Delhi",
		},
	}

	// แสดงข้อมูลของออบเจ็กต์ Person
	fmt.Println(p)
}

/*
ในฟังก์ชัน `main()` คุณสามารถสร้างออบเจ็กต์ `Person` โดยใช้ชื่อของ package หรือ module ที่นำเข้ามา ตามด้วยชื่อของสตรักเจอร์ เช่น `module_name.Person` เพื่อกำหนดค่าของสนามต่างๆ ในสตรักเจอร์ Person
*/
