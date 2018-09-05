// struct  结构
// Go中的struct与C中的struct非常相似，并且Go没有class
// 使用type <Name> struct{}定义结构，名称遵循可见性规则
// 支持指向自身的指针类型成员
// 支持匿名结构，可用作成员或定义成员变量
// 匿名结构也可以用于map的值
// 可以使用字面值对结构进行初始化
// 允许直接通过指针来读写结构成员
// 相同类型的成员可进行直接拷贝赋值
// 支持 == 与 != 比较运算符，但不支持 > 或 <
// 支持匿名字段，本质上是定义了以某个类型名为名称的字段
// 嵌入结构作为匿名字段看起来像继承，但不是继承
// 可以使用匿名指针

package main

import (
	"fmt"
	)

type person struct {
	Name string
	Age int
	Contact struct{
		Phone, City string
	}
}

// 嵌入结构
type people struct {
	Name string
	Age int
}

type student struct {
	people
	Name string
}

type teacher struct {
	people
}

func main()  {
	// 初始化
	a := person{
		Name : "cocoyo",
		Age : 18,
	}
	fmt.Println(a) //{cocoyo 18}

	// 赋值
	d := person{}
	d.Name = "coocyo"
	d.Age = 18
	fmt.Println(d) //{cocoyo 18}

	ae := &person{
		Name : "cocoyo",
		Age : 18,
	}
	//可以这样赋值
	ae.Name = "Ning"
	doSomeThing(ae)
	fmt.Println(ae)

	// 匿名结构
	aa := struct {
		 Name string
		 Age int
	}{
		 Name : "cocoyo",
		 Age : 18,
	}
	fmt.Println(aa) // cocoyo 18}
	g := person{Name : "cocoyo", Age:18}
	// 初始化结构里面的匿名结构
	g.Contact.Phone = "21511213"
	g.Contact.City = "asdsd"
	fmt.Println(g) // { 0 { }}  {cocoyo 18 {21511213 asdsd}

	// 结构比较 需要相同类型才能比较
	b := person{Name : "Ning", Age : 20}
	c := person{Name : "Ning", Age : 21}
	fmt.Println(b == c) // false

	// 结构组合
	student := student{people : people{Name : "cocoyo", Age : 18}} // {{cocoyo 18}}
	teacher := teacher{people : people{Name: "Ning", Age : 20}} // {{Ning 20}}

	// 赋值
	student.people.Name = "jelly" // {{jelly 18}}
	fmt.Println(student, teacher)
}

func doSomeThing(person *person)  {
	person.Age = 20
	fmt.Println(person)
}