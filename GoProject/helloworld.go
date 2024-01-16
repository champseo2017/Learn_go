package main

import (
	"fmt"
)

/*
ขอบเขตของตัวแปร
*/

// global_var เป็น global variable
var global_var int = 10

func main() {

	// local_var เป็น local variable
	var local_var int = 20

	fmt.Println("global variable is ", global_var)
	fmt.Println("local variable is ", local_var)
}

func show() {
	fmt.Println("global variable is ", global_var)
}
