package main

import (
	"fmt"
	"math"
)

/*

Polymorphism (พอลิมอร์ฟิซึม) เป็นแนวคิดในการเขียนโปรแกรมเชิงวัตถุ ที่ออบเจกต์สามารถมีหลายรูปแบบหรือพฤติกรรมที่แตกต่างกันได้ แม้ว่าจะถูกอ้างอิงผ่านตัวแปรหรือพารามิเตอร์ชนิดเดียวกัน

เปรียบเหมือนกับรูปร่างต่างๆ เช่น สี่เหลี่ยม วงกลม และสามเหลี่ยม ซึ่งถือเป็น "รูปร่าง" เหมือนกัน แต่มีวิธีคำนวณพื้นที่แตกต่างกัน อย่างไรก็ตาม เราสามารถใช้คำสั่งเดียวกันเพื่อให้ทุกรูปร่างคำนวณพื้นที่ของตัวเองได้

ในภาษา Go เราใช้ interface เพื่อทำ Polymorphism โดยกำหนด method ที่จำเป็นใน interface และให้แต่ละ type ไปเขียน method เหล่านั้น เมื่อเราสร้างฟังก์ชันที่รับ interface เป็นพารามิเตอร์ เราสามารถส่งออบเจกต์ของ type ต่างๆ ที่ implement interface นั้นเข้าไปได้

*/

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}


func main() {
    shapes := []Shape{
        Rectangle{Width: 5, Height: 6},
        Circle{Radius: 7},
    }

    for _, s := range shapes {
        fmt.Println(s.Area())
    }
}

/* 

*/