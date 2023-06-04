package main

import "fmt"

func main() {

	a := make([]int, 1, 10)

	fmt.Printf("a为%#v,长度:%d,容量:%d,底层数组的内存地址:%p\n", a, len(a), cap(a), a)

	fmt.Println("--------")

	b := append(a, 2)
	fmt.Printf("b为%#v,长度:%d,容量:%d,底层数组的内存地址:%p\n", b, len(b), cap(b), b)

	fmt.Printf("此时a为%#v,长度:%d,容量:%d,底层数组的内存地址:%p\n", a, len(a), cap(a), a)

	fmt.Println("--------")

	c := append(a, 3)
	fmt.Printf("c为%#v,长度:%d,容量:%d,底层数组的内存地址:%p\n", c, len(c), cap(c), c)

	fmt.Printf("最后b为%#v,长度:%d,容量:%d,底层数组的内存地址:%p\n", b, len(b), cap(b), b)

	fmt.Printf("最后a为%#v,长度:%d,容量:%d,底层数组的内存地址:%p\n", a, len(a), cap(a), a)

	fmt.Println(a) // [0]
	fmt.Println(b) // [0, 2]

	fmt.Println(a) // 输出什么？
	fmt.Println(c) // 输出什么？

}
