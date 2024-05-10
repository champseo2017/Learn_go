package main

import (
	"fmt"
)

/*
Switch Case เป็นโครงสร้างควบคุมที่ใช้สำหรับการตรวจสอบเงื่อนไขเมื่อมีเงื่อนไขจำนวนมาก และช่วยให้โค้ดมีความชัดเจนและเหมาะสมกว่าการใช้ if-else หลายๆ ครั้ง
*/
func Switch_Case_Example(x string) {
	switch x {
	case "Rahul": // ถ้า x เป็น "Rahul"
		{
			fmt.Println("This is Rahul") // แสดงข้อความ "This is Rahul"
		}
	case "Vikas": // ถ้า x เป็น "Vikas"
		{
			fmt.Println("This is Vikas") // แสดงข้อความ "This is Vikas"
		}
	case "Arjun": // ถ้า x เป็น "Arjun"
		{
			fmt.Println("This is Arjun") // แสดงข้อความ "This is Arjun"
		}
	default: // ถ้า x ไม่ตรงกับเคสใดๆ
		{
			fmt.Println("He has no name") // แสดงข้อความ "He has no name"
		}
	}
}

func main() {
	Switch_Case_Example("Arjun")
}
