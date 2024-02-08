package main

import (
	"fmt"
	"strconv"
)

/*
 ฟังค์ชั่น Atoi() แปลงค่าสตริงไปเป็นเลขจำนวนเต็ม int ฐาน 10

*/

func main() {

	var str = "9876"
	number, err := strconv.Atoi(str)
	number2, err := strconv.ParseInt(str, 10, 0)

	if err == nil {
		fmt.Printf("string \"%s\" with Atoi() to int : %d\n", str, number)
		fmt.Printf("string \"%s\" with ParseInt() to int : %d\n", str, number2)
	} else {
		fmt.Printf("Error is %v\n", err)
	}

}
