package main

import (
	"fmt"
	"sort"
)

/*
Interface ใน package sort ของ Go ใช้สำหรับเรียงลำดับ slice ของ user-defined types โดยมีเมธอด 3 ตัวที่ต้อง implement คือ Len, Swap และ Less

หากต้องการใช้ sort.Sort เพื่อเรียงลำดับ slice เราต้องสร้าง type ใหม่ที่เป็น slice ของ type ที่ต้องการเรียงลำดับ และ implement Interface ให้กับ type ใหม่นั้น โดย:
- Len คืนความยาวของ slice
- Swap สลับตำแหน่งของ element ใน slice
- Less เปรียบเทียบว่า element ตัวใดน้อยกว่า

จากนั้นเรียกใช้ sort.Sort โดยส่ง slice ที่ implement Interface เข้าไป ซึ่งจะเรียงลำดับ slice ตามเงื่อนไขที่กำหนดใน Less

ตัวอย่างโค้ดแสดงการสร้าง struct Student และ type ByMarks ที่เป็น slice ของ Student และ implement Interface เพื่อเรียงลำดับ slice ของ Student ตามคะแนน (Marks) จากน้อยไปมาก
*/

type Student struct {
    Name  string
    Marks int
}

func (s Student) String() string {
    return fmt.Sprintf("%s: %d", s.Name, s.Marks)
}

type ByMarks []Student

func (b ByMarks) Len() int { return len(b) }
func (b ByMarks) Swap(i, j int)  { b[i], b[j] = b[j], b[i] }
func (b ByMarks) Less(i, j int) bool { return b[i].Marks < b[j].Marks }

func main() {
    students := []Student{
        {"Bob", 31},
        {"John", 42},
        {"Michael", 17},
        {"Jenny", 26},
    }
    fmt.Println("Before sorting =", students)
    sort.Sort(ByMarks(students))
    /* 
    เรียกใช้ sort.Sort โดยแปลง students เป็น ByMarks เพื่อให้ Sort ใช้ Interface ที่เรา implement ไว้
    */
    fmt.Println("After sorting =", students)
}
