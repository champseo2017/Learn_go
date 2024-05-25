package main

/*

 */

func readVal(ch <-chan int) {
	// ฟังก์ชัน readVal รับพารามิเตอร์ ch ซึ่งเป็น channel แบบอ่านอย่างเดียว
	ch <- 10 // พยายามส่งค่า 10 เข้าไปใน channel ch ซึ่งจะทำให้เกิด compilation error
}

func writeVal(ch chan<- int) {
	// ฟังก์ชัน writeVal รับพารามิเตอร์ ch ซึ่งเป็น channel แบบเขียนอย่างเดียว
	<-ch // พยายามอ่านค่าจาก channel ch ซึ่งจะทำให้เกิด compilation error
}
func main() {

}

/*

 */
