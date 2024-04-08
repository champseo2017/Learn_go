package main

import "fmt"

/*
ตัวอย่างการใช้คำสั่ง if…else:
if age >= 18 {
    fmt.Println("You are eligible to vote")
} else {
    fmt.Println("You are not eligible to vote")
}

ตัวอย่างการใช้คำสั่ง if…else if…else:
if grade >= 80 {
    fmt.Println("Grade: A")
} else if grade >= 70 {
    fmt.Println("Grade: B")
} else if grade >= 60 {
    fmt.Println("Grade: C")
} else {
    fmt.Println("Grade: F")
}

ตัวอย่างการใช้คำสั่ง switch…case:
switch dayNumber {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
case 4:
    fmt.Println("Thursday")
case 5:
    fmt.Println("Friday")
case 6:
    fmt.Println("Saturday")
case 7:
    fmt.Println("Sunday")
default:
    fmt.Println("Invalid day number")
}

ตัวอย่างการใช้ for loop:
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}

ตัวอย่างการใช้ nested for loop:
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("(%d, %d) ", i, j)
    }
    fmt.Println()
}

ตัวอย่างการใช้ break และ continue:
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue
    }
    if i == 4 {
        break
    }
    fmt.Println(i)
}
*/

func main() {
	// nested for loop:
	// for i := 1; i <= 3; i++ {
	// 	for j := 1; j <= 3; j++ {
	// 		fmt.Printf("(%d, %d) ", i, j)
	// 	}
	// 	fmt.Println()
	// }
	// ตัวอย่างการใช้คำสั่ง fallthrough:
	switch num := 2; num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
		// คำสั่ง fallthrough มักใช้ในกรณีที่ต้องการให้โปรแกรมเรียกใช้โค้ดในกรณีถัดไปโดยอัตโนมัติ
	case 3:
		fmt.Println("Three")
	}
	/* 
	กำหนด switch แบบปกติ
	num := 2
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}
	
	*/
}