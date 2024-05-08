package main

import "fmt"

/*
การกำหนด address ให้กับตัวแปร interface มีประโยชน์ในแง่ของการเข้าถึงและแก้ไขข้อมูลใน value ของ type ผ่าน method ที่มี pointer receiver

การกำหนด address ให้กับ interface จะมีประโยชน์เฉพาะกับ pointer receiver method เท่านั้น ที่จะสามารถเปลี่ยนแปลงข้อมูลใน value ของ type ได้โดยตรง
*/

type IPrint interface {
    Print()
}

type Student struct {
    name string
}

func (s *Student) Print() {
    fmt.Printf("Name: %s\n", s.name)
}

func (s *Student) ChangeName(newName string) {
    s.name = newName
}

func main() {
    s := Student{"John"}
    
    i := IPrint(&s)
    i.Print() // Output: Name: John
    
    s.ChangeName("Jane")
    i.Print() // Output: Name: Jane
    
    fmt.Printf("i: %v\n", i) // Output: i: &{Jane}
    fmt.Printf("i: %v\n", s)
}
