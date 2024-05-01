package main

import "fmt"

/*
สรุปแบบสั้นๆ Map ในภาษา Go:

Map ในภาษา Go คือคอลเลกชันของคู่ key-value โดยที่ key ต้องไม่ซ้ำกัน

วิธีสร้าง map มี 2 แบบ คือ
1. map literal เช่น m := map[string]int{"a": 1, "b": 2}
2. ใช้ฟังก์ชัน make() เช่น m := make(map[string]int)

การเพิ่ม/แก้ไข value ทำได้โดย m[key] = value
การดึง value ทำได้โดย value = m[key]
การลบ ทำได้โดยใช้ delete(m, key)

การวนลูปเพื่ออ่านค่าใน map ทำได้โดยใช้ for range loop

map เป็น reference type เมื่อกำหนด m2 := m1 ทั้งสองจะชี้ไปยังโครงสร้างข้อมูลเดียวกัน การเปลี่ยนแปลงที่ m2 จะส่งผลต่อ m1 ด้วย

โดยสรุปแล้ว map ใช้เก็บคู่ key-value มีความสามารถในการเพิ่ม แก้ไข ลบ และค้นหา value จาก key ได้อย่างรวดเร็ว
*/

func main() {
	// สร้าง map ด้วย map literal
	myMap := map[string]int {
		"apple": 5,
		"banana": 3,
		"orange": 7,
	}
	fmt.Println("myMap:", myMap)
	// สร้าง map ด้วย make()
	myNumbers := make(map[int]float64)
	myNumbers[1] = 3.14
	myNumbers[2] = 1.618

	fmt.Println("myNumbers:", myNumbers) // แสดงผล: myNumbers: map[1:3.14 2:1.618]
	// วนลูปไปยัง map 
	for key, value := range myMap {
		fmt.Printf("%s -> %d\n", key, value)
	}
	 /*
		ผลลัพธ์ที่ได้อาจแตกต่างกันในแต่ละครั้งที่รันโปรแกรม เนื่องจากเป็น unordered collection เช่น
		apple -> 5 
		banana -> 3
		orange -> 7
    */
	// การเพิ่ม key ที่มีอยู่แล้วจะเปลี่ยน value
	myMap["banana"] = 10
	fmt.Println("myMap after update:", myMap) // แสดงผล: myMap after update: map[apple:5 banana:10 orange:7]
	// การเพิ่ม key ใหม่
	myMap["grape"] = 15
	fmt.Println("myMap after adding:", myMap) // แสดงผล: myMap after adding: map[apple:5 banana:10 grape:15 orange:7]
	// ตรวจสอบว่า key มีอยู่หรือไม่
	value, ok := myMap["kiwi"]
	if !ok {
		fmt.Println("Key 'kiwi' does not exist.") // แสดงผล: Key 'kiwi' does not exist.
	} else {
		fmt.Println("Value of 'kiwi':", value)
	}
	// ลบ key ออกจาก map
	delete(myMap, "orange")
	fmt.Println("myMap after deletion:", myMap) // แสดงผล: myMap after deletion: map[apple:5 banana:10 grape:15]
	// ตัวอย่างการใช้ map เป็น reference type
	newMap := myMap
	newMap["pear"] = 20
	fmt.Println("newMap:", newMap)        // แสดงผล: newMap: map[apple:5 banana:10 grape:15 pear:20]
	fmt.Println("myMap after changing newMap:", myMap) // แสดงผล: myMap after changing newMap: map[apple:5 banana:10 grape:15 pear:20]
}

/* 

*/