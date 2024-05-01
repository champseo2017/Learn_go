package main

import "fmt"

/*
Q8: เขียนโปรแกรม Go เพื่ออธิบายการตรวจสอบว่ามี key อยู่ใน map หรือไม่
*/

func main() {
	// สร้าง map
	employees := map[string]int {
		"Alice": 1001,
        "Bob":   1002,
        "Charlie": 1003,
	}

	// ตรวจสอบว่ามี key หรือไม่
	id, exists := employees["Alice"]
	if exists {
		fmt.Printf("Alice's ID is %d\n", id)
	} else {
		fmt.Println("Alice does not exist")
	}

	_, exists2 := employees["Eve"]
	if !exists2 {
		fmt.Println("Eve does not exist")
	}
}

/* 
การตรวจสอบว่ามี key อยู่ในmap หรือไม่ ทำได้โดยการเขียนในรูปแบบ value, ok := mapname[key]
ถ้า ok เป็น true แสดงว่ามี key นั้นอยู่ใน map และ value จะได้รับค่า value ที่เชื่อมโยงกับ key นั้น
ถ้า ok เป็น false แสดงว่าไม่มี key นั้นอยู่ใน map และ value จะได้รับ zero value ของ value type นั้น
ในตัวอย่างนี้ เรามีการตรวจสอบ key "Alice" และ "Eve" ว่ามีอยู่ใน map หรือไม่ พร้อมทั้งแสดงผลลัพธ์ด้วย
*/