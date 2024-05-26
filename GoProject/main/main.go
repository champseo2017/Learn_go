package main

import (
	"fmt"
)

/*
สามารถสร้าง custom error ของตัวเองได้โดยการ implement interface error ซึ่งมีเพียงหนึ่งเมธอดคือ Error() ที่คืนค่าเป็น string ตัวอย่างเช่นการสร้าง InvalidAgeError
*/

type student struct {
	id   int
	name string
	age  int
}

// InvalidAgeError struct ที่ implement interface error
type InvalidAgeError struct {
	age int
}

type InvalidNameError struct {
	name string
}

var stuMap = make(map[int]student)

// implement เมธอด Error() เพื่อแสดงข้อความ error สำหรับ InvalidAgeError
func (i InvalidAgeError) Error() string {
	return fmt.Sprintf("Invalid Age %v, Age can't be <= 0", i.age)
}

func (i InvalidNameError) Error() string {
	return fmt.Sprintf("Invalid Name %s", i.name)
}

// ตัวอย่างการคืนค่า InvalidAgeError ในฟังก์ชัน create()
func create(name string, age int) (int, error) {
	if name == "" {
		return 0, InvalidNameError{name}

	} else if age <= 0 {
		// คืนค่า InvalidAgeError พร้อมกับ age ที่ไม่ถูกต้อง
		return 0, InvalidAgeError{age}
	}
	// สร้างและคืนค่า id และ nil error เหมือนเดิม
	id := len(stuMap) + 1
	stu := student{id, name, age}
	stuMap[id] = stu
	return id, nil
}

func main() {
	// เรียกใช้ฟังก์ชัน create() โดยส่งค่า name เป็นค่าว่าง และ age เป็น 10
	_, err := create("", 10)
	if ageError, ok := err.(InvalidAgeError); ok {
		// ถ้า error เป็น InvalidAgeError ให้แสดงข้อความพร้อมกับค่า age ที่ไม่ถูกต้อง
		fmt.Println("This is InvalidAgeError. Age:", ageError.age)
		// ดำเนินการอื่น ๆ สำหรับ InvalidAgeError
	} else if nameError, ok := err.(InvalidNameError); ok {
		// ถ้า error เป็น InvalidNameError ให้แสดงข้อความพร้อมกับค่า name ที่ไม่ถูกต้อง
		fmt.Println("This is InvalidNameError. Name:", nameError.name)
		// ดำเนินการอื่น ๆ สำหรับ InvalidNameError
	}
}

/*
1. `ok` ตัวแรก (`ok := err.(InvalidAgeError)`) เป็นการประกาศตัวแปร `ok` เพื่อรับค่าผลลัพธ์จากการแปลงประเภท (type assertion) ของ `err` เป็น `InvalidAgeError`
   - ถ้าการแปลงประเภทสำเร็จ `ageError` จะมีค่าเป็น error ที่แปลงเป็นประเภท `InvalidAgeError` และ `ok` จะมีค่าเป็น `true`
   - ถ้าการแปลงประเภทไม่สำเร็จ `ageError` จะมีค่าเป็น zero value ของประเภท `InvalidAgeError` และ `ok` จะมีค่าเป็น `false`

2. `ok` ตัวที่สอง (`; ok {`) เป็นการใช้ค่าของตัวแปร `ok` ที่ได้จากการแปลงประเภทในข้อ 1 เพื่อตรวจสอบว่าการแปลงประเภทสำเร็จหรือไม่ ถ้า `ok` เป็น `true` (แปลงสำเร็จ) โค้ดภายในบล็อก `if` จะถูกเรียกทำงาน

ดังนั้นบรรทัดนี้จึงเป็นการตรวจสอบว่า error ที่ได้รับ (`err`) สามารถแปลงเป็นประเภท `InvalidAgeError` ได้หรือไม่ โดยใช้ type assertion ถ้าแปลงได้ (แปลงสำเร็จ) ค่า `ok` จะเป็น `true` และโค้ดภายในบล็อก `if` จะถูกเรียกทำงาน ซึ่งเราสามารถเข้าถึงค่า `age` จาก `ageError` ได้ภายในบล็อกนั้น
*/
