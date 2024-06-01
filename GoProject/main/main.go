package main

import "fmt"

/*
เกี่ยวกับฟังก์ชันใน Go:

1. ฟังก์ชันใน Go เป็น First-class สามารถปฏิบัติได้เหมือนกับชนิดข้อมูลอื่นๆ
2. ตัวแปรสามารถ Assign ด้วยฟังก์ชันได้ และฟังก์ชันสามารถส่งผ่านแบบ Dynamic ได้
3. ฟังก์ชันสามารถส่งเป็น Argument และส่งค่ากลับเป็น Return value ได้
4. ฟังก์ชันสามารถเก็บไว้ในตัวแปร รวมถึง Array หรือ Slice ได้
*/

func PrintName(name string) {
	fmt.Println("Person name is", name)
}

func GetPrintPersonDetails(printName func(string), name string, age int) {
	// เรียกใช้ฟังก์ชันที่ส่งเข้ามาเป็น Argument โดยส่งชื่อเป็น Argument
	printName(name)
	fmt.Println("Person age is", age)
}

func main() {
	// ส่งฟังก์ชัน PrintName เป็น Argument ให้กับฟังก์ชัน GetPrintPersonDetails
	GetPrintPersonDetails(PrintName, "Udit", 29)
}

/*
ในตัวอย่าง ฟังก์ชัน `PrintName` ถูกส่งเป็น Argument ให้กับฟังก์ชัน `GetPrintPersonDetails` และถูกเรียกใช้งานภายในฟังก์ชันนั้น แสดงให้เห็นถึงความสามารถในการส่งฟังก์ชันเป็น Argument ในภาษา Go
*/
