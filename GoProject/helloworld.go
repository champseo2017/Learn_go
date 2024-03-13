package main

import (
	"fmt"
)



func main() {
   // ประกาศค่าคงที่ const
   // const Pi = 3.14
  // ประเภทสไลซ์ (Slice types)
  // ประกาศและกำหนดค่าให้กับสไลซ์ numbers
  var numbers = []int{1, 2, 3, 4, 5}
  // แสดงค่าของสไลซ์ numbers
  fmt.Println("Numbers:", numbers)
  // เพิ่มค่าเข้าไปในสไลซ์
  numbers = append(numbers, 6)
  fmt.Println("Numbers after appending 6:", numbers)
  // การเข้าถึงค่าในสไลซ์ด้วย index
  fmt.Println("The first number is:", numbers[0])
  // การเข้าถึงส่วนย่อยของสไลซ์
  subSlice := numbers[1:4] // ค่าที่ index 1 ถึง index 3 (ไม่รวม index 4)
  fmt.Println("Sub-slice from index 1 to 3:", subSlice)
  // การเปลี่ยนค่าในสไลซ์
  numbers[0] = 10
  fmt.Println("Numbers after changing the first element to 10:", numbers)

}