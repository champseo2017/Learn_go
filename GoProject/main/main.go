package main

import "fmt"

/*

 */

// ผลลัพธ์ของโค้ดต่อไปนี้คืออะไร
type person struct {
    id int
    name string
}

func (p person) setName(name string) {
    p.name = name
}


func main() {
	p := person{101, "Xyz"}
    p.setName("Abc")
    fmt.Println(p)
}
/* 
ผลลัพธ์จะเป็น {101 Xyz} เพราะเมธอด setName ใช้ตัวรับแบบค่า (value receiver) ดังนั้นการเปลี่ยนแปลงใดๆ ที่ทำภายในเมธอดจะไม่ส่งผลต่อค่าของ p ในฟังก์ชัน main
*/