// 反射reflection

// 反射可大大提高程序的灵活性，使得interface{}有更大的发挥余地
// 反射使用TypeOf或ValueOf 函数从接口中获取目标对象信息
// 反射会将匿名字段作为独立字段（匿名字段本质）
// 想要利用反射修改对象状态,前提是interface.data是settable，即pointer-interface
// 通过反射可以"动态"调用方法

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int
	Id int
}

func main()  {
	user := User{"cocoyo", 20, 1}
	user.Hello("ning")
	Info(user)

	// 利用反射修改基本类型状态
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)

	// 利用反射修改结构
	Set(&user)
	fmt.Println(user)

	// 利用反射方法的调用
	reflectTransferMethods(user)
}

func (user User) Hello(name string) {
	fmt.Println(user.Name + " say:hello " + name)
}

func Info(o interface{})  {
	t := reflect.TypeOf(o)
	// t.Name 获取结构的名称
	fmt.Println(t.Name())

	v := reflect.ValueOf(o)

	for i := 0; i < t.NumField(); i++{
		f := t.Field(i)
		v := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\r\n", f.Name, f.Type, v)
	}
}

func Set(o interface{})  {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("不能修改")
	} else {
		v = v.Elem()
	}

	if f := v.FieldByName("Id"); f.Kind() == reflect.Int {
		f.SetInt(22)
	}
}

func reflectTransferMethods(user interface{})  {
	v := reflect.ValueOf(user)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("不能修改")

		return
	}

	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("jelly")}

	mv.Call(args)
}