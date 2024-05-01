package main

import "fmt"

/*
Q7: การดึงค่า value จาก key ใน map โดยการเขียนโปรแกรม Go
*/

func main() {
	// สร้าง map
	capitals := map[string]string {
		"Thailand": "Bangkok",
        "Japan":    "Tokyo",
        "USA":      "Washington D.C.",
	}
	// ดึงค่า value จาก key
	capital := capitals["Thailand"]
	fmt.Printf("Capital of Thailand: %s\n", capital)

	// ดึงค่า value จากkey ที่ไม่มีอยู่ใน map
	capital, ok := capitals["Germany"]
	if ok {
		fmt.Printf("Capital of Germany: %s\n", capital)
	} else {
		fmt.Println("Key 'Germany' does not exist")
	}
}

/* 
การดึงค่า value จาก key ทำได้โดยการเขียนในรูปแบบ value := mapname[key] เช่น capital := capitals["Thailand"] จะได้ value "Bangkok"
กรณีที่เราดึงค่า value จาก key ที่ไม่มีอยู่ใน map เช่น key "Germany" เราจะได้ค่า zero value ของ value type นั้นกลับมา (string ก็จะได้ "" กลับมา)
วิธีตรวจสอบว่า key มีอยู่ในmap หรือไม่ทำได้โดยการใช้คำสั่ง value, ok := mapname[key] ซึ่งจะได้ value กลับมาพร้อมกับ boolean ค่า ok เป็น true ถ้ามีkey นั้นอยู่
*/