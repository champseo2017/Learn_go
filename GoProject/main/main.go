package main

import (
	"fmt"
)

/*
แสดง interface `error` และการ implement โดย struct `errorString` ในแพ็คเกจ `errors`

// ตัวอย่าง code ของแพ็คเกจ errors

package errors

// error interface มีเพียงเมธอด Error() ที่คืนค่าเป็น string
type error interface {
    Error() string
}

// struct errorString ที่เป็นแบบ unexported
// มีฟิลด์ s เป็น string เพื่อเก็บข้อความ error
type errorString struct {
    s string
}

// errorString implement interface error โดยมีเมธอด Error() ที่คืนค่า s
func (e *errorString) Error() string {
    return e.s
}

// ฟังก์ชัน New สร้างค่า error จาก string ที่รับเข้ามา
// โดยจะสร้าง struct errorString และคืนค่าเป็น pointer ไปยัง errorString
func New(text string) error {
    return &errorString{text}
}

- มีการประกาศ interface `error` ที่มีเพียงเมธอด `Error()` ซึ่งคืนค่าเป็น string
- มีการประกาศ struct `errorString` แบบ unexported ที่มีฟิลด์ `s` เป็น string เพื่อเก็บข้อความ error
- `errorString` implement interface `error` โดยมีเมธอด `Error()` ที่คืนค่า `s`
- มีฟังก์ชัน `New` ที่รับ string `text` เพื่อสร้างค่า error โดยจะสร้าง struct `errorString` และคืนค่าเป็น pointer ไปยัง `errorString`

เนื่องจาก struct `errorString` เป็นแบบ unexported เราจึงไม่สามารถสร้างค่าของ `errorString` ได้โดยตรงจากภายนอกแพ็คเกจ `errors` ดังนั้นจึงต้องใช้ฟังก์ชัน `New()` เพื่อสร้างค่า error ผ่านการสร้าง `errorString` และคืนค่าเป็น `error` interface

เมื่อต้องการสร้างค่า error ในแพ็คเกจอื่น เราสามารถเรียกใช้ฟังก์ชัน `errors.New("error message")` เพื่อสร้างค่า error ได้
*/

type student struct {
	id   int
	name string
	age  int
}

var students = make(map[int]student)

func create(id int, name string, age int) (int, error) {
	if name == "" || age <= 0 {
		// ถ้าชื่อว่างหรืออายุน้อยกว่าหรือเท่ากับ 0 จะคืนค่า error
		return 0, fmt.Errorf("Invalid Input: name: %s, age: %d", name, age)
	}
	students[id] = student{id, name, age}
	// คืนค่า id ของนักเรียนและ nil error
	return id, nil
}

func main() {
	id, err := create(1, "", -1)
	if err != nil {
		// ถ้ามี error จะแสดงข้อความ error และจบการทำงาน
		fmt.Println(err)
		return
	}
	// ถ้าไม่มี error จะแสดงข้อมูลนักเรียน
	fmt.Println(students[id])
}

/*

 */
