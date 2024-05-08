package main

import "fmt"

/*
Stringer interface เป็น interface ใน standard library ของ Go ที่ใช้บ่อย มีเพียงเมธอดเดียวคือ String() ซึ่งคืนค่าเป็น string เมื่อเรา implement Stringer interface กับ type ใดๆ และใช้ fmt.Println() เพื่อแสดงค่าของตัวแปรที่มี type นั้น เมธอด String() จะถูกเรียกใช้โดยอัตโนมัติ และแสดงผลตามที่เมธอดคืนค่ามา

ตัวอย่างโค้ดแสดงการสร้าง type ใหม่ชื่อ Dollar และ implement Stringer interface เพื่อแสดงค่าของ Dollar ในรูปแบบที่กำหนดเอง คือมีเครื่องหมาย $ นำหน้า เมื่อแสดงค่าตัวแปรชนิด Dollar ด้วย fmt.Println() ผลลัพธ์ที่ได้จะเป็นไปตามที่กำหนดใน String() method
*/

type Dollar float64

func (d Dollar) String() string {
    return fmt.Sprintf("$%f", d)
    /* 
        fmt.Sprintf เป็นฟังก์ชันใน package fmt ของ Go ที่ใช้สำหรับจัดรูปแบบข้อความ (string formatting) โดยมีหน้าที่หลักดังนี้

        1. รับ format string ซึ่งเป็นแม่แบบของข้อความที่ต้องการ โดยใน format string จะมีตัวแทน (placeholder) สำหรับค่าที่ต้องการแทรกลงไป เช่น %s สำหรับ string, %d สำหรับ integer, %f สำหรับ float เป็นต้น

        2. รับค่าที่ต้องการแทรกลงใน format string โดยระบุตามลำดับหลังจาก format string

        3. จัดรูปแบบข้อความตาม format string และค่าที่ระบุ แล้วคืนค่าเป็น string ที่จัดรูปแบบแล้ว
    */
}

func main() {
  var d Dollar = 23.3
  fmt.Println(d)
}
