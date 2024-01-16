package main

import (
	"fmt"
)

/*
ตัวแปรแบบโกลบอล และ โลคอลของ บล็อกการทำงาน
*/

func main() {

	var myvar1 int = 10
	{
		var myvar2 int = 20
		fmt.Println("variable1 is ", myvar1)
		fmt.Println("variable2 is ", myvar2)
	}

}
