package main

import "fmt"

/*
ตัวอย่างการใช้ฟังก์ชัน new() เพื่อสร้างพอยน์เตอร์ในภาษา Go
*/

func main() {
   // สร้างพอยน์เตอร์ชี้ไปยังค่า int โดยใช้ฟังก์ชัน new()
   ptr := new(int)
    
   // กำหนดค่าให้กับตำแหน่งที่พอยน์เตอร์ชี้ไป
   *ptr = 42
   
   fmt.Println("Value at pointer:", *ptr)  // แสดงค่าที่พอยน์เตอร์ชี้ไป
   fmt.Println("Pointer address:", ptr)    // แสดงที่อยู่ของพอยน์เตอร์
   
   // สร้างพอยน์เตอร์ชี้ไปยังค่า string โดยใช้ฟังก์ชัน new()
   strPtr := new(string)
   
   // กำหนดค่าให้กับตำแหน่งที่พอยน์เตอร์ชี้ไป
   *strPtr = "Hello, World!"
   
   fmt.Println("Value at pointer:", *strPtr)  // แสดงค่าที่พอยน์เตอร์ชี้ไป
   fmt.Println("Pointer address:", strPtr)    // แสดงที่อยู่ของพอยน์เตอร์
}
/* 

*/