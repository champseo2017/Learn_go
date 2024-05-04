package main

import "fmt"

/*
ตัวอย่างการกำหนด Method ให้กับ Type ต่างๆ
*/

// Method ของ Basic Type (int)
type MyInt int

func (m MyInt) Double() int {
	return int(m * 2)
}

func main() {
	num := MyInt(5)
	result := num.Double()
	fmt.Println(result)
}
/* 

*/