package main

import "fmt"

//https://dashen.tech/2010/03/02/golang%E4%B9%8Bslice%E4%B8%AD%E7%9A%84%E5%B0%8Ftips/
// https://dashen.tech/2020/08/05/%E4%B8%A4%E4%B8%AAgolang%E5%B0%8F%E9%97%AE%E9%A2%98/
//https://dashen.tech/2021/03/01/%E4%B8%80%E4%B8%8D%E7%95%99%E7%A5%9E%E5%B0%B1%E6%8E%89%E5%9D%91/#map%E5%92%8Cslice%E5%8F%98%E9%87%8F%E7%9A%84%E8%B5%8B%E5%80%BC%E4%BD%9C%E7%94%A8%E8%8C%83%E5%9B%B4%E9%97%AE%E9%A2%98

// 特别留意append

func main() {

	demo := make([]int, 10)

	demo2 := demo

	fmt.Printf("切片demo为:%#v,长度为%d,容量为%d,内存地址为%p\n", demo, len(demo), cap(demo), demo)
	fmt.Printf("切片demo2为:%#v,长度为%d,容量为%d,内存地址为%p\n", demo2, len(demo2), cap(demo2), demo2)

	demo3 := append(demo, 1)
	fmt.Printf("切片demo3为:%#v,长度为%d,容量为%d,内存地址为%p\n", demo3, len(demo3), cap(demo3), demo3)
	demo4 := append(demo3, 1, 2, 3)
	fmt.Printf("切片demo4为:%#v,长度为%d,容量为%d,内存地址为%p\n", demo4, len(demo4), cap(demo4), demo4)

	fmt.Println()

	// 对于这种sli2 = append(sli1,6,6,6),如果没发生扩容，sli1和sli2底层数组一样

	// 对于sli1 = append(sli,6,6), sli2 = append(sli,8,8),即便容量一样，sli1和sli2底层数组也不一样..

	//var sli []int
	//sli := make([]int, 0)
	sli := make([]int, 11)

	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

	sli1 := append(sli, 1)
	fmt.Printf("[sli1] %v  len:%d  cap:%d  ptr:%p\n", sli1, len(sli1), cap(sli1), sli1)

	fmt.Println()

	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

	sli2 := append(sli, 1, 2)
	fmt.Printf("[sli2] %v  len:%d  cap:%d  ptr:%p\n", sli2, len(sli2), cap(sli2), sli2)

	fmt.Println()

	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

	sli3 := append(sli, 1, 2, 3)
	fmt.Printf("[sli3] %v  len:%d  cap:%d  ptr:%p\n", sli3, len(sli3), cap(sli3), sli3)
	fmt.Println()

	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

	sli4 := append(sli, 1, 2, 3, 4)
	fmt.Printf("[sli4] %v  len:%d  cap:%d  ptr:%p\n", sli4, len(sli4), cap(sli4), sli4)

	fmt.Println()

	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

	sli5 := append(sli, 1, 2, 3, 4, 5)
	fmt.Printf("[sli5] %v  len:%d  cap:%d  ptr:%p\n", sli5, len(sli5), cap(sli5), sli5)
	fmt.Println()

	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

	sli6 := append(sli, 1, 2, 3, 4, 5, 6)
	fmt.Printf("[sli6] %v  len:%d  cap:%d  ptr:%p\n", sli6, len(sli6), cap(sli6), sli6)

	fmt.Println()
	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

	sli7 := append(sli, 1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("[sli7] %v  len:%d  cap:%d  ptr:%p\n", sli7, len(sli7), cap(sli7), sli7)
	fmt.Println()

	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

	sli8 := append(sli, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Printf("[sli8] %v  len:%d  cap:%d  ptr:%p\n", sli8, len(sli8), cap(sli8), sli8)
	fmt.Println()

	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

	sli9 := append(sli, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("[sli9] %v  len:%d  cap:%d  ptr:%p\n", sli9, len(sli9), cap(sli9), sli9)
	fmt.Println()

	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

	sli10 := append(sli, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("[sli10] %v  len:%d  cap:%d  ptr:%p\n", sli10, len(sli10), cap(sli10), sli10)
	fmt.Println()

	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

	sli11 := append(sli, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Printf("[sli11] %v  len:%d  cap:%d  ptr:%p\n", sli11, len(sli11), cap(sli11), sli11)
	fmt.Println()

	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

	sli12 := append(sli, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	fmt.Printf("[sli12] %v  len:%d  cap:%d  ptr:%p\n", sli12, len(sli12), cap(sli12), sli12)
	fmt.Println()

	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

	sli13 := append(sli, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	fmt.Printf("[sli13] %v  len:%d  cap:%d  ptr:%p\n", sli13, len(sli13), cap(sli13), sli13)
	fmt.Println()

	fmt.Printf("切片sli为:%#v,长度为%d,容量为%d,内存地址为%p\n", sli, len(sli), cap(sli), sli)

}
