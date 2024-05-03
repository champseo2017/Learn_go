package main

import "fmt"

/*
1. Struct เป็น user-defined type ที่ใช้จัดกลุ่มข้อมูลที่เกี่ยวข้องกันเป็นหน่วยเดียว
2. ประกาศ struct ด้วยคีย์เวิร์ด `type` ตามด้วยชื่อและคีย์เวิร์ด `struct` และกำหนดฟิลด์และประเภทข้อมูลภายในวงเล็บปีกกา
3. สร้าง instance ของ struct ด้วย struct literal และเข้าถึงฟิลด์ด้วยเครื่องหมายจุด (`.`)
4. Struct เป็น value type เมื่อกำหนดค่าหรือส่งเป็นพารามิเตอร์ จะมีการสร้างสำเนาใหม่
5. เปรียบเทียบความเท่ากันของ struct ด้วยเครื่องหมาย `==` โดยฟิลด์ทั้งหมดต้องเท่ากัน
6. Struct และฟิลด์ที่ขึ้นต้นด้วยตัวพิมพ์ใหญ่จะถูก export ไปยังแพ็คเกจอื่นได้
7. ใช้ pointer (`&`) เพื่อชี้ไปยัง struct และเข้าถึงฟิลด์ผ่าน pointer โดยไม่ต้อง dereference
*/
// ประกาศ struct ชื่อ Person
type Person struct {
	Name string
	Age int
}

func main() {
	 // สร้าง instance ของ Person
	 p1 := Person{Name: "John", Age: 30 }
	 fmt.Println(p1) // Output: {John 30}
	 // สร้าง pointer ชี้ไปยัง p1
	 p2 := &p1
	 fmt.Println(p2) // Output: &{John 30}
	 // เข้าถึงฟิลด์ของ struct ผ่านทาง instance
	 fmt.Println(p1.Name) // Output: John
	 // เข้าถึงฟิลด์ของ struct ผ่านทาง pointer
	 fmt.Println(p2.Age) // Output: 30
	 // กำหนดค่า struct ให้กับตัวแปรอื่น (สร้างสำเนา)
	 p3 := p1
	 p3.Name = "Alice"
	 fmt.Println(p1) // Output: {John 30}
	 fmt.Println(p3) // Output: {Alice 30}

	 // เปรียบเทียบความเท่ากันของ struct
	 p4 := Person{Name: "John", Age: 30}
	 fmt.Println(p1 == p4) // Output: true


}
/* 

*/