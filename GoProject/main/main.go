package main

import (
	"fmt"
)

/*
โปรแกรม 12.1 เป็นตัวอย่างการจัดการข้อผิดพลาดในภาษา Go โดยมีฟังก์ชัน create() ที่ใช้สำหรับสร้างข้อมูลนักเรียนและเพิ่มลงใน (map)
*/

type student struct {
	id   int
	name string
	age  int
}

var stuMap = make(map[int]student)

func create(name string, age int) (int, error) {
	if name == "" {
		// ถ้าชื่อว่างเปล่า ให้คืนค่า error
		return 0, fmt.Errorf("Invalid Name, Name can't be Empty")
	} else if age <= 0 {
		// ถ้าอายุน้อยกว่าหรือเท่ากับ 0 ให้คืนค่า error
		return 0, fmt.Errorf("Invalid Age %v, Age can't be <= 0", age)
	}
	// สร้าง id ใหม่จากความยาวของ stuMap บวก 1
	id := len(stuMap) + 1
	// สร้างข้อมูลนักเรียนใหม่
	stu := student{id, name, age}
	// เพิ่มข้อมูลนักเรียนลงใน stuMap โดยใช้ id เป็น key
	stuMap[id] = stu
	// คืนค่า id และ nil error
	return id, nil
}

func main() {
	// เรียกใช้ฟังก์ชัน create() และรับค่าที่คืนมาเป็น id และ error
	id, err := create("John", 10)
	if err != nil {
		// ถ้ามี error ให้แสดงข้อความแจ้งเตือนและ error message
		fmt.Println("Error while creating student")
		fmt.Println(err)
		return
	}
	// ถ้าไม่มี error ให้แสดงข้อความยืนยันการสร้างนักเรียนพร้อม id
	fmt.Println("Student created with Id=", id)
}

/*

 */
