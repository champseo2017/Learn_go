package main

import "fmt"

// ประกาศตัวแปร p, q, r ที่ระดับ package พร้อมกำหนดค่าเริ่มต้น
var p, q, r int = 150, 250, 300

func main() {
    // ประกาศและกำหนดค่าเริ่มต้นให้กับตัวแปร a ในฟังก์ชัน main
    a := 10
    // ประกาศและกำหนดค่าเริ่มต้นให้กับตัวแปร x และ y ในฟังก์ชัน main
    x, y := 100, 200
    // ไม่จำเป็นต้องประกาศ p, q, r ใหม่ เพราะมีการประกาศไว้นอก main แล้ว
    // แสดงค่าของตัวแปร a, x, y, p, q, r ออกมา
    fmt.Println("a:", a)
    fmt.Println("x:", x, "y:", y)
    fmt.Println("p:", p, "q:", q, "r:", r)
}