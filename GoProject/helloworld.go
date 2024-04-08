package main

import (
	"fmt"
)

/*
ไวยากรณ์ในการสร้างฟังก์ชันในภาษา Go:
func functionName(parameter1 type1, parameter2 type2) (returnType1, returnType2) {
    // function body
    // statements
    return value1, value2
}

ความแตกต่างระหว่าง multiple returns และ named returns:
func multipleReturns(x, y int) (int, int) {
    return x + y, x - y
}

func namedReturns(x, y int) (sum, diff int) {
    sum = x + y
    diff = x - y
    return
}

ตัวอย่างโปรแกรม Go ที่อธิบาย call by value:
func callByValue(x int) {
    x = x + 1
    fmt.Println("Inside function:", x)
}

func main() {
    x := 10
    callByValue(x)
    fmt.Println("Outside function:", x)
}

ฟังก์ชันแบบ variadic:
func variadicFunction(numbers ...int) {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    fmt.Println("Sum:", sum)
}

func main() {
    variadicFunction(1, 2, 3)
    variadicFunction(4, 5, 6, 7)
}

การใช้งานคำสั่ง defer:
func deferExample() {
    defer fmt.Println("Deferred message")
    fmt.Println("Normal message")
}

func main() {
    deferExample()
}
*/

func deferExample() {
	defer fmt.Println("Deferred message")
	fmt.Println("Normal message")
}

func main() {
	deferExample()
}