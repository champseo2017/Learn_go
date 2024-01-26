package main

import (
	"fmt"
)

/*
การทำงานของตัวดำเนินการ ทางตรรกศาสตร์

*/

func main() {

	var x int = 25
	var y int = 30

	result1 := x < y && y == 30
	result2 := x > y || y == 30
	result3 := !(y == 30)
	fmt.Println("x < y && y == 30 is", result1)
	fmt.Println("x > y || y == 30 is", result2)
	fmt.Println("!(y == 30) is", result3)

}
