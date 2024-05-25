package main

import (
	"fmt"
	"sync"
)

/*
หาข้อผิดพลาดในโปรแกรม
*/

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		fmt.Println("Running go routine")
		wg.Done()
	}()

	wg.Wait()
}

/*
ข้อผิดพลาดคือ มีการเรียกใช้ wg.Add(2) แต่สร้าง goroutine เพียงตัวเดียว ทำให้โปรแกรมไม่จบการทำงาน
*/
