package main

import "fmt"

/*
ภาษา Go อนุญาตให้ตั้งชื่อให้กับพารามิเตอร์ที่เป็นค่าคืนกลับ (return หรือ result) ของฟังก์ชันได้ในการประกาศฟังก์ชัน นอกจากนี้ Go ยังมีการรองรับการคืนค่าหลายค่า (multiple return values) ในตัว ความสามารถนี้ถูกใช้เพื่อคืนทั้งค่าผลลัพธ์และค่าข้อผิดพลาด (error) จากฟังก์ชัน
*/

func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = fmt.Errorf("cannot divide by zero")
		return
	}
	result = a / b
	return
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result", result)
	}
}