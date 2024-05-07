package main

import "fmt"

/*
Empty interface คือ interface ที่ไม่มี method ใดๆ เลย ทำให้ทุก type สามารถ satisfy interface นี้ได้ เราจึงสามารถกำหนดค่าของ type ใดๆ ให้กับตัวแปรของ empty interface ได้

ประโยชน์ของ empty interface คือ เราสามารถใช้มันเป็น parameter ของฟังก์ชันที่ยอมรับ argument หลายๆ type ได้ เช่นในฟังก์ชัน fmt.Println() ที่รับ var-args ของ empty interface ทำให้เราสามารถส่งค่าของ type ใดๆ เข้าไปได้

ตัวอย่าง code:

```go
type MyEmpty interface {}

func main() {
    var empty MyEmpty

    empty = 10
    empty = "a"
    empty = 22.3
    empty = Student{101, "Shyam"}
}
```

ในตัวอย่าง เรากำหนดค่าของ type ต่างๆ เช่น int, string, float32 และ custom type Student ให้กับตัวแปร empty ของ MyEmpty interface ได้ เพราะทุก type สามารถ satisfy empty interface
*/

type MyEmpty interface {

}

type Student struct {
    id int
    name string
}

func main() {
    empty := MyEmpty(10)
    fmt.Println(empty)

    empty = "a"
    fmt.Println(empty)

    empty = 22.3
    fmt.Println(empty)

    empty = Student{101, "Shyam"}
    fmt.Println(empty)
}
