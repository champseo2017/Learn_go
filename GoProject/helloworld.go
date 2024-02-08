package main

import (
	"fmt"
	"strconv"
)

/*
 ฟังค์ชั่น ParseUint() แปลงสตริงไปเป็นเลขจำนวนเต็ม แบบไม่คิดเครื่องหมาย เก็บได้เฉพาะจำนวนเต็ม บวก

*/

func main() {

	var str = "-65535"
	var str2 = "65535"
	number, err := strconv.ParseInt(str, 10, 64)
	number2, err := strconv.ParseUint(str2, 10, 32)

	if err == nil {
		fmt.Printf("string \"%s\" with int32 base 10 : %d\n", str, number)
		fmt.Printf("string \"%s\" with uint32 base 10 : %d\n", str, number2)
	} else {
		fmt.Printf("Error is %v\n", err)
	}

}
