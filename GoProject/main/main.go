package main

import "fmt"

/*

 */
// เข้าถึงฟิลด์ของ struct

type Person struct {
	Name string
	Age int
}
 
func main() {
	p := Person{Name: "John", Age: 30}
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
/* 

*/