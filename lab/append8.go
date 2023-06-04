package main

import "fmt"

func main() {

	demo := make([]int, 9)

	demo2 := demo

	fmt.Printf("切片demo为:%#v,长度为%d,容量为%d,内存地址为%p\n", demo, len(demo), cap(demo), demo)
	fmt.Printf("切片demo2为:%#v,长度为%d,容量为%d,内存地址为%p\n", demo2, len(demo2), cap(demo2), demo2)

	demo3 := append(demo, 1)
	fmt.Printf("切片demo3为:%#v,长度为%d,容量为%d,内存地址为%p\n", demo3, len(demo3), cap(demo3), demo3)
	demo4 := append(demo3, 1, 2, 3)
	fmt.Printf("切片demo4为:%#v,长度为%d,容量为%d,内存地址为%p\n", demo4, len(demo4), cap(demo4), demo4)

	fmt.Println()

}
