package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
Slices ใน Go ซึ่งเป็นโครงสร้างข้อมูลที่ยืดหยุ่นและขยายได้ตามความต้องการ Slices มีลักษณะคล้ายกับ arrays โดยสามารถเข้าถึงข้อมูลผ่านดัชนีและมีความยาวได้ อย่างไรก็ตาม arrays ใน Go ไม่มีเมธอดในตัวสำหรับการเพิ่มขนาดแบบไดนามิกหรือการดึงส่วนย่อยของอาร์เรย์ ซึ่ง slices สามารถแก้ไขข้อจำกัดนี้ได้

บทนี้มีวัตถุประสงค์เพื่อแนะนำแนวคิดเกี่ยวกับ slices ให้กับผู้อ่าน และแสดงวิธีการประกาศ slice, วิธีการสร้าง slices, การแก้ไขและเปรียบเทียบ slices, slices หลายมิติ, การเรียงลำดับ slices และวิธีการวนลูปผ่าน slices
*/

func main() {
	// ประกาศ slice ของ strings
	// var sliceOfStrings []string

	// ประกาศ slice ของ integers 
	// var sliceOfInts []int

	// สร้าง slice จาก array
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4] // slice ประกอบด้วยเลข 2, 3, 4
	// สร้าง slice โดยใช้ make
	sliceMake := make([]int, 5) // สร้าง slice ของ integers ที่มีความยาว 5
	// แก้ไขค่าใน slice
	slice[0] = 10
	// เปรียบเทียบ slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	if reflect.DeepEqual(slice1, slice2) {
		fmt.Println("Slices are equal")
	}

	// Slice หลายมิติ
	sliceMui := [][]int{{1, 2}, {3, 4}}
	// เรียงลำดับ slice
	sort.Ints(slice)
	// วนลูปผ่าน slice
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	println(sliceMui)
	println(sliceMake)
	println(sliceMake)

}

/* 
ในตัวอย่าง code ข้างต้น แสดงวิธีการประกาศ slice, สร้าง slice จาก array และโดยใช้ make, การแก้ไขค่าใน slice, การเปรียบเทียบ slices โดยใช้ reflect.DeepEqual, การสร้าง slice หลายมิติ, การเรียงลำดับ slice ด้วยฟังก์ชัน sort.Ints และการวนลูปผ่าน slice โดยใช้ for range เพื่อเข้าถึงดัชนีและค่าของแต่ละองค์ประกอบใน slice

การสร้าง slice ด้วย `arr[1:4]` จะทำให้ได้ slice ที่ประกอบด้วยเลข 2, 3 และ 4 โดยมีการทำงานดังนี้:

1. `startIndex` คือ 1 หมายถึงเริ่มต้นที่ตำแหน่งที่ 1 ของ array (ซึ่งมีค่าเป็น 2)
2. `endIndex` คือ 4 หมายถึงสิ้นสุดที่ตำแหน่งที่ 4 ของ array (ซึ่งมีค่าเป็น 5) แต่ไม่รวมตำแหน่งที่ 4
3. ดังนั้น slice ที่ได้จะประกอบด้วยเลข 2, 3 และ 4
*/