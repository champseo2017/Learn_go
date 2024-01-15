package main

import (
	"fmt"
	"reflect"
)

/*
แสดงการทำงานของ type infrernce

var จะประกาศตัวแปรไว้ภายใน หรือ ภายนอก ฟังก์ชั่นก็ได้
:= จะต้องประกาศตัวแปรไว้ภายในฟังก์ชั่นใด ฟังก์ชั่นหนึ่งเท่านั้น

*/
func main() {
	str := "my number"
	num := 28
	status := true

	fmt.Println("type inference of str variable is ", reflect.TypeOf(str))
	fmt.Println("type inference of num variable is ", reflect.TypeOf(num))
	fmt.Println("type inference of status variable is ", reflect.TypeOf(status))

}