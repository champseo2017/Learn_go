package main

import (
	"fmt"
	"strconv"
)

/*
 ParseBool() แปลงค่าสตริงไปเป็นบูลีน

*/

func main() {

	str1 := "1"
	str2 := "0"
	str3 := "true"
	str4 := "false"

	fmt.Println(strconv.ParseBool(str1))
	fmt.Println(strconv.ParseBool(str2))
	fmt.Println(strconv.ParseBool(str3))
	fmt.Println(strconv.ParseBool(str4))

}
