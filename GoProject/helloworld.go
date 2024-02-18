package main

import (
	"fmt"
)

/*

 การใช้ Sprintf() แปลงค่าเลขจำนวนเต็มไปเป็นสตริง

 - `fmt.Scanf` เป็นฟังก์ชันในภาษา Go ที่ใช้สำหรับอ่านข้อมูลที่ป้อนเข้ามาจาก standard input (stdin) ตามรูปแบบที่กำหนด ในกรณีนี้ "%d %d %d" หมายถึง รูปแบบการรับข้อมูลเป็นจำนวนเต็มทั้งสามตัว และจะเก็บค่าเหล่านั้นไว้ในตัวแปร num1, num2, และ num3 ตามลำดับ

- `fmt.Sprintf` เป็นฟังก์ชันที่ใช้สร้าง string จากข้อมูลที่กำหนดตามรูปแบบที่ระบุ เช่นในกรณีนี้ "%d, %d and %d" ใช้สำหรับสร้าง string จากตัวเลขทั้งสามที่รับมา โดยจะเปลี่ยนค่าตัวเลขให้เป็นข้อความตามรูปแบบที่กำหนดและเก็บไว้ในตัวแปร str

- `%T` เป็น verb ที่ใช้ในฟังก์ชันที่เกี่ยวข้องกับการพิมพ์ข้อมูลในภาษา Go (เช่น `fmt.Printf`) สำหรับแสดงประเภทของตัวแปร ในตัวอย่างนี้ ใช้ `%T` เพื่อแสดงประเภทของตัวแปร `str` ซึ่งจะแสดงผลว่าเป็น `string`

*/

func main() {
   
	var num1, num2, num3 int

	fmt.Print("Please enter 3 numbers : ")
	_, err := fmt.Scanf("%d %d %d", &num1, &num2, &num3)

	if err == nil {
		str := fmt.Sprintf("%d, %d and %d", num1, num2, num3)
		fmt.Printf("Type of str is %T, value is %s\n", str, str)
	} else {
		fmt.Printf("Error is \"%v\"\n", err)
	}
	

}
