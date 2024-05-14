package main

import (
	"fmt"
)

/*
Map ใน Go เป็นแบบ unordered collection หมายความว่าลำดับของ key-value ใน map ไม่ได้ถูกกำหนดแน่นอนตายตัว และอาจมีการจัดเรียงลำดับใหม่ในแต่ละครั้งที่เรา loop map
*/

func Loop_Map() {
	// สร้าง map ที่มี key เป็น string และ value เป็น int
	scores := map[string]int{
		"Alice":   90,
		"Bob":     80,
		"Charlie": 95,
	}

	// loop map ด้วย for และ range
	for name, score := range scores {
		fmt.Printf("Name: %s, Score: %d\n", name, score)
	}
}

func main() {
	Loop_Map()
}

/*
 เมื่อเรา loop map scores ด้วย for และ range ลำดับของ key-value ที่ถูกแสดงผลอาจแตกต่างกันในแต่ละครั้งที่ run เนื่องจาก map ใน Go เป็นแบบ unordered

 การที่ map เป็น unordered ช่วยให้การเข้าถึง value ด้วย key เป็นไปอย่างรวดเร็ว เพราะไม่จำเป็นต้องเก็บลำดับของ key-value เอาไว้ แต่ผลข้างเคียงคือเมื่อเรา loop map ลำดับที่ได้อาจไม่เหมือนเดิมในแต่ละครั้ง
	ดังนั้นหากต้องการ loop map และต้องการลำดับที่แน่นอน เราอาจต้องเรียง key ของ map ก่อนแล้วค่อย loop ตาม key ที่เรียงลำดับแล้ว จึงจะได้ลำดับที่แน่นอนเหมือนกันในทุกครั้งที่ run
*/
