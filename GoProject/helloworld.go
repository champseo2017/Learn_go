package main

import "fmt"

// ประเภทอินเทอร์เฟซ (Interface types)
type Greeter interface {
    Greet() string
}

// ประกาศโครงสร้าง EnglishSpeaker ที่จะเป็นหนึ่งในผู้รับใช้งานอินเทอร์เฟซ Greeter
type EnglishSpeaker struct {}
// ประกาศเมธอด Greet สำหรับ EnglishSpeaker ที่ตอบกลับด้วยข้อความ "Hello!"
/* 
func: คีย์เวิร์ดที่ใช้ในการประกาศฟังก์ชันหรือเมธอดใน Go
(e EnglishSpeaker): ส่วนนี้เรียกว่า "receiver" ซึ่งระบุว่าเมธอดนี้เป็นของโครงสร้าง EnglishSpeaker. ตัวแปร e ในที่นี้เป็นอินสแตนซ์ของ EnglishSpeaker ที่เมธอด Greet จะทำงานกับมัน
Greet(): ชื่อของเมธอด
string: ประเภทของค่าที่เมธอดนี้คืนกลับ ในที่นี้คือ string
*/
func (e EnglishSpeaker) Greet() string {
    return "Hello!"
}

// ประกาศโครงสร้าง ThaiSpeaker ที่จะเป็นหนึ่งในผู้รับใช้งานอินเทอร์เฟซ Greeter
type ThaiSpeaker struct{}

// ประกาศเมธอด Greet สำหรับ ThaiSpeaker ที่ตอบกลับด้วยข้อความ "สวัสดี!"
func (t ThaiSpeaker) Greet() string {
    return "สวัสดี!"
}

// ฟังก์ชันที่รับอินเทอร์เฟซ Greeter เป็นพารามิเตอร์ และแสดงข้อความที่ได้จากเมธอด Greet
func greetSomeone(g Greeter) {
    fmt.Println(g.Greet())
}

func main() {
   englishSpeaker := EnglishSpeaker{}
   thaiSpeaker := ThaiSpeaker{}

   // เรียกใช้ฟังก์ชัน greetSomeone กับ EnglishSpeaker
   greetSomeone(englishSpeaker)
   // เรียกใช้ฟังก์ชัน greetSomeone กับ ThaiSpeaker
   greetSomeone(thaiSpeaker)
}