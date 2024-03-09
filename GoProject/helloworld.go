package main

import (
	"fmt"
	"strings"
)

/*

Derived types
ประเภทข้อมูลที่ได้มาจากการแปลงในภาษา Go คือการขยายความสามารถของประเภทข้อมูลพื้นฐานเพื่อสร้างโครงสร้างข้อมูลใหม่ๆ ที่มีความซับซ้อนและเฉพาะเจาะจงมากขึ้น ได้แก่:

โครงสร้าง (Structures): รวมข้อมูลหลายประเภทเข้าด้วยกัน
ฟังก์ชัน (Functions): กำหนดการทำงานหรือการดำเนินการ
สไลซ์ (Slices): ชุดข้อมูลที่ขนาดสามารถเปลี่ยนแปลงได้
อินเทอร์เฟซ (Interfaces): กำหนดชุดเมธอดที่ประเภทข้อมูลต้องการใช้งาน
แมพ (Maps): จัดเก็บข้อมูลในรูปแบบคีย์และค่า
ชาแนล (Channels): สำหรับการสื่อสารระหว่างกอรูทีน

*/

// ประเภทโครงสร้าง
type Person struct {
	Name string
	Age int
}

// ฟังก์ชัน (Functions)
/* 
// วิธีการทำงาน
// 1. ฟังก์ชัน greet รับพารามิเตอร์ name ที่เป็น string
// 2. ฟังก์ชันจะส่งค่ากลับเป็น string ที่เป็นการรวมคำว่า "Hello, " กับค่าของ name
// 3. เมื่อเรียกใช้ฟังก์ชัน greet และส่งค่า name เข้าไป เช่น greet("Alice") จะได้ผลลัพธ์คือ "Hello, Alice"
*/
func greet(name string) string {
	return "Hello, " + name
}

// สไลซ์ (Slices)
// ทำไมต้องกำหนด range for _, number := range numbers
// 1. range ใช้เพื่อวนซ้ำ (iterate) ผ่านทุกสมาชิกใน slice, array, map หรือ string
// 2. ในการวนซ้ำผ่าน slice หรือ array, range จะส่งคืนค่าสองตัวในแต่ละรอบของ loop: index และค่าของสมาชิก
func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// อินเทอร์เฟซ (Interfaces)
type Animal interface {
	Speak() string
}

// กำหนดโครงสร้างของ Dog ที่มีชื่อเป็นสตริง
type Dog struct {
    Name string
}

// กำหนดเมธอด Speak สำหรับ struct Dog
func (d Dog) Speak(times int) string {
    // ใช้ฟิลด์ Name จาก struct Dog ในการสร้างข้อความ
    // สร้างข้อความที่ประกอบด้วยชื่อของ Dog ตามด้วย " says Woof! "
    // และใช้ strings.Repeat เพื่อซ้ำข้อความนี้ตามจำนวนครั้งที่ระบุในพารามิเตอร์ times
    woof := strings.Repeat(d.Name + " says Woof! ", times)
    // ส่งคืนข้อความที่สร้างขึ้น
    return woof
}

func main() {
   myDog := Dog{Name: "Buddy"}
   // เรียกใช้เมธอด Speak โดยส่งจำนวนครั้งที่ต้องการให้ Dog "Woof!"
   sound := myDog.Speak(3) // ต้องการให้ Dog "Woof!" 3 ครั้ง

   // แสดงผลลัพธ์
   fmt.Println(sound) // จะแสดง "Buddy says Woof! Buddy says Woof! Buddy says Woof! "
}