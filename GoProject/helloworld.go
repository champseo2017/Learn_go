package main

import "fmt"

/*
ในภาษา Go เมื่อคุณส่ง array ไปยังฟังก์ชัน มันจะถูกส่งโดยค่า (pass by value) เช่นเดียวกับประเภทข้อมูล int หรือ string ฟังก์ชันจะได้รับเพียงสำเนาของ array ดังนั้น เมื่อคุณทำการเปลี่ยนแปลง array ภายในฟังก์ชัน มันจะไม่ส่งผลกระทบต่อ array เดิม แต่นี่ไม่ใช่กรณีของ slice

ในการส่ง array เป็นอาร์กิวเมนต์ในฟังก์ชัน คุณควรสร้างพารามิเตอร์อย่างเป็นทางการโดยใช้ไวยากรณ์ดังนี้:
- สำหรับ array ที่มีขนาดกำหนด: `func function_name(variable_name [size]type)`
- สำหรับ array ที่ไม่มีขนาดกำหนด: `func function_name(variable_name []type)`
*/

func arr_sum(a [5]int, size int) int {
	var i, sum int
	for i = 0; i < size; i++ {
		sum += a[i]
	}
	return sum
}

func main() {
	// Create and initialize array
	var x = [5]int{10, 20, 30, 40, 50}

	// Pass array as an argument
	sum := arr_sum(x, 5)
	fmt.Printf("Sum of array elements: %d", sum)
}

/*
โปรแกรมนี้มีฟังก์ชันชื่อ arr_sum() ซึ่งยอมรับ array เป็นอาร์กิวเมนต์ ในฟังก์ชัน main เราส่ง x[5] ของประเภท int ไปยังฟังก์ชันพร้อมกับขนาดของ array และฟังก์ชันจะคืนค่าผลรวมขององค์ประกอบใน array
*/