package main

import "fmt"

/*
Go สามารถสร้าง slice ได้โดยใช้ slice literal ซึ่งเป็นวิธีการสร้าง slice ด้วยการระบุค่าของ slice ไว้ในเครื่องหมาย [] ตามด้วยประเภทข้อมูลของ slice เช่น []int{2, 4, 6, 8, 10} เมื่อสร้าง slice ด้วยวิธีนี้ ภาษา Go จะสร้าง array ขึ้นมาก่อน แล้วจึงคืนค่า reference ของ slice ที่อ้างอิงไปยัง array นั้น นอกจากนี้ยังสามารถสร้าง slice ได้ด้วยการใช้ shorthand declaration เช่น s := []int{1, 3, 5, 7, 9}
*/

func main() {
	// Create a slice using a slice literal
	var s = []int{2, 4, 6, 8, 10}
	fmt.Println("slice = ", s)
}

/* 
ในโปรแกรมนี้ เราสร้าง slice ชื่อ s ด้วยการใช้ slice literal []int{2, 4, 6, 8, 10} ซึ่งจะสร้าง slice ที่มีค่าเป็น [2, 4, 6, 8, 10] จากนั้นเราใช้ fmt.Println() เพื่อแสดงค่าของ slice s ออกทางหน้าจอ ผลลัพธ์ที่ได้คือ "slice = [2 4 6 8 10]"
*/