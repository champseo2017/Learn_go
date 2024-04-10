package main

import "fmt"

/*
อธิบายเกี่ยวกับคำสั่ง Fallthrough ในภาษา Go ซึ่งใช้ในคำสั่ง switch เพื่อส่งการควบคุมไปยังคำสั่งแรกของ case ถัดไป โดยปกติเมื่อมีการตรงกับเงื่อนไขใน case ใดๆ โปรแกรมจะออกจาก switch ทันที แต่ด้วยคำสั่ง Fallthrough จะทำให้โปรแกรมทำงานใน case ถัดไปต่อ แม้ว่าจะไม่ตรงกับเงื่อนไข อย่างไรก็ตาม Fallthrough ไม่สามารถใช้ใน case สุดท้ายได้ เพราะไม่มี case ถัดไปแล้ว หากใช้จะเกิด compilation error
*/

func main() {
	switch number := 200; {
	case number < 300:
		fmt.Printf("%d is less than 300\n", number)
		fallthrough
	case number > 100:
		fmt.Printf("%d is greater than 100", number)
	}
}

/*
1. `package main` และ `import "fmt"` เป็นการประกาศแพ็คเกจและนำเข้าแพ็คเกจ fmt สำหรับใช้งานฟังก์ชันพิมพ์ข้อความ
2. `func main()` เป็นฟังก์ชันหลักที่จะถูกเรียกเมื่อโปรแกรมเริ่มทำงาน
3. `switch number := 200; { ... }` เป็นการประกาศตัวแปร number และกำหนดค่าเป็น 200 ภายในคำสั่ง switch
4. `case number < 300:` เป็นเงื่อนไขแรก ถ้า number น้อยกว่า 300 จะทำคำสั่งใน case นี้
5. `fmt.Printf("%d is less than 300\n", number)` พิมพ์ข้อความ "200 is less than 300"
6. `fallthrough` ทำให้โปรแกรมทำงานใน case ถัดไปต่อ แม้ว่าเงื่อนไขจะไม่ตรงกับ case ถัดไป
7. `case number > 100:` เป็นเงื่อนไขที่สอง ถ้า number มากกว่า 100 จะทำคำสั่งใน case นี้
8. `fmt.Printf("%d is greater than 100", number)` พิมพ์ข้อความ "200 is greater than 100"

ผลลัพธ์ที่ได้จากโปรแกรมนี้ คือ:
200 is less than 300
200 is greater than 100

โดยปกติโปรแกรมจะพิมพ์แค่ข้อความแรก แต่เนื่องจากมีคำสั่ง Fallthrough จึงทำให้โปรแกรมพิมพ์ข้อความที่สองต่อ แม้ว่าเงื่อนไขที่สองจะไม่เป็นจริงก็ตาม

ส่วนโปรแกรมที่สองแสดงให้เห็นว่าเราไม่สามารถใช้ Fallthrough ใน case สุดท้ายได้ เพราะจะเกิด compilation error
*/