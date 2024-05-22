package module

// สตรักเจอร์ Person ประกอบด้วยชื่อ อายุ และที่อยู่
type Person struct {
	Name    string
	Age     int
	Address Address
}

// สตรักเจอร์ Address ประกอบด้วยบ้านเลขที่และเมือง
type Address struct {
	HouseNo string
	City    string
}
