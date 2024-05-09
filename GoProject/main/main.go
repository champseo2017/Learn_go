package main

import (
	"fmt"
)

/*



 */

type I interface {
    m1()
}

type T struct {
}

func (t *T) m1() {
    fmt.Println("M1 of T")
}

func main() {
    var i I
    t := T{}
    // i = t // กำหนดค่า i = t ซึ่งจะเกิด error เพราะ T ไม่ได้ implement method m1() แต่ *T ต่างหากที่ implement
    i = &t // เพราะว่า i เป็น interface I ซึ่งมี method m1() และ *T (pointer ของ T) นั้น implement method m1() ตาม interface I ดังนั้นเราจึงต้องกำหนดค่าให้ i ด้วย &t ซึ่งเป็นการส่ง address ของ t (หรือ pointer ไปยัง t) เพื่อให้ i ชี้ไปยัง instance ของ type ที่ implement interface I นั่นเอง
    i.m1()
}

/* 

*/