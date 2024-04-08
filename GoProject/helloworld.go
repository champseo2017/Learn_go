package main

/*
กล่าวถึงโครงสร้างควบคุมการทำงานในภาษา Go ซึ่งใช้ในการเปลี่ยนแปลงลำดับการทำงานของโปรแกรมและทำให้โปรแกรมสามารถเรียกใช้โค้ดตามเงื่อนไขที่กำหนด โดยจะครอบคลุมเรื่องการตัดสินใจในภาษา Go โดยใช้คำสั่ง if, if…else, if…else if…else, switch…case และ fallthrough นอกจากนี้ยังครอบคลุมแนวคิดเกี่ยวกับการวนลูป ได้แก่ for loop, nested for loop และคำสั่งควบคุมลูป เช่น break, goto, continue เป็นต้น

คำสั่ง if และ if…else:
if condition {
    // code to execute if condition is true
} else {
    // code to execute if condition is false
}
ถ้าเงื่อนไข condition เป็นจริง โค้ดภายในบล็อก if จะถูกเรียกใช้งาน
ถ้าเงื่อนไข condition เป็นเท็จ โค้ดภายในบล็อก else จะถูกเรียกใช้งาน

คำสั่ง if…else if…else:
if condition1 {
    // code to execute if condition1 is true
} else if condition2 {
    // code to execute if condition2 is true
} else {
    // code to execute if all conditions are false
}
ถ้าเงื่อนไข condition1 เป็นจริง โค้ดภายในบล็อก if จะถูกเรียกใช้งาน
ถ้าเงื่อนไข condition1 เป็นเท็จ แต่เงื่อนไข condition2 เป็นจริง โค้ดภายในบล็อก else if จะถูกเรียกใช้งาน
ถ้าทุกเงื่อนไขเป็นเท็จ โค้ดภายในบล็อก else จะถูกเรียกใช้งาน

คำสั่ง switch…case:
switch expression {
case value1:
    // code to execute if expression matches value1
case value2:
    // code to execute if expression matches value2
default:
    // code to execute if no case matches
}
expression จะถูกประเมินและเปรียบเทียบกับค่าในแต่ละ case
ถ้าค่าของ expression ตรงกับค่าในกรณีใด โค้ดภายในกรณีนั้นจะถูกเรียกใช้งาน
ถ้าไม่มีกรณีใดตรงกับค่าของ expression โค้ดภายในบล็อก default จะถูกเรียกใช้งาน

คำสั่ง for loop:
for initialization; condition; update {
    // code to execute in each iteration
}
initialization คือการกำหนดค่าเริ่มต้นของตัวแปรที่ใช้ในลูป
condition คือเงื่อนไขในการทำงานของลูป ลูปจะทำงานต่อไปจนกว่าเงื่อนไขจะเป็นเท็จ
update คือการอัปเดตค่าของตัวแปรที่ใช้ในลูปหลังจากแต่ละรอบ

คำสั่งควบคุมลูป (break, continue):
for condition {
    if anotherCondition {
        break
    }
    if yetAnotherCondition {
        continue
    }
    // code to execute in each iteration
}
break ใช้เพื่อหยุดการทำงานของลูปทันที
continue ใช้เพื่อข้ามการทำงานในรอบปัจจุบันและไปเริ่มต้นรอบถัดไปของลูป
*/

func main() {
	
}