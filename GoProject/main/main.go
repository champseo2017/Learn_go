package main

import "fmt"

/*
พิจารณาโปรแกรมจากคำถามข้อ 1 ผลลัพธ์จะเป็นอย่างไรหากเราเรียกใช้เมธอด setName() บนแอดเดรสของประเภท person
ถ้าเราเรียกใช้เมธอด setName() บนแอดเดรสของ person เช่น (&p).setName("Abc") ผลลัพธ์จะเป็น {101 Abc} เพราะการเปลี่ยนแปลงที่ทำภายในเมธอดจะส่งผลต่อค่าของ p ในฟังก์ชัน main
*/
type person struct {
    id   int
    name string
}

func (p *person) setName(name string) {
    p.name = name
}


func main() {
	p := person{101, "Xyz"}
    p.setName("Abc")
    fmt.Println(p)
}
/* 

*/