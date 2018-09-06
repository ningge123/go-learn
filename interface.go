// interface
// Go中虽然没有class，但依旧有methods
// 只能为同一个包中的类型定义方法
// Receiver 可以是类型的值或者指针
// 不存在方法重载
// 可以使用值或指针类型来调用方法，编译器会自动完成转换
// 冲某种意义上说，方法是函数的语法糖，因为receiver其实就是方法所接受的第一个参数(Method Value vs. Method Expression)
// 如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
// 类型别名不会拥有底层类型所附带的方法
// 方法可以调用结构中的非公开字段

package main

import "fmt"

type cocoyo struct {
	Name string
}

func main()  {
	a := cocoyo{}
	a.Name = "cocoyo"
	panic(55)
	a.Print(a.Name)
}

func (cocoyo cocoyo) Print(a string) {
	fmt.Print(a)
}