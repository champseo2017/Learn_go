package main

import (
	oops "GoProject/package"
	"fmt"
)

/*
หากเรามี interface ที่ถูก implement โดย type ใดๆ เราสามารถเก็บ object ของ type นั้นไว้ในตัวแปรที่เป็น interface ได้
*/

func main() {
	var s oops.Shape = oops.Square{Side: 10}
	square_area := s.Area()
	fmt.Println(square_area)

	var r oops.Shape = oops.Rectangle{Width: 10, Height: 20}
	rectangle_area := r.Area()
	fmt.Println(rectangle_area)
}

/*
จะเห็นว่าเราสามารถใช้ Shape interface เพื่ออ้างอิงถึง object ของทั้ง Square และ Rectangle ได้ และเรียกใช้ method Area() ผ่านมันได้ตามปกติ โดย Go จะรู้ว่าต้องเรียก Area() ของ type ไหนจาก object ที่ถูกส่งเข้ามาครับ
*/
