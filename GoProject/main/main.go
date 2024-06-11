package main

import (
	oops "GoProject/package"
)

/*
Go สามารถจัดกลุ่มข้อมูลต่างๆ เข้าด้วยกันได้โดยใช้ struct ซึ่งเหมือนกับ class ในภาษาอื่นๆ ภายใน struct เราสามารถประกาศ properties และ methods ได้
*/

func main() {
	p1 := oops.Player{
		Name:   "Udit",
		Age:    29,
		Sports: "football",
	}
	p1.Print_Details()
}

/*
ตัวอย่างนี้ Player เป็น struct มี properties Name, Age, Sports และ method Print_Details() ที่แสดงข้อมูลเหล่านี้

ด้วย struct และ receiver functions เราสามารถใช้หลักการ encapsulation และ abstraction ได้ในการพัฒนาแบบ OOP
*/
