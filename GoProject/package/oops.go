package oops

type Int int // ประกาศชนิดข้อมูล Int

func (a *Int) Add_One() { // Pointer Receiver Function
	*a = *a + 1 // เปลี่ยนแปลงค่าของ a โดยตรงผ่าน Pointer
}

func (a Int) Double() Int { // Normal Receiver Function
	return (a * 2) // คืนค่า a * 2 กลับไป
}
