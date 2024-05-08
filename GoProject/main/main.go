package main

import "fmt"

/*
เมื่อเรา implement interface ด้วย value receiver เราสามารถกำหนดทั้ง value และ address ของ type ให้กับ interface ได้
*/

type Itr interface {
    m1()
    m2()
}

type St struct {

}

func (s St) m1() {
    fmt.Println("In m1 of St")
}

func (s St) m2() {
    fmt.Println("In m2 of St")
}

func main() {
   fmt.Println("Assigning value of St")
   i := Itr(St{})
   i.m1()
   i.m2()

   fmt.Println("Assigning Address of St")
   i = Itr(&St{})
   i.m1()
   i.m2()

}
