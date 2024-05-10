package main

import (
	"fmt"
	"strings"
)

/*

 */

func main() {
	str := "Hello how are u. u will go to school" // กำหนดสตริง "Hello how are u. u will go to school" ให้กับตัวแปร str
	substr := "how"                               // กำหนดสตริง "how" ให้กับตัวแปร substr

	fmt.Println(strings.Contains(str, substr)) // ตรวจสอบว่าสตริง str มีสตริงย่อย substr อยู่ภายในหรือไม่ และแสดงผลลัพธ์เป็น true หรือ false

	rep := strings.ReplaceAll(str, "u", "Udit") // แทนที่ทุกการปรากฏของ "u" ในสตริง str ด้วย "Udit" และกำหนดผลลัพธ์ให้กับตัวแปร rep

	arr := strings.Split(str, " ") // แยกสตริง str ออกเป็นสไลซ์ของสตริงย่อยโดยใช้ช่องว่าง (" ") เป็นตัวคั่น และกำหนดผลลัพธ์ให้กับตัวแปร arr

	fmt.Println(rep) // แสดงค่าของตัวแปร rep ซึ่งเป็นสตริงที่ถูกแทนที่ "u" ด้วย "Udit"

	fmt.Println(arr) // แสดงค่าของตัวแปร arr ซึ่งเป็นสไลซ์ของสตริงย่อยที่ถูกแยกออกจากสตริง str

	for i := 0; i < len(arr); i++ { // วนลูปตามความยาวของสไลซ์ arr
		fmt.Println(arr[i]) // แสดงสตริงย่อยแต่ละตัวในสไลซ์ arr
	}

}

/*

 */
