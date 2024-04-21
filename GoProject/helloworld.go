package main

import "fmt"

/*
Go เราสามารถสร้าง slice จาก array ได้โดยระบุ index ต่ำสุดและสูงสุดที่คั่นด้วยเครื่องหมายโคลอน (:) ในรูปแบบ array_name[low:high] โดย slice ที่ได้จะประกอบด้วยองค์ประกอบตั้งแต่ index ต่ำสุดจนถึงก่อนหน้า index สูงสุด นอกจากนี้ index ต่ำสุดและสูงสุดยังสามารถละไว้ได้ โดยค่าเริ่มต้นของ index ต่ำสุดคือ 0 และค่าเริ่มต้นของ index สูงสุดคือความยาวของ slice
*/

func main() {
	var a = [6]string{"Jaipur", "Gorakhpur", "Pune", "Bengaluru", "Mumbai", "Delhi"}
	var s []string = a[1:5]
	fmt.Println("Array a =", a)
	fmt.Println("Slice s =", s)
}

/* 
ในโปรแกรมนี้ เราสร้าง array a ที่มี 6 องค์ประกอบ จากนั้นเราสร้าง slice s จาก array a โดยระบุ index ต่ำสุดเป็น 1 และ index สูงสุดเป็น 5 ดังนั้น slice s จะประกอบด้วยองค์ประกอบตั้งแต่ index 1 ถึง 4 ของ array a เมื่อเราแสดงค่าของ array a และ slice s ผลลัพธ์ที่ได้คือ "Array a = [Jaipur Gorakhpur Pune Bengaluru Mumbai Delhi]" และ "Slice s = [Gorakhpur Pune Bengaluru Mumbai]"
*/