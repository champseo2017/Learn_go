package main

import "fmt"

/*
Basic Types และการสร้างเมธอดบนประเภทที่กำหนดเองในภาษา Go:

1. ในภาษา Go เราไม่สามารถสร้างเมธอดบน basic types (เช่น int, float64, string) ได้โดยตรง

2. หากต้องการสร้างเมธอดบน basic type เราจำเป็นต้องสร้างประเภทที่กำหนดเองโดยใช้ basic type เป็นพื้นฐาน เช่น:
   ```go
   type MyInt int
   ```

3. หลังจากสร้างประเภทที่กำหนดเองแล้ว เราสามารถสร้างเมธอดบนประเภทนั้นได้ เช่น:
   ```go
   func (m MyInt) IsEven() bool {
       return m%2 == 0
   }
   ```

4. เมธอดสามารถมี value receiver หรือ pointer receiver ก็ได้ ขึ้นอยู่กับว่าเราต้องการแก้ไขค่าของ receiver ภายในเมธอดหรือไม่

5. การสร้างประเภทที่กำหนดเองจาก basic type ช่วยให้เราสามารถเพิ่มเมธอดและพฤติกรรมเฉพาะให้กับ basic type ได้ ทำให้โค้ดมีความชัดเจนและง่ายต่อการใช้งานมากขึ้น

ดังนั้น หากเราต้องการเพิ่มฟังก์ชันการทำงานพิเศษให้กับ basic type ในภาษา Go เราสามารถทำได้โดยการสร้างประเภทที่กำหนดเองและสร้างเมธอดบนประเภทนั้น

การใช้ pointer receiver และ value receiver ในเมธอดของภาษา Go:

1. ใช้ pointer receiver (`*Type`) เมื่อต้องการแก้ไขค่าของ receiver ภายในเมธอด
   - ในเมธอด สามารถแก้ไขค่าของ receiver ได้โดยตรง โดยไม่จำเป็นต้องใช้ `*` เพื่อเข้าถึงค่าที่ pointer ชี้ไปหา
   - ไม่จำเป็นต้อง return ค่าจากเมธอด เพราะการแก้ไขค่าผ่าน pointer จะส่งผลต่อตัวแปรต้นฉบับ

2. ใช้ value receiver (`Type`) เมื่อไม่จำเป็นต้องแก้ไขค่าของ receiver ภายในเมธอด
   - Go จะส่งค่าสำเนา (copy) ของ receiver ไปยังเมธอด
   - การแก้ไขค่าภายในเมธอดจะไม่ส่งผลกระทบต่อค่าของ receiver ดั้งเดิม

การเลือกใช้ pointer receiver หรือ value receiver ขึ้นอยู่กับว่าเมธอดนั้นจำเป็นต้องแก้ไขค่าของ receiver หรือไม่ หากต้องการแก้ไข ให้ใช้ pointer receiver แต่ถ้าไม่จำเป็น ใช้ value receiver ก็เพียงพอ
*/

type number int

func (n number) isPrime() bool {
	for i := 2; i < int(n); i++ {
		if int(n)%i == 0 {
			return false
		}
	}
	return true
}

func (n number) isDivisible(i number) bool {
	return (n % i) == 0
}

func (n *number) increamentBy(i number) {
	*n = *n + i
}

func main() {
	var n number = 7
	fmt.Println(n.isPrime())
	fmt.Println(n.isDivisible(3))
	n.increamentBy(3)
	fmt.Println(n)
}
/* 

*/