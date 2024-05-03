package main

import "fmt"

/*

 */
// struct เป็น value type
type Person struct {
	Name string
	Age int
}
 
func main() {
	p1 := Person{Name: "John", Age: 30}
	p2 := p1
	p2.Name = "Alice"
	fmt.Println(p1.Name)
	fmt.Println(p2.Name)
}
/* 

*/