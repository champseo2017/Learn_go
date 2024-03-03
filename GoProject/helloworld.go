package main

import (
	"fmt"
)

/*

สตริง "" และ สตริงแบบ backticks


*/

func main() {
   
	// สตริง "" ลำดับของ newline ถูกมองว่าเป็นค่าพิเศษ
	x := "apple\norange"
	fmt.Println(x) // แสดงผลเป็น apple และ orange บนสองบรรทัด

	// สตริงแบบ backticks
	y := `apple\norange`
	fmt.Println(y) // แสดงผลเป็น apple\norange บนบรรทัดเดียว
	
}
