package main

import (
	"fmt"
)

/*
ประกาศตัวแปรแบบค่าคงที่หลายตัว
*/

func main() {

	const (
		TAX = 0.12
		VAT = 0.1
	)

	const PRODUCT_TAG = "Nat"
	const PRODUCT_QA = 2500
	TAX = 0.15

	fmt.Println("product tag = ", PRODUCT_TAG)
	fmt.Println("product quantity = ", PRODUCT_QA)
	fmt.Println("tax = ", TAX)
	fmt.Println("vat = ", VAT)

}
