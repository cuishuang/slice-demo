package main

import (
	"fmt"
)

func main() {

	a := [...]int{0, 1, 2, 3}

	fmt.Printf("数组为:%#v,长度为%d,容量为%d,内存地址为%p\n", a, len(a), cap(a), &a) // [4][0 1 2 3],4,4,地址a
	fmt.Println()

	x := a[:1]
	fmt.Printf("切片x为:%#v,长度为%d,容量为%d,底层数组的内存地址为%p\n", x, len(x), cap(x), x) // [0 1], 2, 4(容量为底层数组的长度)，地址a (也是底层数组的地址，而不是x这个切片本身的地址)

	y := a[2:]
	fmt.Printf("切片y为:%#v,长度为%d,容量为%d,底层数组的内存地址为%p\n", y, len(y), cap(y), y) // [2 3], 2, 4, 地址a (同上例)

	x = append(x, y...)
	fmt.Printf("切片x为:%#v,长度为%d,容量为%d,底层数组的内存地址为%p\n", x, len(x), cap(x), x) // [0 1 2 3], 4, 4, 地址a（依然没有扩容）

	x = append(x, y...)
	fmt.Printf("切片x为:%#v,长度为%d,容量为%d,底层数组的内存地址为%p\n", x, len(x), cap(x), x) // []

	fmt.Println(a, x)
}
