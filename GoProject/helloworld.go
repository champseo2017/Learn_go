package main

import "fmt"

/*
ในภาษา Go เราสามารถใช้ for loop เพื่อวนซ้ำผ่าน array ได้ นอกจากนี้ Go ยังมีตัวดำเนินการ range ซึ่งจะคืนค่าดัชนีและค่าของแต่ละองค์ประกอบใน array ในลูป for
*/

func main() {
	x := [...]int{10, 20, 30, 40, 50}
	for index := 0; index < len(x); index++ {
		fmt.Printf("x[%d] = %d\n", index, x[index])
	}
}

/*
โปรแกรมนี้ประกาศตัวแปร x เป็น array ของจำนวนเต็ม จากนั้นใช้ for loop เพื่อวนซ้ำผ่าน array โดยใช้ตัวแปร index เป็นดัชนี และแสดงค่าดัชนีและค่าขององค์ประกอบที่สอดคล้องกันโดยใช้ fmt.Printf()
*/