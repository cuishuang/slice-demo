package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var sli []int64

	fmt.Println(sli)

	fmt.Println(unsafe.Sizeof(sli)) // 24, 其中unsafe.utp指针占8字节，len和cap也都占8个字节

	sli = append(sli, 1, 2, 3, 4, 5)

	fmt.Println(sli)

	fmt.Println(unsafe.Sizeof(sli)) // 24, 其中unsafe.utp指针占8字节，len和cap也都占8个字节

	m := make(map[int64]int64)

	fmt.Println(unsafe.Sizeof(m)) // 8, 只是一个指针...

}
