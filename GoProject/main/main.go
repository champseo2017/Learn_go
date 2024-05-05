package main

import "fmt"

/*

Composite Types และการสร้างเมธอดบน Slice ในภาษา Go:

1. Go อนุญาตให้สร้างเมธอดบน Composite Types เช่น Slice, Array, Map, Struct เป็นต้น

2. ในการสร้างเมธอดบน Slice เราต้องสร้างประเภทที่กำหนดเองจาก Slice ก่อน เช่น:
   ```go
   type MySlice []string
   ```

3. จากนั้นเราสามารถสร้างเมธอดบนประเภท `MySlice` ได้ เช่น:
   ```go
   func (s MySlice) Len() int {
       return len(s)
   }
   ```

4. ตัวอย่างเมธอดที่มีประโยชน์บน Slice ได้แก่:
   - `intersect` : หาส่วนตัดของ Slice ปัจจุบันกับ Slice ที่ส่งเข้ามาเป็นพารามิเตอร์
   - `remove` : ลบสมาชิกออกจาก Slice (ใช้ pointer receiver เพื่อแก้ไขค่าของ Slice)
   - `indexOf` : หาดัชนีของสมาชิกใน Slice

5. การสร้างเมธอดบน Composite Types ช่วยให้เราสามารถเพิ่มฟังก์ชันการทำงานเฉพาะให้กับประเภทข้อมูลเหล่านั้นได้ ทำให้โค้ดมีความชัดเจนและง่ายต่อการใช้งานมากขึ้น

ดังนั้น Go ให้ความยืดหยุ่นในการสร้างเมธอดบน Composite Types ซึ่งรวมถึง Slice ด้วย เราสามารถสร้างประเภทที่กำหนดเองและเพิ่มเมธอดที่มีประโยชน์บนประเภทนั้นได้ ช่วยให้การทำงานกับ Slice และประเภทข้อมูลอื่นๆ เป็นไปได้ง่ายและมีประสิทธิภาพมากขึ้น

*/
type list []string

func (l list) intersect(input list) list {
    mp := make(map[string]bool)
    for _, ele := range l {
        mp[ele] = true
    }
    var result list
    for _, ele := range input {
        if mp[ele] {
            result = append(result, ele)
        }
    }
    return result
}

func (l *list) remove(in string) {
    index := l.indexOf(in)
    *l = append((*l)[:index], (*l)[index+1:]...)
}

func (l list) indexOf(in string) int {
    for index, ele := range l {
        if ele == in {
            return index
        }
    }
    return -1
}

func main() {
	var l list = []string{"A", "B", "C"}
    var input list = []string{"B", "C", "D"}
    intersResult := l.intersect(input)
    fmt.Println("Intersection:", intersResult)
    l.remove("B")
    fmt.Println("After deletion:", l)

	updateIntersResult := l.intersect(input)
	fmt.Println("Updated Intersection:", updateIntersResult)

    fmt.Println("Index of B:", l.indexOf("B"))
    fmt.Println("Index of C:", l.indexOf("C"))
}
/* 

*/