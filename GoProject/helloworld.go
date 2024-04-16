package main

import "fmt"

/*
ในภาษา Go เมื่อคุณส่ง array ไปยังฟังก์ชัน มันจะถูกส่งโดยค่า (pass by value) เช่นเดียวกับประเภทข้อมูล int หรือ string ฟังก์ชันจะได้รับเพียงสำเนาของ array ดังนั้น เมื่อคุณทำการเปลี่ยนแปลง array ภายในฟังก์ชัน มันจะไม่ส่งผลกระทบต่อ array เดิม แต่นี่ไม่ใช่กรณีของ slice

ในการส่ง array เป็นอาร์กิวเมนต์ในฟังก์ชัน คุณควรสร้างพารามิเตอร์อย่างเป็นทางการโดยใช้ไวยากรณ์ดังนี้:
- สำหรับ array ที่มีขนาดกำหนด: `func function_name(variable_name [size]type)`
- สำหรับ array ที่ไม่มีขนาดกำหนด: `func function_name(variable_name []type)`
*/

// This function accepts array as an argument
func arr_avg(a [4]int, size int) int {
	var i, sum, avg int
	for i = 0; i < size; i++ {
		sum += a[i]
	}
	avg = sum / size
	return avg
}

// Main function
func main() {
	// Create and initialize array
	var x = [4]int{20, 40, 60, 80}
	// Pass array as an argument
	avg := arr_avg(x, 4)
	fmt.Printf("Average of array elements: %d ", avg)
}

/*
ใช้ฟังก์ชันชื่อ arr_avg() ซึ่งยอมรับ array เป็นอาร์กิวเมนต์ ในฟังก์ชัน main เราส่ง x[4] ของประเภท int ไปยังฟังก์ชันพร้อมกับขนาดของ array และฟังก์ชันจะคืนค่าเฉลี่ยขององค์ประกอบใน array
*/