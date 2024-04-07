package main

import (
	"fmt"
	"strings"
)

/*
Variadic function คือฟังก์ชันที่สามารถรับอาร์กิวเมนต์ได้หลายค่า จำนวนไม่คงที่ ต่างจากฟังก์ชันทั่วไปที่รับได้จำนวนคงที่ตามที่ประกาศไว้ โดยจะใช้ ... นำหน้าพารามิเตอร์ตัวสุดท้ายในการประกาศฟังก์ชัน และพารามิเตอร์ที่เป็น variadic จะต้องเป็นตัวสุดท้ายเท่านั้น

วิธีการทำงานของ variadic function คือ แปลงอาร์กิวเมนต์ทั้งหมดให้เป็น slice ของประเภทเดียวกับพารามิเตอร์ที่เป็น variadic ทำให้ง่ายต่อการจัดการและเพิ่มความยืดหยุ่นให้กับฟังก์ชัน เหมาะกับกรณีที่ไม่รู้จำนวนอาร์กิวเมนต์ที่แน่นอน

นอกจากส่งอาร์กิวเมนต์ทีละค่าแล้ว เรายังสามารถส่ง slice เข้าไปใน variadic function ได้ด้วย และสามารถเข้าถึงอาร์กิวเมนต์แต่ละตัวผ่าน index ของ slice ได้

ตัวอย่างการเข้าถึงอาร์กิวเมนต์แต่ละตัวใน variadic function

func variadicExample(depts ...string) {
    fmt.Println(depts[1])
    fmt.Println(depts[4])
}

func main() {
    variadicExample("IT", "Finance", "HR", "Recruitment", "Payroll")
}

*/

// ฟังก์ชัน join เป็น variadic function ที่รับพารามิเตอร์ elements เป็น string ได้หลายค่า
func join(elements ...string) string {
	// ภายในฟังก์ชันใช้ strings.Join เพื่อรวม elements เข้าด้วยกัน โดยใช้ _ เป็นตัวคั่น
	return strings.Join(elements, "_")
}

func main() {
	fmt.Println(join("GO", "language", "book"))

	elements := []string{"GO", "language", "book"}
	fmt.Println(join(elements...))
	// ... ที่เขียนต่อท้าย elements เรียกว่า "variadic argument" หรือ "argument unpacking" มันเป็นวิธีการส่ง slice เข้าไปใน variadic function โดยไม่ต้องส่งทีละค่า

	// join([]string{"GO", "language", "book"}...)
	// แทนที่จะต้องเขียนแบบนี้ join("GO", "language", "book")
}