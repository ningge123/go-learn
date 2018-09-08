// 并发 concurrency
// 并发不是并行
// 并发主要是由切换时间片来实现"用时"运行"，在并行则是直接利用多核实现多线程的运行，但Go可以设置用核数，以发挥多核计算机的能力
// Goroutine 奉行通过通信来共享内存，而不是共享内存来通信

// Channel

// Channel 是 goroutine 沟通的桥梁，大都是阻塞同步的
// 通过 make 创建，close关闭
// Channel是引用类型
// 可以使用for range来迭代不断操作Channel
// 可以设置单向或双向通道
// 可以设置缓存大小，在未被填满前不会发生阻塞

// Select

// 可以处理一个或多个channel的发送与接收
// 同时有多个可用的channel时按随机顺序处理
// 可用空的select来阻塞main函数
// 可设置超时
package main

import (
	"fmt"
	"sync"
)

func main()  {
	//c := make(chan bool) // 双向通道 可存 可取
	//go func() {
	//	fmt.Println("go go go")
	//	c <- true
	//	close(c)
	//}()
	//
	//// 读取channel
	//for v := range c {
	//	fmt.Println(v)
	//}
	s := sync.WaitGroup{}
	s.Add(10)

	for i := 0; i < 10; i++ {
		go Go(&s, i)
	}

	s.Wait()
}

func Go(s *sync.WaitGroup, num int)  {
	v := 1
	for i := 0; i < 1000000000; i++ {
		v += i
	}
	fmt.Println(num, v)

	s.Done()
}
