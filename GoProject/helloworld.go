package main

import (
	"fmt"
	"strconv"
)

/*
 แปลงค่าเลขจำนวนเต็มไปเป็นสตริง

 Itoa()
 FormatInt()
 FormatUint()

*/

func main() {

	s1 := strconv.Itoa(97)
	s2 := strconv.FormatInt(97, 10)
	s3 := strconv.FormatInt(-97, 16)
	s4 := strconv.FormatUint(97, 16)

	fmt.Printf("Itoa() \"%s\"\n", s1)
	fmt.Printf("FormatInt() \"%s\"\n", s2)
	fmt.Printf("FormatInt() \"%s\"\n", s3)
	fmt.Printf("FormatUint() \"%s\"\n", s4)

}
