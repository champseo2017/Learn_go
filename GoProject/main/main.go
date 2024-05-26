package main

import (
	"errors"
	"fmt"
)

/*

เกี่ยวกับการจัดการ Error ในภาษา Go:

- ภาษา Go เรียกข้อผิดพลาดที่ไม่คาดคิดว่า "Error" แทนที่จะเป็น "Exception" เหมือนในภาษาอื่น ๆ
- ฟังก์ชันหรือเมธอดสามารถส่งคืน Error ได้ และผู้เรียกใช้ต้องตรวจสอบและจัดการกับ Error เหล่านั้น
- Go มี Interface ของ Error ในตัว และแพ็คเกจ Error ที่ให้เมธอด ฟังก์ชัน และ Struct สำหรับจัดการ Error
- สามารถสร้าง Error ใหม่หรือกำหนด Error Type เองได้
- Panic คือสถานการณ์ที่โปรแกรมไม่สามารถดำเนินการต่อไปได้ และสามารถจัดการ Panic ได้โดยใช้ Recover
- Defer เป็นคำสั่งที่ใช้เพื่อให้ฟังก์ชันทำงานหลังจากที่ฟังก์ชันปัจจุบันจบการทำงานแล้ว ซึ่งมีประโยชน์ในการจัดการทรัพยากรหรือการทำความสะอาด

การจัดการ Error ในภาษา Go จะใช้การตรวจสอบค่า Error ที่ส่งคืนจากฟังก์ชันหรือเมธอด ถ้ามี Error เกิดขึ้นก็จะแสดงข้อความ Error แต่ถ้าไม่มีก็จะดำเนินการต่อไปตามปกติ ซึ่งแตกต่างจากภาษาอื่น ๆ ที่ใช้ try/catch ในการจัดการ Exception
*/

// ตัวอย่างการจัดการ Error ในภาษา Go
func divide(a, b int) (int, error) {
	if b == 0 {
		// ถ้า b เป็น 0 จะ return error
		return 0, errors.New("cannot divide by zero")
	}
	// ถ้า b ไม่เป็น 0 จะ return ผลหารและ nil error
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		// ถ้ามี error เกิดขึ้น จะแสดงข้อความ error
		fmt.Println("Error:", err)
	} else {
		// ถ้าไม่มี error จะแสดงผลลัพธ์
		fmt.Println("Result:", result)
	}

	// เรียกใช้ฟังก์ชัน divide โดยให้ b เป็น 0 เพื่อทดสอบการจัดการ error
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

/*

 */
