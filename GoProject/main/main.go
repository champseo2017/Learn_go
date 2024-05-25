package main

import (
	"encoding/json"
	"fmt"
)

/*
แสดงการใช้งานฟังก์ชัน json.Marshal และ json.Unmarshal ในการแปลงข้อมูลจากและไปยังรูปแบบ JSON ตามลำดับ โดยมีการสร้างข้อมูลประเภท Student ขึ้นมาก่อน แล้วจึงแปลงข้อมูลนั้นให้อยู่ในรูปแบบ JSON ด้วย json.Marshal และแสดงผลลัพธ์ออกมา จากนั้นจะนำผลลัพธ์ที่ได้ไปแปลงกลับเป็นข้อมูลประเภท Student ด้วย json.Unmarshal และแสดงผลลัพธ์สุดท้ายออกมา
*/

// ประกาศ struct ชื่อ Student ที่มีฟิลด์ Name และ Subject
type Student struct {
	Name    string
	Subject string
}

// ฟังก์ชัน JSON_Data_Marshalling() ทำหน้าที่แปลงข้อมูล Student ให้อยู่ในรูปแบบ JSON
func JSON_Data_Marshalling() []byte {
	// สร้างข้อมูล Student
	s := Student{"Udit", "Physics"}

	// แปลงข้อมูล Student ให้อยู่ในรูปแบบ JSON ด้วย json.Marshal
	d, err := json.Marshal(s)
	if err != nil {
		fmt.Println("the error is as follows")
		return nil
	} else {
		fmt.Println(d)
		return d
	}
}

// ฟังก์ชัน JSON_Data_UnMarshalling() ทำหน้าที่แปลงข้อมูลที่อยู่ในรูปแบบ JSON กลับมาเป็นข้อมูลประเภท Student
func JSON_Data_UnMarshalling() Student {
	// เรียกใช้ฟังก์ชัน JSON_Data_Marshalling() เพื่อได้ข้อมูลในรูปแบบ JSON
	d := JSON_Data_Marshalling()

	// สร้างตัวแปรเก็บข้อมูลประเภท Student
	var s Student

	// แปลงข้อมูลที่อยู่ในรูปแบบ JSON กลับมาเป็นข้อมูลประเภท Student ด้วย json.Unmarshal
	json.Unmarshal(d, &s)

	// ส่งคืนข้อมูลประเภท Student
	return s
}

func main() {
	// เรียกใช้ฟังก์ชัน JSON_Data_UnMarshalling() และแสดงผลลัพธ์
	s := JSON_Data_UnMarshalling()
	fmt.Println("This is student", s)
}

/*

 */
