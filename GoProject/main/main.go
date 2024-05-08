package main

import "fmt"

/*
ถ้า type T implement interface I โดยกำหนด pointer receiver method แสดงว่าเฉพาะ *T เท่านั้นที่จะ implement interface I ไม่ใช่ T ดังนั้นเราสามารถกำหนดเฉพาะ address ของ T ให้กับตัวแปร interface เท่านั้น
แต่ถ้า type T implement interface I โดยกำหนด value receiver method แสดงว่าทั้ง T และ *T จะ implement interface I ดังนั้นเราสามารถกำหนดได้ทั้ง value และ address ของ T ให้กับตัวแปร interface
*/

// กรณี pointer receiver
type IPrint interface {
    Print()
}

type Student struct {
    name string
}

func (s *Student) Print() {
    fmt.Println("Name:", s.name)
}

func main() {
  i := IPrint(&Student{"John"})
  fmt.Printf("%s", i)
}
