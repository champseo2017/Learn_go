package main

import "fmt"

/*

 */
// อธิบายค่า zero value ของ struct
type Person struct {
	Name string
	Age int
}
 
func main() {
	var p Person
	fmt.Println("Name: %q, Age: %d\n", p.Name, p.Age)
}
/* 

*/