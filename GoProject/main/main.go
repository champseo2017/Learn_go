package main

import "fmt"

/*

 */
// อธิบายแนวคิดของพอยน์เตอร์ใน struct

type Person struct {
	Name string
	Age int
}
 
func main() {
	p := Person{Name: "John", Age: 30}
	ptr := &p
	fmt.Println(ptr.Name)
	ptr.Age = 31
	fmt.Println(p.Age)
}
/* 

*/