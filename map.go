// map
// 类似其它语言中的哈希表或字典，以key-value形式存储数据
// key必须是支持==或!=比较运算的类型，不可以是函数、map或slice
// map查找比线性搜索快很多，但比使用索引访问数据的类型慢100倍
// map使用make()创建，支持:=这种方式简写

// make([keyType]valueType, cap), cap表示容量，可省略
// 超出容量时会自动扩容，但尽量提供一个合理的初始值
// 使用len()获取元素的个数

// 键值对不存在时自动添加，使用delete()删除某个键值对
// 使用 for range 对map和slice进行迭代操作

package main

import "fmt"

func main()  {
	// 定义一个map
	//var m map[int]string = make(map[int]string)
	// or
	m := make(map[int]string)
	m[1] = "ok"

	fmt.Println(m) // map[1:ok]

	// 复杂的map
	m1 := make(map[int]map[int]string)
	m1[1] = make(map[int]string)
	m1[1][1] = "cocoyo"

	fmt.Println(m1) // map[1:map[1:cocoyo]]

	// 以map为元素类型的slice
	sm := make([]map[int]string, 5)

	for _,v := range sm{
		// 判断map是否初始化
		a,e := v[1]

		if !e {
			v = make(map[int]string)
		}
		v[1] = "cocoyo"
        a, e = v[1]

		fmt.Println(a, e)
	}
}