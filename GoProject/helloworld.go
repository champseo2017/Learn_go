package main

import (
	"fmt"
)

/*
การทำ Type Conversion

*/

func main() {

	var x int = 5
	var y int = 10
	var result float32

	// explicit type conversion
	result = float32(x) * float32(y)

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	fmt.Println("Multiplication of x and y = ", result)
}
