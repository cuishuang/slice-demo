package main

func main() {
}

func noPrepare() {
	n := 10
	src := make([]int, n)

	// 无预分配: 162 ns/op
	var dst []int

	for i := 0; i < n; i++ {
		// src[i]为默认值0
		dst = append(dst, src[i])
	}
}

func prePare() {
	n := 10
	src := make([]int, n)

	// 预分配，32.3 ns/op，提升了5倍
	dst2 := make([]int, 0, n)

	for i := 0; i < n; i++ {
		// src[i]为默认值0
		dst2 = append(dst2, src[i])
	}
}

func prePareAppend() {
	n := 10
	src := make([]int, n)

	dst2 := make([]int, 0, n)
	// 预分配+append...: 30.1 ns/op
	dst2 = append(dst2, src...)

}

func prePareCopy() {
	n := 10
	src := make([]int, n)

	dst2 := make([]int, 0, n)
	// 预分配+copy: 26.0 ns/op
	copy(dst2, src)
}
