package main

import (
	"fmt"
)

/*
ตัวอย่างการกำหนด Method ให้กับ Type ต่างๆ
*/

// Method ของ Slice
type IntSlice []int

func (s IntSlice) Average() float64 {
	sum := 0
	for _, value := range s {
		sum += value
	}
	return float64(sum) / float64(len(s))
}

func main() {
	slices := IntSlice{1, 2, 3, 4, 5}
	avg := slices.Average()
	fmt.Println(avg) // Output: 3
}
/* 

*/