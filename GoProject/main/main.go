package main

import (
	"fmt"
)

/*
Pointer เป็นตัวแปรที่เก็บ memory address ของตัวแปรอื่น เมื่อค่าของตัวแปรที่ pointer ชี้ไปถูกเปลี่ยนแปลง ค่าของ pointer ก็จะเปลี่ยนแปลงตามไปด้วย

Pointer มีประโยชน์มากเมื่อเราต้องการส่งตัวแปรไปยังฟังก์ชัน เพราะเมื่อมีการเปลี่ยนแปลงค่าของตัวแปรภายในฟังก์ชัน ค่าของตัวแปรในฟังก์ชันที่เรียกใช้ก็จะเปลี่ยนแปลงตามไปด้วย

ในบทนี้ เราจะเรียนรู้เกี่ยวกับ pointer ในภาษา Go ว่ามีประโยชน์อย่างไร วิธีการประกาศตัวแปร pointer และชนิดของ pointer นอกจากนี้ เราจะเรียนรู้ว่าทำไมเราจึงไม่สามารถทำ arithmetic operation บน pointer ได้ในภาษา Go

*/

func main() {
    x := 10
    p := &x

    fmt.Println("Value of x:", x)
    fmt.Println("Address of x:", &x)
    fmt.Println("Value of p:", p)

    *p = 20 // * คือ dereference operator ที่ใช้เข้าถึงค่าที่ memory address ที่ pointer ชี้ไป

    fmt.Println("New value of x:", x)
}
/* 
- เราประกาศตัวแปร `x` และกำหนดค่าเป็น 10
- เราประกาศตัวแปร `p` เป็น pointer ที่ชี้ไปยัง `x` โดยใช้ `&x` (& คือ address-of operator ที่ return memory address ของตัวแปร)
- เราแสดงค่าของ `x`, memory address ของ `x`, และค่าของ `p` (ซึ่งก็คือ memory address ของ `x`)
- เราเปลี่ยนค่าที่ memory address ที่ `p` ชี้ไปเป็น 20 โดยใช้ `*p` (* คือ dereference operator ที่ใช้เข้าถึงค่าที่ memory address ที่ pointer ชี้ไป)
- เราแสดงค่าใหม่ของ `x` ซึ่งถูกเปลี่ยนเป็น 20 ผ่านการเปลี่ยนค่าผ่าน pointer `p`
*/