package main

import "fmt"

/*

การกำหนดเมธอดบนฟังก์ชันในภาษา Go:

1. Go อนุญาตให้กำหนดเมธอดบนชนิดฟังก์ชันได้ ซึ่งมีประโยชน์เมื่อต้องการใช้ฟังก์ชันัลลิตี้ที่คล้ายกันกับหลายๆ สถานการณ์

2. ในตัวอย่างโค้ด มีการสร้างชนิดฟังก์ชัน `By` ที่รับ `User` เป็นพารามิเตอร์และคืนค่าเป็น `bool` และมีเมธอด `Filter` ที่รับสไลซ์ของ `User` และคืนค่าเป็นสไลซ์ของ `User` ที่ถูกกรองแล้ว

3. `Filter` ใช้ฟังก์ชัน `By` เพื่อตรวจสอบเงื่อนไขการกรองสำหรับแต่ละสมาชิกในสไลซ์ของ `User` และคืนค่าเฉพาะสมาชิกที่ผ่านเงื่อนไข

4. ในฟังก์ชัน `main` มีการสร้างฟังก์ชันที่ใช้เงื่อนไขการกรองที่แตกต่างกัน (`frequent`, `appreciated`, `respected`, `matured`) แล้วแปลงเป็นชนิด `By` และเรียกใช้เมธอด `Filter` กับสไลซ์ของ `User`

5. การกำหนดเมธอดบนชนิดฟังก์ชันช่วยให้สามารถใช้โค้ดซ้ำได้ โดยไม่ต้องเขียนฟังก์ชันกรองแยกสำหรับแต่ละเงื่อนไข และทำให้โค้ดอ่านง่ายและมีโครงสร้างที่ดีขึ้น

สรุปแล้ว การกำหนดเมธอดบนชนิดฟังก์ชันเป็นเทคนิคที่มีประโยชน์ในการเขียนโค้ดที่ยืดหยุ่นและนำกลับมาใช้ใหม่ได้ในภาษา Go โดยเฉพาะเมื่อต้องการใช้ฟังก์ชันัลลิตี้ที่คล้ายกันกับข้อมูลหรือเงื่อนไขที่แตกต่างกัน

*/

type User struct {
    id         int
    name       string
    age        int
    visitDays  int
    totalLikes int
    followers  int
}

var users []User = []User{
    {101, "A", 18, 20, 645, 2342},
    {102, "B", 23, 110, 323, 110},
    {103, "C", 40, 125, 1120, 4577},
    {104, "D", 36, 45, 323, 1201},
    {105, "D", 42, 45, 323, 1201},
}

type By func(user User) bool

func (by By) Filter(s1 []User) []User {
    filtered := make([]User, 0)
    for _, ele := range s1 {
        if by(ele) {
            filtered = append(filtered, ele)
        }
    }
    return filtered
}

func main() {
	//Frequent User
    frequent := func(user User) bool {
        return user.visitDays > 100
    }

    //Liked User
    appreciated := func(user User) bool {
        return user.totalLikes > 500
    }

    // Large followers
    respected := func(user User) bool {
        return user.followers > 1000
    }

    // Matured User
    matured := func(user User) bool {
        return user.age > 35
    }

    frequestUsers := By(frequent).Filter(users)
    fmt.Println(frequestUsers)

    appreciatedUsers := By(appreciated).Filter(users)
    fmt.Println(appreciatedUsers)

    respectedUsers := By(respected).Filter(users)
    fmt.Println(respectedUsers)

    maturedUsers := By(matured).Filter(users)
    fmt.Println(maturedUsers)
}
/* 

*/