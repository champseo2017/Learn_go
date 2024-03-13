package main

import (
	"fmt"
)

// ประเภทโครงสร้าง (Structure types)
type Person struct {
    Name string
    Age int
}


func main() {
  // สร้างตัวแปร p ประเภท Person และกำหนดค่าให้กับฟิลด์
  p := Person{Name: "John Doe", Age: 30}
  // แสดงค่าของ p ออกมา
  fmt.Println(p)
  // การเข้าถึงฟิลด์ของโครงสร้างและแสดงค่า
  fmt.Printf("%s is %d years old.\n", p.Name, p.Age)
}