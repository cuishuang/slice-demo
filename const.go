package main

import "fmt"

const GoarchArm64 = 1
const GoarchMips = 0
const GoarchMipsle = 0
const GoarchWasm = 0

const GoosIos = 0

func main() {
	// _64bit = 1 on 64-bit systems, 0 on 32-bit systems
	_64bit := 1 << (^uintptr(0) >> 63) / 2

	fmt.Println("_64bit值为：", _64bit)

	heapAddrBits := (_64bit*(1-GoarchWasm)*(1-GoosIos*GoarchArm64))*48 + (1-_64bit+GoarchWasm)*(32-(GoarchMips+GoarchMipsle)) + 33*GoosIos*GoarchArm64

	fmt.Println("heapAddrBits值为：", heapAddrBits)

	// maxAlloc is the maximum size of an allocation. On 64-bit,
	// it's theoretically possible to allocate 1<<heapAddrBits bytes. On
	// 32-bit, however, this is one less than 1<<32 because the
	// number of bytes in the address space doesn't actually fit
	// in a uintptr.
	maxAlloc := (1 << heapAddrBits) - (1-_64bit)*1

	fmt.Println("maxAlloc值为：", maxAlloc)

}
