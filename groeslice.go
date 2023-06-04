package main

import "fmt"

func main() {

	const PtrSize = 4 << (^uintptr(0) >> 63) // unsafe.Sizeof(uintptr(0)) but an ideal const

	fmt.Println("uintptr(0)值为:", uintptr(0))
	fmt.Println("^uintptr(0)值为:", ^uintptr(0))           // 对0取反，在64位操作系统得到18446744073709551615； uint64(所有无符号 64 位整数的集合)，范围：0 到 18446744073709551615。
	fmt.Println("^uintptr(0) >> 63值为:", ^uintptr(0)>>63) // 1

	fmt.Println("PtrSize值为:", PtrSize) // 4左移一位，即8
}
