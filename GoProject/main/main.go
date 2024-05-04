package main

import "fmt"

/*
ตัวอย่างการกำหนด Method ให้กับ Type ต่างๆ
*/

// Method ของ Basic Type (string):
type MyString string

func (s MyString) Reverse() string {
    runes := []rune(s)
	/* 
	runes := []rune(s) เป็นการแปลง MyString เป็น Slice ของ rune ([]rune) เพื่อให้สามารถทำงานกับตัวอักษรแต่ละตัวในสตริงได้อย่างถูกต้อง โดยเฉพาะอย่างยิ่งเมื่อสตริงมีตัวอักษรที่ไม่ใช่ ASCII เช่น ตัวอักษรภาษาไทย
	*/
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
	str := MyString("สวัสดี")
	reversed := str.Reverse()
	fmt.Println(reversed) //  ีดสัวส
}
/* 

*/