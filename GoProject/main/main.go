package main

import (
	"fmt"
	"math"
)

/*
ในภาษา Go การใช้งานอินเทอร์เฟซเป็นแบบ Implicit Implementation หมายความว่า เมื่อมีโครงสร้างหรือประเภทที่มีเมธอดตรงตามที่ระบุในอินเทอร์เฟซ ก็ถือว่าโครงสร้างหรือประเภทนั้นได้ implement อินเทอร์เฟซนั้นแล้วโดยอัตโนมัติ โดยไม่จำเป็นต้องระบุชัดเจน (explicit)

ข้อดีของ Implicit Implementation คือ

1. สามารถ implement อินเทอร์เฟซที่เราสร้างขึ้นเองกับโครงสร้างหรือประเภทที่มาจาก third-party library ได้ แม้ว่าเราจะไม่สามารถแก้ไขโค้ดของ library นั้นได้ก็ตาม

2. เมื่อเราต้องการสร้างอินเทอร์เฟซใหม่ที่มีเมธอดไม่มาก และต้องการ implement อินเทอร์เฟซนั้นในหลายๆ โครงสร้างหรือประเภท เราไม่จำเป็นต้องไปแก้ไขโค้ดในทุกๆ ที่

ตัวอย่างเช่น เรามีอินเทอร์เฟซ `Shape` และโครงสร้าง `Rectangle` และ `Circle` ที่มีเมธอดตรงตามที่ระบุในอินเทอร์เฟซ `Shape` ถึงแม้เราจะไม่ได้ระบุชัดเจนว่าต้องการ implement อินเทอร์เฟซ `Shape` แต่ก็ถือว่า `Rectangle` และ `Circle` ได้ implement อินเทอร์เฟซ `Shape` แล้วโดยอัตโนมัติ

เราสามารถสร้างตัวแปรชนิด `Shape` และกำหนดค่าเป็น `Rectangle` หรือ `Circle` ได้ และเรียกใช้เมธอดผ่านตัวแปรชนิด `Shape` ได้ ซึ่งแสดงให้เห็นว่า Implicit Implementation ช่วยให้เราใช้งานอินเทอร์เฟซได้อย่างยืดหยุ่นและสะดวกมากขึ้น
*/

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width float64
    Height float64
}

type Circle struct {
    Radius float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}


func main() {
   s := Shape(Rectangle{Width: 5, Height: 10})
   fmt.Printf("Area of rectangle: %.2f\n", s.Area())
   fmt.Printf("Perimeter of rectangle: %.2f\n", s.Perimeter())

   s = Shape(Circle{Radius: 7})
   fmt.Printf("Area of circle: %.2f\n", s.Area())
   fmt.Printf("Perimeter of circle: %.2f\n", s.Perimeter())
}

/* 
จะเห็นว่าเราสามารถเรียกใช้เมธอด Area() และ Perimeter() ผ่านตัวแปร s ชนิด Shape ได้ ทั้งๆ ที่ s ถูกกำหนดค่าเป็น Rectangle และ Circle ซึ่งแสดงให้เห็นว่า Implicit Implementation ช่วยให้เราสามารถใช้งานอินเทอร์เฟซได้อย่างยืดหยุ่นและสะดวกมากขึ้น
*/