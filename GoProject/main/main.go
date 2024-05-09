package main

import (
	"fmt"
	"math"
)

/*

ในภาษา Go เราสามารถฝัง (embed) interface ลงใน struct ได้ เมื่อเราฝัง interface ลงใน struct เราจะสามารถเรียกใช้ method ทั้งหมดของ interface ผ่าน struct ได้โดยตรง

ข้อควรระวังคือ เมื่อเราสร้าง value ของ struct ที่มีการฝัง interface เราจำเป็นต้องส่งค่า implementation ที่ถูกต้องของ interface เข้าไปด้วย ไม่เช่นนั้นโปรแกรมจะเกิด panic เมื่อมีการเรียกใช้ method ของ interface

การฝัง interface ลงใน struct ช่วยให้เราสามารถเขียนโค้ดได้อย่างยืดหยุ่นและมีประสิทธิภาพมากขึ้น โดยเราสามารถใช้ struct เดียวกันกับ implementation ที่แตกต่างกันของ interface ได้ ทำให้โค้ดของเรามีความ reusable มากขึ้น

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

type MyShape struct {
    Shape
}

func main() {
    r := Rectangle{Width: 5, Height: 6}
    c := Circle{Radius: 7}

    shapes := []MyShape{
        {
            Shape: r,
        },
        {
            Shape: c,
        },
    }

    for _, s := range shapes {
        fmt.Println(s.Area())
    }
    
}

/* 
1. เราสร้าง interface `Shape` ที่มี method `Area() float64`
2. เราสร้าง struct `Rectangle` และ `Circle` ที่ implement method `Area() float64` ตาม interface `Shape`
3. เราสร้าง struct `MyShape` ที่ฝัง interface `Shape` เข้าไป
4. ในฟังก์ชัน `main()` เราสร้าง value ของ `Rectangle` และ `Circle`
5. เราสร้าง slice ของ `MyShape` และเก็บ value ของ `Rectangle` และ `Circle` เข้าไป โดยระบุ field `Shape` ให้ตรงกับ type ที่ต้องการ
6. เราวน loop ผ่าน slice ของ `MyShape` และเรียกใช้ method `Area()` ผ่าน `MyShape` ได้โดยตรง เนื่องจาก `MyShape` มีการฝัง interface `Shape` เอาไว้
*/