package main

import (
	"fmt"
	"reflect"
)

/*
เราสามารถพิมพ์ประเภทของตัวแปรใน Golang ได้อย่างไร?
*/

func main() {
	var x = 10.5
	fmt.Println("Type of x:", reflect.TypeOf(x))
}
