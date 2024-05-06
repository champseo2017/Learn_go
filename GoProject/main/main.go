package main

import "fmt"

/*
ภาษา Go ไม่อนุญาตให้ใช้พอลิมอร์ฟิซึม (polymorphism) กับประเภทที่ถูกฝัง (embedded types) แต่สามารถทำได้โดยใช้อินเทอร์เฟซ (interface)
*/

type Person struct {
    id int
    name string
}

type Student struct {
    Person
    marks []float32
}

func processPerson(person Person) {
    fmt.Printf("Processing person: %d, %s\n", person.id, person.name)
}

func main() {
	s := Student{Person{101, "Prithvi"}, []float32{9.3, 3.4, 6.7}}
    processPerson(s.Person)
}
/* 
ซึ่งแสดงให้เห็นว่าเราสามารถส่งค่า Person ที่เป็นส่วนหนึ่งของ Student ให้กับฟังก์ชัน processPerson ได้ แม้ว่าเราจะไม่สามารถส่งค่า Student โดยตรงให้กับฟังก์ชันนี้ก็ตาม
*/