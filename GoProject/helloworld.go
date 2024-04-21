package main

import "fmt"

/*
Go เราสามารถสร้าง slice ใหม่จาก slice ที่มีอยู่แล้วได้ โดยระบุ index ต่ำสุดและสูงสุดของ slice ต้นทางในรูปแบบเดียวกับการสร้าง slice จาก array
*/

func main() {
	a := [8]string{"Shanghai", "Hiroshima", "Jaipur", "Gorakhpur", "Nagpur", "Mumbai", "Nashik", "Lucknow"}
	ind_cities := a[2:]
	maha_cities := ind_cities[2:5]
	fmt.Println("Cities =", a)
	fmt.Println("Indian Cities =", ind_cities)
	fmt.Println("Maharashtra Cities =", maha_cities)
}

/* 
ในโปรแกรมนี้ เราสร้าง array a ที่มี 8 องค์ประกอบเป็นชื่อเมืองต่างๆ จากนั้นเราสร้าง slice ind_cities จาก array a โดยระบุ index ต่ำสุดเป็น 2 ซึ่งหมายความว่า slice ind_cities จะประกอบด้วยองค์ประกอบตั้งแต่ index 2 ถึงสุดท้ายของ array a

ต่อมา เราสร้าง slice maha_cities จาก slice ind_cities โดยระบุ index ต่ำสุดเป็น 2 และ index สูงสุดเป็น 5 ซึ่งหมายความว่า slice maha_cities จะประกอบด้วยองค์ประกอบตั้งแต่ index 2 ถึง 4 ของ slice ind_cities

เมื่อเราแสดงค่าของ array a, slice ind_cities และ slice maha_cities ผลลัพธ์ที่ได้คือ
"Cities = [Shanghai Hiroshima Jaipur Gorakhpur Nagpur Mumbai Nashik Lucknow]"
"Indian Cities = [Jaipur Gorakhpur Nagpur Mumbai Nashik Lucknow]"
"Maharashtra Cities = [Nagpur Mumbai Nashik]"

จากผลลัพธ์ จะเห็นได้ว่า slice ind_cities ประกอบด้วยองค์ประกอบตั้งแต่ index 2 ถึงสุดท้ายของ array a ซึ่งเป็นชื่อเมืองในประเทศอินเดีย และ slice maha_cities ประกอบด้วยองค์ประกอบตั้งแต่ index 2 ถึง 4 ของ slice ind_cities ซึ่งเป็นชื่อเมืองในรัฐมหาราษฏระของอินเดีย
*/