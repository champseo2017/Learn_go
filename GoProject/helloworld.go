package main

import (
	"fmt"
)

/*
การทำงานของตัวดำเนินการเปรียบเทียบ

*/

func main() {

	var x int = 100
	var y int = 50
	var z int = 100

	fmt.Println("x == z is ", x == z)
	fmt.Println("x != y is ", x != y)
	fmt.Println("x > y is ", x > y)
	fmt.Println("x < y is ", x < y)
	fmt.Println("x >= z is ", x >= z)
	fmt.Println("x <= y is ", x <= y)

}
