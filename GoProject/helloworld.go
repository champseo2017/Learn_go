package main

import (
	"fmt"
)

/*
แสดงการประกาศตัวแปรหลายตัว
*/

func main() {

	var (
		prd_id1   = "D001"
		prd_name1 = "T-Shirt"
		price1    = 250
	)

	var (
		prd_id2   = "D002"
		prd_name2 = "Shoe"
		price2    = 499
	)

	fmt.Println("Product : ", prd_id1, " | ", prd_name1, " | ", price1, "Baht")
	fmt.Println("product : ", prd_id2, " | ", prd_name2, " | ", price2, "Baht")
}
