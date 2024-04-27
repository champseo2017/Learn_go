package main

import "fmt"

/*
ในภาษา Go เราสามารถวนลูปเพื่อเข้าถึงข้อมูลใน map ได้โดยใช้ range form ของ for loop ซึ่งจะให้คู่ key และ value ในแต่ละรอบของลูป
อย่างไรก็ตาม map เป็น unordered collection ดังนั้นลำดับการวนลูปใน map จึงไม่ได้รับการการันตีว่าจะเหมือนกันทุกครั้งที่วนลูป ดังนั้นถ้ารันโปรแกรมที่มีการวนลูป map หลายๆ ครั้ง ผลลัพธ์ที่ได้อาจมีลำดับที่แตกต่างกันออกไป
*/

func main() {
	// var my_map = map[string]string {
	// 	"Maharashtra":   "Mumbai",
    //     "Uttar_Pradesh": "Lucknow",
    //     "Rajasthan":     "Jaipur",
    //     "Karnataka":     "Bengaluru",
	// }
	// for state, capital := range my_map {
	// 	fmt.Println(state, capital)
	// }
	/* 
		- ประกาศตัวแปร `my_map` เป็น map ที่มี key และ value เป็น string โดยมีข้อมูลเป็นคู่ของรัฐ (state) และเมืองหลวง (capital)
		- ใช้ for loop ร่วมกับ range เพื่อวนลูปใน `my_map`
		- ในแต่ละรอบ ตัวแปร `state` จะเก็บ key (ชื่อรัฐ) และ `capital` จะเก็บ value (ชื่อเมืองหลวง)
		- แสดงค่า `state` และ `capital` ในแต่ละรอบของลูป

		ผลลัพธ์จะเป็นการแสดงคู่ของรัฐและเมืองหลวงทีละคู่ แต่ลำดับอาจไม่เหมือนกับที่กำหนดใน map เนื่องจาก map เป็น unordered collection
	*/
	var my_map2 = map[string]int {
		"India":     1947,
        "Singapore": 1965,
        "Australia": 1901,
        "Malaysia":  1957,
	}
	for country, year_of_independence := range my_map2 {
		fmt.Println(country, year_of_independence)
	}
	/* 
	- ประกาศตัวแปร `my_map` เป็น map ที่มี key เป็น string และ value เป็น int โดยมีข้อมูลเป็นคู่ของประเทศ (country) และปีที่ได้รับเอกราช (year of independence)
	- ใช้ for loop ร่วมกับ range เพื่อวนลูปใน `my_map`
	- ในแต่ละรอบ ตัวแปร `country` จะเก็บ key (ชื่อประเทศ) และ `year_of_independence` จะเก็บ value (ปีที่ได้รับเอกราช)
	- แสดงค่า `country` และ `year_of_independence` ในแต่ละรอบของลูป

	ผลลัพธ์จะเป็นการแสดงคู่ของประเทศและปีที่ได้รับเอกราชทีละคู่ แต่ลำดับอาจไม่เหมือนกับที่กำหนดใน map เนื่องจาก map เป็น unordered collection
	*/
}
