package main

import "fmt"

/*
ในการประกาศและกำหนดค่าเริ่มต้นให้กับ struct type ในภาษา Go เราเริ่มต้นด้วยคีย์เวิร์ด type ตามด้วยชื่อของ struct ใหม่ และปิดท้ายด้วยคีย์เวิร์ด struct ภายในวงเล็บปีกกา เราระบุชื่อและประเภทของฟิลด์ข้อมูลเป็นชุดๆ เราสามารถกำหนดค่าเริ่มต้นให้กับตัวแปรของ struct type โดยใช้ struct literal และต้องส่งค่าให้กับฟิลด์ตามลำดับที่ประกาศไว้ใน struct
*/

// กำหนด struct type
type city_list struct {
	city string
	state string
	country string
}

func main() {
	// ประกาศและกำหนดค่าเริ่มต้นให้ struct โดยใช้ struct literal
	// c1 := city_list{"Mumbai", "Maharashtra", "India"}
	// c2 := city_list{"Chicago", "Illinois", "US"}
	// c3 := city_list{"Sydney", "New South Wales", "Australia"}

	// fmt.Println("City1:", c1)
	// fmt.Println("City2:", c2)
	// fmt.Println("City3:", c3)
	/* 
	1. เรากำหนด struct type ชื่อ `city_list` ที่มีฟิลด์ `city`, `state` และ `country` เป็น string
	2. ในฟังก์ชัน `main` เราประกาศตัวแปร `c1`, `c2` และ `c3` ของ struct type `city_list` และกำหนดค่าเริ่มต้นโดยใช้ struct literal
	3. เราส่งค่าให้กับฟิลด์ตามลำดับที่ประกาศไว้ใน struct คือ `city`, `state` และ `country`
	4. สุดท้ายเราแสดงค่าของตัวแปร struct ทั้งสามโดยใช้ `fmt.Println`
	
	*/


	// กำหนดค่าเริ่มต้นให้กับ struct โดยระบุชื่อฟิลด์ด้วย
	c1 := city_list{city: "Mumbai", state: "Maharashtra", country: "India"}
	c2 := city_list{city: "Chicago", state: "Illinois", country: "US"}
	c3 := city_list{city: "Sydney", state: "New South Wales", country: "Australia"}
	/* 
		เราระบุชื่อฟิลด์ในขณะที่กำหนดค่าเริ่มต้นให้กับ struct ด้วยการใช้รูปแบบ `field: value` ซึ่งช่วยให้โค้ดอ่านง่ายและชัดเจนมากขึ้น

		การประกาศและกำหนดค่าเริ่มต้นให้กับ struct type ในภาษา Go เป็นเรื่องง่ายและตรงไปตรงมา เราสามารถกำหนดค่าเริ่มต้นโดยใช้ struct literal และระบุค่าตามลำดับของฟิลด์หรือระบุชื่อฟิลด์ก็ได้ ซึ่งช่วยให้เราสามารถจัดการและจัดกลุ่มข้อมูลที่เกี่ยวข้องกันได้อย่างสะดวกและเป็นระเบียบ
	*/

	fmt.Println("City1:", c1)
	fmt.Println("City1:", c2)
	fmt.Println("City1:", c3)
}

/* 

*/