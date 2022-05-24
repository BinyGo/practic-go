package main

import "fmt"

type User struct {
	Name string
}

func (u *User) SetName(name string) {
	u.Name = name
	fmt.Println(u.Name)
}

type Employee User

func main() {
	employee := new(Employee)
	fmt.Println(employee)
	//employee.SetName("biny") //employee.SetName undefined (type *Employee has no field or method SetName)
	//参考答案及解析：编译不通过。当使用 type 声明一个新类型，它不会继承原有类型的方法集。
	employee2 := new(User)
	employee2.SetName("biny")
}
