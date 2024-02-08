package main

import (
	"fmt"
	"strconv"
)

/*
 ParseFloat() แปลงค่าสตริงไปเป็นเลขทศนิยม

*/

func main() {

	var str = "3.14156"
	number, err := strconv.ParseFloat(str, 64)
	
	if err == nil {
		fmt.Printf("string \"%s\" with int32 base 10 : %f\n", str, number)
	} else {
		fmt.Printf("Error is %v\n", err)
	}

}
