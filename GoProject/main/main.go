package main

import "fmt"

/*
ในภาษา Go ไม่มีการสืบทอดแบบ IS-A (inheritance) แต่เราสามารถใช้การฝังแบบ (embedding) เพื่อให้ได้ฟีเจอร์ที่คล้ายกัน โดยการฝังประเภท (type) หนึ่งไว้ในอีกประเภทหนึ่งโดยไม่ต้องระบุชื่อตัวแปร

ในตัวอย่างโปรแกรม เรามี struct ชื่อ `Person` ที่มีฟิลด์และเมธอดของตัวเอง และเรามี struct ชื่อ `Student` ที่ฝัง `Person` ไว้ข้างใน ทำให้ `Student` สามารถเข้าถึงฟิลด์และเมธอดของ `Person` ได้โดยตรง เหมือนกับการสืบทอด แต่ความสัมพันธ์เป็นแบบ HAS-A แทนที่จะเป็น IS-A

การฝังแบบนี้เรียกว่าการส่งเสริมเมธอด (method promotion) ซึ่งทำให้เมธอดของประเภทที่ถูกฝังสามารถเรียกใช้ได้โดยตรงบนประเภทที่ฝังมัน นอกจากนี้ เรายังสามารถเรียกใช้เมธอดของประเภทที่ถูกฝังโดยระบุชื่อประเภทนั้นก่อนเรียกใช้เมธอดก็ได้

*/

type Person struct {
    id int
    firstName string
    lastName string
}

func (p Person) getFullName() string {
    return p.firstName + " " + p.lastName
}

type Student struct {
    Person
    marks []float32
}

func (s Student) getTotalMarks() (total float32) {
    for _, val := range s.marks {
        total += val
    }
    return
}

func main() {
	p := Person{101, "Prithvi", "Singh"}
    marks := []float32{10.2, 23.3, 19.5}
    s := Student{p, marks}

    fmt.Println(s.getFullName())
    fmt.Println(s.getTotalMarks())
}
/* 

*/