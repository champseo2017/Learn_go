package main

import "fmt"

/*
ในภาษา Go การกำหนดฟังก์ชันมีรูปแบบดังนี้
*/

func add(x int, y int) int {
	return x + y
}

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	
}
