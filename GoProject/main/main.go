package main

import (
	"fmt"
)

/*

แสดงให้เห็นถึงการฝัง (embed) interface เข้าไปใน struct โดยตรง ซึ่งแตกต่างจากโปรแกรม 9.13 ที่ใช้ interface เป็น field ของ struct

การฝัง interface เข้าไปใน struct ทำให้เราสามารถเรียกใช้ method ของ interface ผ่าน struct ได้โดยตรง โดยไม่ต้องผ่าน field ของ struct ที่เป็น interface เหมือนในโปรแกรม 9.13

ในโปรแกรมนี้ เรามี interface `Executor` ที่มี method `Execute()` และมี struct `Thread` ที่ implement method นี้ เราสร้าง struct `Process` ที่ฝัง interface `Executor` เข้าไปโดยตรง และสร้าง instance ของ `Process` ด้วย `Thread` เพื่อเรียกใช้ method `Execute()` ผ่าน `Process` ได้เลย

*/

type Executor interface {
    Execute()
}

type Thread struct {
}

func (t Thread) Execute() {
    fmt.Println("Executing thread")
}

type Process struct {
    Executor
}

func main() {
    p := Process{Thread{}}
    p.Execute()
}

/* 
จากโปรแกรมนี้ เราจะเห็นว่าการฝัง interface เข้าไปใน struct โดยตรง ทำให้เราสามารถเรียกใช้ method ของ interface ผ่าน struct ได้เลย โดยไม่ต้องผ่าน field ของ struct ที่เป็น interface เหมือนในโปรแกรม 9.13
*/