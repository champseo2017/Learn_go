package main

import (
	"fmt"
)

/*
ตัวอย่างการกำหนด Method ให้กับ Type ต่างๆ
*/

// Method ของ Map
type StringMap map[string]string

func (m StringMap) GetValue(key string) string {
	return m[key]
}

func main() {
	dict := StringMap{"apple": "fruit", "carrot": "vegetable"}
	value := dict.GetValue("apple")
	fmt.Println(value) // Output: fruit
}
/* 

*/