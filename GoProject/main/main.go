package main

import (
	"fmt"
)

/*

Interface สามารถใช้ภายใน struct ได้ โดยมีวิธีใช้ 2 แบบ คือ:
1. ใช้ interface เป็น embedded type ใน struct
2. ประกาศตัวแปร interface ใน struct

เราไม่สามารถส่งค่า interface โดยตรงตอนสร้าง struct ได้ เพราะ interface เป็นนามธรรม ต้องส่งเป็น concrete type ที่ implement interface แทน

เมื่อเรียกใช้เมธอดของ interface ผ่าน struct การเรียกใช้จะถูกส่งต่อไปยัง concrete type ที่อยู่ใน struct นั้น

ตัวอย่างการใช้ interface ใน struct:
1. Embedded type:
```go
type MyStruct struct {
    MyInterface
}

s := MyStruct{MyImpl{}}
s.Method1()
```

2. ประกาศตัวแปร interface:
```go
type MyStruct struct {
    Iface MyInterface
}

s := MyStruct{Iface: MyImpl{}}
s.Iface.Method1()
```

ในทั้งสองกรณี เรากำหนด concrete type (MyImpl) ที่ implement MyInterface ให้กับ struct และเรียกใช้เมธอดของ interface ผ่าน struct ซึ่งจะเรียกใช้เมธอดจาก MyImpl ที่อยู่ใน struct
*/

type MyInterface interface {
    Method1()
    Method2()
}

type MyStruct struct {
    MyInterface
}

type MyImpl struct {}

func (m MyImpl) Method1() {
    fmt.Println("Method1")
}

func (m MyImpl) Method2() {
    fmt.Println("Method2")
}

func main() {
    s := MyStruct{MyImpl{}}
    s.Method1()
    s.Method2()
}

/* 
สร้าง interface ชื่อ MyInterface ที่มีเมธอด Method1 และ Method2
สร้าง struct ชื่อ MyStruct ที่มี field เป็น MyInterface (embedded type)
สร้าง struct ชื่อ MyImpl ที่ implement MyInterface โดยมีเมธอด Method1 และ Method2
ในฟังก์ชัน main สร้างตัวแปร s ชนิด MyStruct โดยส่ง MyImpl เข้าไปเป็นค่าของ embedded type
เรียกใช้ s.Method1() และ s.Method2() ซึ่งจะเรียกใช้เมธอดของ MyImpl ที่อยู่ใน s
*/