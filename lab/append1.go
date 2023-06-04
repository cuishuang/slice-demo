package main

import (
	"fmt"
)

func main() {

	sli := make([]string, 1)

	sli[0] = "宋江"

	fmt.Println("slice is:", sli) // ["宋江"]

	fmt.Printf("原始sli的长度%d,容量%d,内存地址%p\n", len(sli), cap(sli), sli) // 1 1 内存地址a
	f1(sli)

	fmt.Println("slice is:", sli)                                         // ["晁盖"]
	fmt.Printf("调用f1()之后sli的长度%d,容量%d,内存地址%p\n", len(sli), cap(sli), sli) // 1 1 内存地址a

}

func f1(sli1 []string) []string {

	sli1[0] = "晁盖"

	return sli1
}
