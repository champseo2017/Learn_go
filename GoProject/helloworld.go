package main

import (
	"fmt"
)

/*
ประกาศตัวแปรเป็นชนิดข้อมูลแบบเลขทศนิยม

%v พิมพ์ค่าข้อมูลของตัวแปรในรูปแบบ default format
%.2f พิมพ์ค่าข้อมูลเลขทศนิยมแบบแสดงจุดทศนิยม 2 หลักโดยมีการปัดเศษ จาก .65 เป็น .66
*/

func main() {

	var myFloat = 587.6599
	fmt.Printf("Type of myFloat is %T \n", myFloat)
	fmt.Printf("Type of myFloat is %v (%.2f)\n", myFloat, myFloat)

}
