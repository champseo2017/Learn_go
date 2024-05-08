package main

import "fmt"

/*

ตัวแปร interface สามารถใช้เป็น field ของ struct ได้ เมื่อสร้าง object ของ struct นั้น เราต้องส่งค่า concrete type ที่ implement interface ให้กับ field ของ struct
ในตัวอย่างโปรแกรม 9.13 มีการสร้าง struct ชื่อ Process ที่มี field ชื่อ exe เป็นตัวแปรของ interface ชื่อ Executor โดย struct ชื่อ Thread ทำการ implement Executor
ในฟังก์ชัน main มีการสร้างตัวแปรของ Process และส่งค่า Thread ให้กับ field exe ที่เป็น Executor จากนั้นเรียกใช้เมธอด Execute ผ่าน field exe ของตัวแปร p ซึ่งจะเรียกใช้เมธอด Execute ของ Thread ที่อยู่ใน p

*/

type Executor interface {
    Execute()
}

type Thread struct{}

func (t Thread) Execute() {
    fmt.Println("Executing thread")
}

type Process struct {
    exe Executor
}

func main() {
    p := Process{Thread{}}
    p.exe.Execute()
}

/* 
สร้าง interface ชื่อ Executor ที่มีเมธอด Execute
สร้าง struct ชื่อ Thread
กำหนดให้ Thread implement interface Executor โดยมีเมธอด Execute ที่แสดงข้อความ "Executing thread"
สร้าง struct ชื่อ Process ที่มี field ชื่อ exe เป็นตัวแปรของ interface Executor
ในฟังก์ชัน main สร้างตัวแปร p เป็น object ของ Process และส่งค่า Thread{} ให้กับ field exe
เรียกใช้เมธอด Execute ผ่าน field exe ของตัวแปร p ซึ่งจะเรียกใช้เมธอด Execute ของ Thread ที่อยู่ใน p

ค่า concrete type ที่ implement interface Executor และถูกส่งให้กับ field exe ของ struct Process คือ Thread{}

ในบรรทัดนี้ p := Process{Thread{}} เราสร้าง object ของ struct Process และกำหนดค่าให้กับ field exe ด้วย Thread{} ซึ่ง Thread เป็น concrete type ที่ implement interface Executor
เมื่อเรียกใช้ p.exe.Execute() การเรียกใช้เมธอด Execute จะถูกส่งต่อไปยังเมธอด Execute ของ Thread ที่อยู่ใน field exe ของ p นั่นเอง
ดังนั้น ค่า concrete type ที่ถูกส่งให้กับ field ของ struct Process ในตัวอย่างนี้คือ Thread{}
*/