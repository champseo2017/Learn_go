package main

/*
การประกาศ Slice
การประกาศ slice จะคล้ายกับการประกาศ array แต่จะไม่ระบุขนาดของ slice ดังนั้น slice จึงมีความยืดหยุ่นและมีขนาดแบบไดนามิกเพื่อให้เหมาะกับความต้องการ
*/

func main() {
	var my_slice []int
	// สร้าง slice ชื่อ team ที่มีความยาวเป็น 0
	var team []int
	// ประกาศ my_slice1 ของ integers โดยไม่ระบุขนาด
	var my_slice1 []int
	// ประกาศ slice ของ strings โดยไม่ระบุขนาด
	var my_slice_string []string
	// ประกาศ slice ของ floats โดยไม่ระบุขนาด
	var my_slice_float []float64
	// ประกาศ slice แบบสั้นโดยใช้ := 
	my_slice := []int{1, 2, 3}
	// ประกาศ slice แบบระบุขนาดเริ่มต้น
	my_slice := make([]int, 5)
}

/* 
- ประกาศ slice ของ integers ชื่อ `my_slice` โดยไม่ระบุขนาด ทำให้มีความยืดหยุ่นและสามารถเปลี่ยนแปลงขนาดได้
- ประกาศ slice ของ strings ชื่อ `my_slice_string` โดยไม่ระบุขนาด
- ประกาศ slice ของ floats ชื่อ `my_slice_float` โดยไม่ระบุขนาด
- ประกาศ slice แบบสั้นโดยใช้ `:=` และระบุค่าเริ่มต้น `{1, 2, 3}`
- ประกาศ slice แบบระบุขนาดเริ่มต้นโดยใช้ฟังก์ชัน `make` และระบุขนาดเป็น 5

การประกาศ slice โดยไม่ระบุขนาดทำให้เราสามารถเปลี่ยนแปลงขนาดของ slice ได้ในภายหลังโดยใช้ฟังก์ชัน `append` หรือ `copy` ซึ่งทำให้ slice มีความยืดหยุ่นและสามารถปรับเปลี่ยนได้ตามความต้องการในขณะรันโปรแกรม
*/