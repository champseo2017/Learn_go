package main

import (
	"fmt"
)

/*
Sscanf() สแกนอ่านข้อมูลจาก สตริงที่กำหนด

fmt.Sscanf("string to parse", "format string", &var1, &var2, ...)

*/

func main() {

	var dept_id, dept_name, emp string
	var num int

	fmt.Sscanf("82201 : IT", "%s : %s", &dept_id, &dept_name)
	fmt.Sscanf("82201 : IT", "%s : %s IT", &num, &emp)

	fmt.Printf("IT : %s\n", dept_name)

}
