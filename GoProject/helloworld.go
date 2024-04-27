package main

import "fmt"

/*

 */

func main() {
	// ในภาษา Go มีวิธีการประกาศ map โดยใช้ syntax ดังนี้
	// var m map[key_type]value_type
	// เราสามารถประกาศ map ที่มี key เป็น string และ value เป็น int
	// var m map[string]int
	/* 
	ค่า zero value ของ map คือ nil ซึ่ง map ที่เป็น nil จะไม่มี key ใดๆ ดังนั้น ถ้าเราพยายามเพิ่ม key ลงใน map ที่เป็น nil จะเกิด runtime error ขึ้น
	*/
	var m1 map[string]int
	fmt.Println(m1)

	if m1 == nil {
		/* 
		คำเตือน "tautological condition: nil == nilnilnesscond" ในภาษา Go หมายความว่า เงื่อนไขในคำสั่ง if เป็นเงื่อนไขที่มีผลลัพธ์เป็นจริงเสมอ (tautology) และไม่มีประโยชน์
		*/
		fmt.Println("Map is empty")
	} else {
		fmt.Println("Map is not empty")
	}
	// Attempt to add keys to a nil map will throw runtime error
    //m1["ten"] = 10

}
/* 
- ประกาศตัวแปร `m1` เป็น map ที่มี key เป็น string และ value เป็น int
- แสดงค่าของ `m1` ออกมา ซึ่งจะเป็น `map[]` เพราะ `m1` ยังไม่ได้ถูกกำหนดค่า (เป็น nil)
- ตรวจสอบว่า `m1` เป็น nil หรือไม่ ถ้าเป็น nil จะแสดงข้อความ "Map is empty" ถ้าไม่เป็น nil จะแสดงข้อความ "Map is not empty"
- ถ้าเราลองเพิ่ม key ลงใน `m1` เช่น `m1["ten"] = 10` (ตอนนี้ถูก comment ไว้) จะเกิด runtime error เพราะเรากำลังพยายามเพิ่ม key ลงใน map ที่เป็น nil

ดังนั้น ก่อนที่จะเพิ่ม key-value ลงใน map เราจำเป็นต้องเริ่มต้น (initialize) map ก่อนเสมอ ไม่เช่นนั้นจะเกิด runtime error ขึ้น
*/