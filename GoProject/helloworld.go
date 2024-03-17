package main

// Interface ใน Go เป็นชุดของ method signatures ที่ไม่มีการเนื้อหาหรือการทำงานของ method ใดๆ กำหนดไว้
type Greeter interface {
    Greet() string // ตัวอย่างของ method signature ใน interface
}

// Struct ใน Go เป็นชุดของ fields ที่สามารถมีทั้งข้อมูลและ methods
// สามารถใช้ struct เพื่อ "implement" interface โดยการเขียน method ที่ตรงกับ signature ใน interface
type EnglishSpeaker struct {} // ตัวอย่างของ struct ที่ไม่มี fields

// การ implement interface ใน struct
func (es EnglishSpeaker) Greet() string {
    return "Hello" // ตัวอย่างของการ implement method จาก interface ใน struct
}

/* 
- Interface ใช้สำหรับกำหนดชุดของ method signatures ที่ struct หรือ type อื่นๆ ต้อง implement
- Struct ใช้สำหรับกำหนดข้อมูลและสามารถมี methods ที่ implement จาก interface หรือเป็น methods ของตัวมันเอง
- การใช้ interface ช่วยให้สามารถเขียนโค้ดที่ flexible และ reusable มากขึ้น โดยไม่ต้องกังวลเกี่ยวกับรายละเอียดการ implement ของแต่ละ type

*/

func main() {
   
}