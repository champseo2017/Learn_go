package main

/*
เขียนโปรแกรมเพื่อกำหนดประเภทที่กำหนดเองต่อไปนี้ของ int: Distance, Millimeter, Centimeter, Meter, Kilometer คุณต้องเพิ่มเมธอดในแต่ละประเภทเพื่อแปลงและส่งคืนค่าของประเภทอื่นๆ ตัวอย่างเช่น คุณสามารถกำหนดเมธอดใน Millimeter เพื่อแปลงค่าของมันเป็น Centimeter, Meter และ Kilometer
*/
// ตัวอย่างการกำหนดประเภทและเมธอดสำหรับ Millimeter
type Distance int

type Millimeter Distance
type Centimeter Distance
type Meter Distance
type Kilometer Distance

func (mm Millimeter) ToCentimeter() Centimeter {
    return Centimeter(mm / 10)
}

func (mm Millimeter) ToMeter() Meter {
    return Meter(mm / 1000)
}

func (mm Millimeter) ToKilometer() Kilometer {
    return Kilometer(mm / 1000000)
}

func (cm Centimeter) ToMillimeter() Millimeter {
    return Millimeter(cm * 10)
}

func (cm Centimeter) ToMeter() Meter {
    return Meter(cm / 100)
}

func (cm Centimeter) ToKilometer() Kilometer {
    return Kilometer(cm / 100000)
}

func (m Meter) ToMillimeter() Millimeter {
    return Millimeter(m * 1000)
}

func (m Meter) ToCentimeter() Centimeter {
    return Centimeter(m * 100)
}

func (m Meter) ToKilometer() Kilometer {
    return Kilometer(m / 1000)
}

func (km Kilometer) ToMillimeter() Millimeter {
    return Millimeter(km * 1000000)
}

func (km Kilometer) ToCentimeter() Centimeter {
    return Centimeter(km * 100000)
}

func (km Kilometer) ToMeter() Meter {
    return Meter(km * 1000)
}


func main() {
	
}
/* 

*/