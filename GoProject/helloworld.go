package main

import "fmt"

/*
ในภาษา Go เราสามารถใช้ for loop เพื่อวนซ้ำผ่าน array ได้ นอกจากนี้ Go ยังมีตัวดำเนินการ range ซึ่งจะคืนค่าดัชนีและค่าของแต่ละองค์ประกอบใน array ในลูป for
*/

func main() {
	x := [...]int{10, 20, 30, 40, 50}
	for index, value := range x {
		fmt.Printf("x[%d] = %d\n", index, value)
	}
}

/*
โปรแกรมนี้คล้ายกับโปรแกรมก่อนหน้า แต่ใช้ตัวดำเนินการ range เพื่อวนซ้ำผ่าน array แทน for loop แบบปกติ range จะคืนค่าดัชนีและค่าของแต่ละองค์ประกอบใน array
*/