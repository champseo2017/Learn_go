package main

import "fmt"

/*

 */
// ประกาศและกำหนดค่าเริ่มต้นให้กับ struct
 type Person struct {
	Name string
	Age int
 }

func main() {
	p := Person{Name: "John", Age: 30}
	fmt.Println(p)
}
/* 

*/