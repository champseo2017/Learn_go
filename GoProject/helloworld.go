package main

import (
	"fmt"
)

/*
function Scanln() รับข้อมูลอินพุตจากหน้าจอโปรแกรม

*/

func main() {

	fmt.Print("Please enter your")
	var name, surname, dept string
	count, err := fmt.Scanln(&name, &surname, &dept)
	fmt.Printf("Number of input is \"%v\"\n", count)
	if err == nil {
		fmt.Printf("Hello 1")
	} else {
		fmt.Printf(" Req 3")
		fmt.Printf("Error is \"%v\"\n", err)
	}
}
