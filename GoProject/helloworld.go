package main

import (
	"fmt"
)

/*
แสดงการประกาศตัวแปรหลายตัว
*/
var prd_id1, prd_id2 = "D001", "D002"

func main() {
	prd_name1, price1 := "T-Shirt", 250
	prd_name2, price2 := "Shoe", 499
	fmt.Println("Product : ", prd_id1, " | ", prd_name1, " | ", price1, "Baht")
	fmt.Println("product : ", prd_id2, " | ", prd_name2, " | ", price2, "Baht")
}
