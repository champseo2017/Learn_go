package main

import "fmt"

/*

 */
// ตรวจสอบความเท่ากันของ struct

type Person struct {
	Name string
	Age int
}
 
func main() {
	p1 := Person{Name: "John", Age: 30}
	p2 := Person{Name: "John", Age: 30}
	fmt.Println(p1 == p2)
}
/* 

*/