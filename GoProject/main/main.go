package main

import (
	"fmt"
)

/*
โปรแกรมแรกจะเกิด compilation error เพราะเราไม่สามารถกำหนดค่าที่อยู่ของตัวแปรชนิด float ให้กับตัวแปรพอยน์เตอร์ชนิด int ได้

โปรแกรมที่สองจะเกิด runtime error เนื่องจากมีการdereference ตัวแปรพอยน์เตอร์ a ซึ่งมีค่าเป็น nil
*/

func main() {
	// var a *int      // ประกาศตัวแปรพอยน์เตอร์ a ชนิด int
	// b := 10.3       // ประกาศและกำหนดค่าให้ตัวแปร b เป็น 10.3 (float64)
	// a = &b          // กำหนดที่อยู่ของ b ให้กับ a (compilation error เพราะชนิดข้อมูลไม่ตรงกัน)
	// fmt.Println(*a) // พยายาม dereference a ซึ่งจะทำให้เกิด error เพราะ a ไม่ได้ชี้ไปที่ valid memory address

	var a *int      // ประกาศตัวแปรพอยน์เตอร์ a ชนิด int (ค่าเริ่มต้นเป็น nil)
	fmt.Println(*a) // พยายาม dereference a ซึ่งมีค่าเป็น nil ทำให้เกิด runtime error (nil pointer dereference)
}

/*

 */
