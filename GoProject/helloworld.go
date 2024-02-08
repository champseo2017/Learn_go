package main

import (
	"fmt"
	"strconv"
)

/*
 ฟังค์ชั่น ParseInt() สตริง ไปเป็นเลขจำนวนเต็ม

*/

func main() {

	var str = "1110"
	number, err := strconv.ParseInt(str, 10, 64)
	number2, err := strconv.ParseInt(str, 2, 64)

	if err == nil {
		fmt.Printf("string \"%s\" with int64 base 10 : %d\n", str, number)
		fmt.Printf("string \"%s\" with int64 base 2 : %d\n", str, number2)
	} else {
		fmt.Printf("Error is %v\n", err)
	}

}
