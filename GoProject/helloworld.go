package main

import "fmt"

/*
ในการเข้าถึงฟิลด์แต่ละตัวของ struct เราต้องใช้เครื่องหมายจุด (.) ตามด้วยชื่อของฟิลด์ที่ต้องการเข้าถึง เราสามารถใช้เครื่องหมายจุดเพื่ออ่านหรือกำหนดค่าให้กับฟิลด์ของ struct ได้
*/

type bike struct {
	name, model, color string
	weight_in_kg float64
}

func main() {
	b := bike{
		name:"Bajaj",
        model:"Pulsar_150cc",
        color:"Black Grey",
        weight_in_kg: 148,
	}
	// เข้าถึงฟิลด์ของ struct โดยใช้เครื่องหมายจุด
	fmt.Println("bike name:", b.name)
	fmt.Println("bike color:", b.color)
	// กำหนดค่าใหม่ให้กับฟิลด์ของ struct
	b.color = "Neon Red"
	// แสดงผลลัพธ์
	fmt.Println("bike:", b)
}

/* 
1. เรากำหนด struct ชื่อ `bike` ที่มีฟิลด์ดังนี้:
   - `name` เป็น string
   - `model` เป็น string
   - `color` เป็น string
   - `weight_in_kg` เป็น float64

2. ในฟังก์ชัน `main` เราสร้างตัวแปร `b` ของประเภท `bike` และกำหนดค่าเริ่มต้นให้กับฟิลด์ต่างๆ โดยใช้ struct literal

3. เราเข้าถึงฟิลด์ของ struct `b` โดยใช้เครื่องหมายจุด:
   - `b.name` เพื่อเข้าถึงฟิลด์ `name`
   - `b.color` เพื่อเข้าถึงฟิลด์ `color`

4. เราแสดงค่าของฟิลด์ `name` และ `color` โดยใช้ `fmt.Println`

5. เรากำหนดค่าใหม่ให้กับฟิลด์ `color` ของ struct `b` โดยใช้เครื่องหมายจุด: `b.color = "Neon Red"`

6. สุดท้าย เราแสดงค่าทั้งหมดของ struct `b` โดยใช้ `fmt.Println`

จะเห็นว่าเราสามารถเข้าถึงและแสดงค่าของฟิลด์ `name` และ `color` ของ struct `b` ได้อย่างถูกต้อง และเมื่อเรากำหนดค่าใหม่ให้กับฟิลด์ `color` ค่าของฟิลด์นั้นก็ถูกอัปเดตอย่างถูกต้องเช่นกัน
*/