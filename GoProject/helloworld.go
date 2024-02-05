package main

import (
	"fmt"
)

/*
Sscan สแกนอ่านข้อมูลจากสตริงที่กำหนด

Sscan() แตก string ออกเป็นส่วนๆ และ นำไปใช้งาน

*/

func main() {

	myStr := "There are 100"
	var s1, s2 string
	var i int

	fmt.Sscan(myStr, &s1, &s2, &i)
	fmt.Printf("Form string \"%s\"\n\n", myStr)
	fmt.Println("We can scan string with sscan()")
	fmt.Printf("String 1 : %s\n", s1)
	fmt.Printf("String 2 : %s\n", s2)
	fmt.Printf("%s birds at the zoo %d \n", s2, i)

}
