package main

import "fmt"

/*
1. การตรวจสอบว่าอาร์เรย์ว่างหรือไม่ในภาษา Go สามารถทำได้โดยการตรวจสอบความยาวของอาร์เรย์ด้วยฟังก์ชัน `len()`
2. การหาความยาวของอาร์เรย์ในภาษา Go สามารถทำได้โดยใช้ฟังก์ชัน `len()`
3. การกำหนดและเข้าถึงค่าที่ดัชนีของอาร์เรย์ในภาษา Go สามารถทำได้โดยใช้วิธี `array[index] = value` และ `array[index]` ตามลำดับ
4. อาร์เรย์ในภาษา Go มีบทบาทสำคัญในการเก็บข้อมูลที่มีชนิดเดียวกันและมีขนาดคงที่ โดยสามารถเข้าถึงข้อมูลได้อย่างรวดเร็วผ่านดัชนี
5. การวนซ้ำอาร์เรย์ในภาษา Go สามารถทำได้โดยใช้ for loop หรือ range operator
6. การส่งอาร์เรย์เข้าไปในฟังก์ชันในภาษา Go จะเป็นการส่งแบบ pass by value ซึ่งฟังก์ชันจะได้รับสำเนาของอาร์เรย์ การเปลี่ยนแปลงอาร์เรย์ในฟังก์ชันจะไม่ส่งผลต่ออาร์เรย์ต้นฉบับ
*/

// ตรวจสอบว่าอาร์เรย์ว่างหรือไม่
func isArrayEmpty(arr []int) bool {
	return len(arr) == 0
}
// หาความยาวของอาร์เรย์
func getArrayLength(arr []int) int {
	return len(arr)
}
// กำหนดและเข้าถึงค่าที่ดัชนีของอาร์เรย์
func setAndGetValue(arr []int, index int, value int) int {
	arr[index] = value
	return arr[index]
}
// วนซ้ำอาร์เรย์ด้วย for loop
func iterateArrayWithForLoop(arr []int) {
	for i:= 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
// วนซ้ำอาร์เรย์ด้วย range operator
func iterateArrayWithRange(arr []int) {
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
// ส่งอาร์เรย์เข้าไปในฟังก์ชัน
func modifyArray(arr []int) {
	arr[0] = 100
}
func main() {
	arr := []int{1, 2, 3, 4, 5}

	// ตรวจสอบว่าอาร์เรย์ว่างหรือไม่
	isEmpty := isArrayEmpty(arr)
	fmt.Println("Is array empty?", isEmpty)

	// หาความยาวของอาร์เรย์
	length := getArrayLength(arr)
	fmt.Println("Array length:", length)

	// กำหนดและเข้าถึงค่าที่ดัชนีของอาร์เรย์
	value := setAndGetValue(arr, 2, 10)
	fmt.Println("Value at index 2:", value)

	// วนซ้ำอาร์เรย์ด้วย for loop
	fmt.Println("Iterating array with for loop:")
	iterateArrayWithForLoop(arr)

	// วนซ้ำอาร์เรย์ด้วย range operator
	fmt.Println("Iterating array with range:")
	iterateArrayWithRange(arr)

	// ส่งอาร์เรย์เข้าไปในฟังก์ชัน
	fmt.Println("Before modifying:", arr)
	modifyArray(arr)
	fmt.Println("After modifying:", arr)
	/*
		ค่าของอาร์เรย์ arr ถูกเปลี่ยนแปลงหลังจากเรียกใช้ฟังก์ชัน modifyArray เนื่องจากการส่งอาร์เรย์เข้าไปในฟังก์ชันเป็นการส่งแบบ pass by reference

		ดังนั้น ในภาษา Go การส่งอาร์เรย์เข้าไปในฟังก์ชันจึงเป็นการส่งแบบ pass by reference ไม่ใช่ pass by value ซึ่งการเปลี่ยนแปลงค่าในอาร์เรย์ภายในฟังก์ชันจะส่งผลต่ออาร์เรย์ต้นฉบับด้วย
	*/
}

