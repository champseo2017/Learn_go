package main

/*
1. เพื่อสร้าง custom error type เราต้อง define method ชื่อ `Error()` ในtype นั้นๆ

2. ข้อความใดเกี่ยวกับ `defer` ถูกต้อง?
   ตอบ: Defer function จะถูก execute เมื่อ parent function ทำงานจบ เมื่อ parent function จบการทำงานแล้วเท่านั้น defer function จึงจะทำงาน

3. จำเป็นไหมที่ต้องเรียกใช้ฟังก์ชัน `recover` ใน defer function?
   ตอบ: ใช่ จำเป็นต้องเรียกใช้ `recover` ใน defer function ถ้าต้องการดักจับและจัดการ panic ที่อาจเกิดขึ้นใน parent function
*/

func main() {

}

/*

 */
