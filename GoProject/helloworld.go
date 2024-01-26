package main

import (
	"fmt"
)

/*
ใช้ Scan() รับข้อมูลอินพุตจากหน้าจอโปรแกรม

newline ถูกนับเป็นช่องว่าง โปรแกรมกำหนดให้รับ 2 input หากป้อนค่าไม่ครบ newline ถือว่าเป็นการขึ้นบรรทัดใหม่

*/

func main() {

	fmt.Print("Pleases enter your name :")
	var name, surname string
	fmt.Scan(&name, &surname)
	fmt.Printf("Hello, %v %v!\n", name, surname)
}
