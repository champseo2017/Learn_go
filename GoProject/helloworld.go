package main

import (
	"fmt"
)

/*
ประกาศตัวแปรเป็นชนิดข้อมูลแบบ จำนวนเชิงซ้อน
*/

func main() {

	var a float32 = 5
	var b float32 = 10

	x := 20 + 30i
	var y complex64 = complex(1, 2)
	z := complex(a, b)

	addComplex := y + z

	fmt.Printf("Type of x is %T, value of x is %v\n", x, x)
	fmt.Printf("Type of x is %T, value of x is %v\n", y, y)
	fmt.Printf("Type of x is %T, value of x is %v\n", z, z)
	fmt.Printf("y+z = %v\n", addComplex)
	fmt.Println("Real of y + z = ", real(addComplex))
	fmt.Println("Imaginary of y + z = ", imag(addComplex))

}
