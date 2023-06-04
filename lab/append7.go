package main

import "fmt"

func create(iterations int) []int {
	a := make([]int, 0)
	for i := 0; i < iterations; i++ {
		a = append(a, i)
	}
	return a
}

func main() {
	sliceFromLoop()

	fmt.Println("-----------------------")

	sliceFromLiteral()

}

func sliceFromLoop() {
	fmt.Printf("** NOT working as expected: **\n\n")
	i := create(11)
	fmt.Printf("切片i为:%#v,长度为%d,容量为%d,内存地址为%p\n", i, len(i), cap(i), i)
	fmt.Println("initial slice: ", i)
	j := append(i, 100)
	fmt.Printf("切片j为:%#v,长度为%d,容量为%d,内存地址为%p\n", j, len(j), cap(j), j)
	fmt.Printf("【原始切片i】为:%#v,长度为%d,容量为%d,内存地址为%p\n", i, len(i), cap(i), i)
	g := append(i, 101)
	fmt.Printf("切片g为:%#v,长度为%d,容量为%d,内存地址为%p\n", g, len(g), cap(g), g)
	fmt.Printf("【原始切片i】为:%#v,长度为%d,容量为%d,内存地址为%p\n", i, len(i), cap(i), i)
	h := append(i, 102)
	fmt.Printf("切片h为:%#v,长度为%d,容量为%d,内存地址为%p\n", h, len(h), cap(h), h)
	fmt.Printf("【原始切片i】为:%#v,长度为%d,容量为%d,内存地址为%p\n", i, len(i), cap(i), i)

	fmt.Println("最后的结果:")
	fmt.Printf("i: %v\nj: %v\ng: %v\nh:%v\n", i, j, g, h)
}

func sliceFromLiteral() {
	fmt.Printf("\n\n** working as expected: **\n")
	i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("切片i为:%#v,长度为%d,容量为%d,内存地址为%p\n", i, len(i), cap(i), i)

	fmt.Println("initial slice: ", i)
	j := append(i, 100)
	fmt.Printf("切片j为:%#v,长度为%d,容量为%d,内存地址为%p\n", j, len(j), cap(j), j)
	fmt.Printf("【原始切片i】为:%#v,长度为%d,容量为%d,内存地址为%p\n", i, len(i), cap(i), i)

	g := append(i, 101)
	fmt.Printf("切片g为:%#v,长度为%d,容量为%d,内存地址为%p\n", g, len(g), cap(g), g)
	fmt.Printf("【原始切片i】为:%#v,长度为%d,容量为%d,内存地址为%p\n", i, len(i), cap(i), i)

	h := append(i, 102)
	fmt.Printf("切片h为:%#v,长度为%d,容量为%d,内存地址为%p\n", h, len(h), cap(h), h)
	fmt.Printf("【原始切片i】为:%#v,长度为%d,容量为%d,内存地址为%p\n", i, len(i), cap(i), i)

	fmt.Println("最后的结果:")
	fmt.Printf("i: %v\nj: %v\ng: %v\nh:%v\n", i, j, g, h)
}
