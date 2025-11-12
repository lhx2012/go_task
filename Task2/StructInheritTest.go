package Task2

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	person     Person
	employeeID uint64
}

func (e *Employee) Print() string {
	str := fmt.Sprintf("该员工信息如下：工号：%d,姓名：%s，年龄：%d", e.employeeID, e.person.name, e.person.age)
	return str
}

type PrintInfo interface {
	Print() string
}

func StructInheritRun() {
	emp := &Employee{employeeID: 123456789, person: Person{name: "张三", age: 18}}
	fmt.Println(emp.Print())
}
