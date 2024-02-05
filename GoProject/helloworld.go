package main

import (
	"fmt"
)

/*
แม่ 2 และ แม่ 3 ใช้คำสั่ง Nested-for

*/

func main() {

	for i:= 2; i <=3; i++ {
		fmt.Printf("%d times table\n", i)
		for j := 1; j <= 12; j++ {
			fmt.Printf("%d*%d \t = \t %d\n", i, j, i * j)
		}
		fmt.Println()
	}
}
