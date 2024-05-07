package main

import "fmt"

/*
ในภาษา Go อินเทอร์เฟซทำหน้าที่เป็นข้อตกลง (contract) ระหว่างฟังก์ชันและผู้เรียกใช้ฟังก์ชัน หากประเภทใดๆ มีเมธอดตรงตามที่ระบุในอินเทอร์เฟซ ก็สามารถกำหนดค่าให้กับตัวแปรชนิดอินเทอร์เฟซนั้นได้

เมื่อฟังก์ชันมีพารามิเตอร์เป็นอินเทอร์เฟซ เราสามารถส่งค่าของประเภทใดๆ ที่มีเมธอดตรงตามอินเทอร์เฟซนั้นเป็นอาร์กิวเมนต์ได้ ฟังก์ชันจะไม่ทราบรายละเอียดของการ implement จริงๆ ที่ถูกส่งเข้ามา มันรู้แค่ว่าสามารถเรียกเมธอดทั้งหมดของอินเทอร์เฟซได้ และการเรียกจะไปที่การ implement นั้น

ในตัวอย่างโปรแกรม มีการสร้างอินเทอร์เฟซ `Bird` ที่มีเมธอด `Fly()` และมีโครงสร้าง `Eagle`, `Pigeon`, `Penguin` ที่ implement อินเทอร์เฟซ `Bird` โดยการ override เมธอด `Fly()`

จากนั้นมีการกำหนดฟังก์ชัน `flyNow()` ที่รับพารามิเตอร์เป็นอินเทอร์เฟซ `Bird` และเรียกเมธอด `Fly()` บนตัวแปรชนิด `Bird`

ในฟังก์ชัน `main()` มีการเรียกใช้ฟังก์ชัน `flyNow()` โดยส่งค่าของ `Eagle`, `Pigeon`, `Penguin` เป็นอาร์กิวเมนต์ เมื่อฟังก์ชัน `flyNow()` เรียกเมธอด `Fly()` การเรียกจะไปยังการ implement ที่เหมาะสมกับประเภทนั้นๆ นี่คือหลักการของ Polymorphism ผ่านการใช้อินเทอร์เฟซในภาษา Go
*/

type Bird interface {
    Fly()
}

type Eagle struct {}
func (e Eagle) Fly() {
    fmt.Println("Eagle is flying over the cloud")
}

type Pigeon struct {}
func (p Pigeon) Fly() {
    fmt.Println("Pigeon is flying on normal height")
}

type Penguin struct {}
func (p Penguin) Fly() {
    fmt.Println("Penguin cannot fly")
}

func flyNow(b Bird) {
    b.Fly()
}
/* 
จากนั้นมีการกำหนดฟังก์ชัน flyNow() ที่รับพารามิเตอร์เป็นอินเทอร์เฟซ Bird และเรียกเมธอด Fly() บนตัวแปรชนิด Bird ฟังก์ชันนี้ไม่ทราบว่าจะได้รับการ implement ของ Bird แบบใด มันรู้แค่ว่าการ implement นั้นจะมีเมธอด Fly()
*/

func main() {
    flyNow(Eagle{})
    flyNow(Pigeon{})
    flyNow(Penguin{})
}

/* 

*/