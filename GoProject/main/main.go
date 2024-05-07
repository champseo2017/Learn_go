package main

import "fmt"

/*
ในโปรแกรม 9.4 มีการกำหนดอินเทอร์เฟซ `Connection` ที่มีสองเมธอด `Open()` และ `Close()` และมีโครงสร้าง `DbConnection` ที่มีเพียงเมธอด `Close()` เท่านั้น

เมื่อพยายามกำหนดค่า `DbConnection{}` ให้กับตัวแปร `conn` ชนิด `Connection` คอมไพเลอร์ของ Go จะตรวจสอบว่าโครงสร้าง `DbConnection` มีเมธอดครบตามที่ระบุในอินเทอร์เฟซ `Connection` หรือไม่

เนื่องจากโครงสร้าง `DbConnection` ขาดเมธอด `Open()` คอมไพเลอร์จึงรายงานข้อผิดพลาด `Error 9.2` ว่าไม่สามารถกำหนดค่า `DbConnection{}` ให้กับตัวแปร `conn` ได้

คอมไพเลอร์ของ Go มีกลไกตรวจสอบความถูกต้องของการใช้อินเทอร์เฟซ ซึ่งช่วยป้องกันข้อผิดพลาดและทำให้โค้ดมีความน่าเชื่อถือ
*/

type Session struct {}

type Connection interface {
    Open(uri string) (Session, error)
    Close() error
}

type DbConnection struct {}

func (conn DbConnection) Close() error {
    fmt.Println("Closing Database Connection")
    return nil
}

func (conn DbConnection) Open(uri string) (Session, error) {
    fmt.Println("Opening Database Connection")
    return Session{}, nil
}

func main() {
    conn := Connection(DbConnection{})
    /* 
    โปรแกรมนี้จะเกิดข้อผิดพลาด Error 9.2 เนื่องจากโครงสร้าง DbConnection ไม่ได้กำหนดเมธอด Open() ตามที่ระบุในอินเทอร์เฟซ Connection ทำให้ไม่สามารถกำหนดค่า DbConnection{} ให้กับตัวแปร conn ได้
    */
    conn.Close()
    /* 
    เมื่อเพิ่มเมธอด Open() แล้ว โปรแกรมจะสามารถคอมไพล์และทำงานได้อย่างถูกต้อง โดยเมื่อเรียกใช้เมธอด Close() ผ่านตัวแปร conn จะแสดงข้อความ "Closing Database Connection" ออกทางหน้าจอ
    */
}