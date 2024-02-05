package main

import (
	"fmt"
)

/*
Sscanln สแกนอ่านข้อมูลจากสตริงที่กำหนด

Sscanln() แตก string ออกเป็นส่วนๆ และ นำไปใช้งาน Sscanln() จะหยุด scan เมื่อพบ \n

*/

func main() {

	message := "I love golang \n Golang programming"

	var s1, s2, s3, s4 string
	var ss1, ss2, ss3, ss4 string

	fmt.Sscan(message, &s1, &s2, &s3, &s4)
	fmt.Printf("%s %s %s %s", s1, s2, s3, s4)
	fmt.Println()
	fmt.Sscanln(message, &ss1, &ss2, &ss3, &ss4)
	fmt.Printf("%s %s %s %s", ss1, ss2, ss3, ss4)

}
