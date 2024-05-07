package main

import "fmt"

/*
Method set คือชุดของ method ที่ทำให้ type หนึ่งๆ สามารถ implement interface ได้โดยปริยาย

ที่ St{} เกิด error เพราะ St เป็น value type ที่ไม่ได้ implement Itr แต่ \*St ซึ่งเป็น pointer type ต่างหากที่ implement Itr

เมื่อเรากำหนด method ให้เป็น pointer receiver เช่น func (s \*St) m1() แสดงว่า method นั้นจะถูก implement โดย pointer type (\*St) ไม่ใช่ value type (St)

ดังนั้น เมื่อเรากำหนดค่า St{} ซึ่งเป็น value ให้กับตัวแปร interface จึงเกิด error เพราะ St ไม่ได้ implement method ของ interface แต่ถ้ากำหนดเป็น &St{} ซึ่งเป็น pointer จะไม่เกิด error เพราะ \*St implement method ของ interface แล้ว

สรุปคือ:
- ถ้า method เป็น value receiver (func (s St) m1()) จะถูก implement โดย value type (St)
- ถ้า method เป็น pointer receiver (func (s \*St) m1()) จะถูก implement โดย pointer type (\*St)

เราต้องกำหนดค่าให้ตรงกับ type ที่ implement interface เท่านั้น ถึงจะไม่เกิด error
*/

type Itr interface {
    m1()
    m2()
}

type St struct {}

func (s *St) m1() {}
func (s *St) m2() {}

func main() {
   i := Itr(&St{}) // ถูกต้อง เพราะ *St implement Itr
   // i := Itr(St{}) ผิด เพราะ St ไม่ได้ implement Itr
   fmt.Println(i)
}
