package main

import "fmt"

// append一定会改变原始slice底层数组的内存地址吗。。不一定，没有发生扩容就不需要

// https://www.zhihu.com/question/265386326/answer/2321716435

func main() {

	names := make([]int, 3)

	fmt.Printf("切片为:%#v,长度为%d,容量为%d,底层数组的内存地址为%p\n", names, len(names), cap(names), names)

	fmt.Println("-------")

	for i := 1; i < 6; i++ {

		fmt.Printf("切片为:%#v,长度为%d,容量为%d,底层数组的内存地址为%p\n", names, len(names), cap(names), names)

		names = append(names, i)
	}

	fmt.Println("-------")
	fmt.Printf("切片为:%#v,长度为%d,容量为%d\n", names, len(names), cap(names))

}
