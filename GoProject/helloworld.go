package main

import "fmt"

/*
การประกาศอาร์เรย์แบบสั้น (shorthand declaration) เป็นอีกวิธีหนึ่งในการประกาศอาร์เรย์ในภาษา Go โดยมีรูปแบบดังนี้

array_name := [length]Type{item1, item2, item3, ..., itemN}

โดย array_name คือชื่อของอาร์เรย์ length คือขนาดของอาร์เรย์ Type คือชนิดข้อมูลของอาร์เรย์ และ item1 ถึง itemN คือค่าเริ่มต้นของสมาชิกในอาร์เรย์


*/

func main() {
	arr := [5]string{"Ganga", "Yamuna", "Godavari", "Kaveri", "Narmada"}

	fmt.Println("Array elements:")

	for i := 0; i <= 4; i++ {
		fmt.Println(arr[i])
	}
}

/*
ในตัวอย่างข้างต้น เราใช้การประกาศอาร์เรย์แบบสั้นเพื่อประกาศอาร์เรย์ arr ที่มีขนาด 5 และชนิดข้อมูลเป็น string พร้อมทั้งกำหนดค่าเริ่มต้นให้กับสมาชิกในอาร์เรย์เป็น {"Ganga", "Yamuna", "Godavari", "Kaveri", "Narmada"}

จากนั้นเราใช้คำสั่ง fmt.Println("Array elements:") เพื่อแสดงข้อความ "Array elements:" บนหน้าจอ

ต่อมาเราใช้ลูป for เพื่อวนลูปตั้งแต่ i := 0 จนถึง i <= 4 (เนื่องจากอาร์เรย์มีขนาด 5 ดังนั้น index สุดท้ายคือ 4) และในแต่ละรอบของลูปเราใช้คำสั่ง fmt.Println(arr[i]) เพื่อแสดงค่าสมาชิกในอาร์เรย์ที่ตำแหน่ง i

เมื่อรันโปรแกรม ผลลัพธ์ที่ได้จะเป็นการแสดงข้อความ "Array elements:" ตามด้วยค่าสมาชิกในอาร์เรย์แต่ละตัวแยกบรรทัดกัน ซึ่งก็คือ "Ganga", "Yamuna", "Godavari", "Kaveri" และ "Narmada" ตามลำดับ
*/