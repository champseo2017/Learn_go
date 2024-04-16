package main

import "fmt"

/*
ในภาษา Go เมื่อคุณส่ง array ไปยังฟังก์ชัน มันจะถูกส่งโดยค่า (pass by value) เช่นเดียวกับประเภทข้อมูล int หรือ string ฟังก์ชันจะได้รับเพียงสำเนาของ array ดังนั้น เมื่อคุณทำการเปลี่ยนแปลง array ภายในฟังก์ชัน มันจะไม่ส่งผลกระทบต่อ array เดิม แต่นี่ไม่ใช่กรณีของ slice

ในการส่ง array เป็นอาร์กิวเมนต์ในฟังก์ชัน คุณควรสร้างพารามิเตอร์อย่างเป็นทางการโดยใช้ไวยากรณ์ดังนี้:
- สำหรับ array ที่มีขนาดกำหนด: `func function_name(variable_name [size]type)`
- สำหรับ array ที่ไม่มีขนาดกำหนด: `func function_name(variable_name []type)`
*/

// array ที่ไม่มีขนาดกำหนดในฟังก์ชันเพิ่มเติม
// This function accepts an unsized array as an argument
func arr_sum(a []int) int {
	var i, sum int
	for i = 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

// Main function
func main() {
	// Create and initialize array
	var x = []int{10, 20, 30, 40, 50}
	
	// Pass array as an argument
	sum := arr_sum(x)
	fmt.Printf("Sun of array elements: %d ", sum)
}

/*
ฟังก์ชัน arr_sum ยอมรับ array ที่ไม่มีขนาดกำหนดเป็นอาร์กิวเมนต์ สังเกตว่าในการประกาศฟังก์ชัน เราใช้ []int แทนที่จะเป็น [size]int เหมือนในตัวอย่างก่อนหน้า

ในฟังก์ชัน main เราสร้าง array x โดยไม่ระบุขนาด และส่งมันเป็นอาร์กิวเมนต์ไปยังฟังก์ชัน arr_sum โดยไม่ต้องระบุขนาดของ array

ภายในฟังก์ชัน arr_sum เราใช้ len(a) เพื่อหาความยาวของ array แทนที่จะใช้ขนาดที่กำหนดไว้ล่วงหน้า นี่ทำให้ฟังก์ชันสามารถทำงานกับ array ที่มีขนาดแตกต่างกันได้

โปรแกรมนี้จะให้ผลลัพธ์เหมือนกับตัวอย่างก่อนหน้า แต่แสดงให้เห็นว่าเราสามารถใช้ array ที่ไม่มีขนาดกำหนดกับฟังก์ชันได้อย่างไร ซึ่งทำให้โค้ดมีความยืดหยุ่นมากขึ้น
*/