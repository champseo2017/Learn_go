package main

import (
	"fmt"
)

/*
Interface ในภาษา Go เป็นเหมือนกับชุดของ method signatures (ชื่อ method, parameter, และค่าที่ return) ที่ใช้เพื่อบอกว่า types ต่างๆ มีความคล้ายคลึงกันในเชิงแนวคิด
ยกตัวอย่างเช่น ในโค้ดด้านล่างนี้ เรามี struct สองตัวคือ Square และ Rectangle ซึ่งทั้งสองตัวนี้ implement interface ชื่อ Shape
*/

type Shape interface {
	area() float64
}

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	square_area := Square{5}
	rectangle_area := Rectangle{4, 6}

	fmt.Println("Area of square:", square_area.area())
	fmt.Println("Area of rectangle:", rectangle_area.area())
}

/*
เนื่องจากทั้ง Square และ Rectangle ต่างก็ implement interface Shape ดังนั้นตัวแปรที่มี type เป็น Shape สามารถเก็บ object ของ Square หรือ Rectangle ได้
เมื่อเราเรียกใช้ method area() ผ่านตัวแปร square_area มันจะไปเรียก method area() ของ struct Square
และเมื่อเราเรียก area() ผ่าน rectangle_area มันจะไปเรียก area() ของ Rectangle แทน
ตรงนี้แสดงให้เห็นแนวคิดของ polymorphism คือเราสามารถใช้ type เดียวกัน (Shape) เพื่อเรียกใช้ฟังก์ชันที่มี behavior คล้ายกันได้ แม้ว่าจริงๆ แล้วมันจะถูก implement โดย struct ที่แตกต่างกัน
สรุปสั้นๆ คือ interface ช่วยให้เราเขียนโค้ดที่ยืดหยุ่นและ reuse ได้มากขึ้น โดยอาศัยหลักการของ polymorphism
*/
