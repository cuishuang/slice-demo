package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var sli []int64

	fmt.Printf("%d %d %p\n", len(sli), cap(sli), &sli) // 0,0,内存地址a
	sli = append(sli, 0)
	fmt.Println(sli) // [0]

	fmt.Printf("%d %d %p\n", len(sli), cap(sli), &sli) // 1,1,内存地址a

	fmt.Println(unsafe.Sizeof(sli)) // 24, 其中unsafe.utp指针占8字节，len和cap也都占8个字节

	sli = append(sli, 1, 2, 3)

	fmt.Println(sli) // [0 1 2 3]

	fmt.Printf("%d %d %p\n", len(sli), cap(sli), &sli) // 4,4,内存地址a (底层数组没有发生变化？？)

	fmt.Println(unsafe.Sizeof(sli)) // 24, 其中unsafe.utp指针占8字节，len和cap也都占8个字节

}
