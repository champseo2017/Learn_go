package main

import "fmt"

/*
ตัวอย่างการกำหนด Method ให้กับ Type ต่างๆ
*/

// Method ของ Array
type IntArray [5]int

func (a IntArray) Sum() int {
	sum := 0
	for _, value := range a {
		sum += value
	}
	return sum
}

func main() {
	arr := IntArray{1, 2, 3, 4, 5}
	sum := arr.Sum()
	fmt.Println(sum) // Output: 15
}
/* 

*/