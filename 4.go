package main

import "fmt"

func main() {

	m := []int64{2, 3}
	fmt.Println("len of old m is ", len(m)) // 2
	fmt.Println("cap of old m is ", cap(m)) // 2
	fmt.Println("")

	m = append(m, 4, 5, 6)
	fmt.Println("len of m is ", len(m)) // 5
	fmt.Println("cap of m is ", cap(m)) // 6 为什么??

	fmt.Println()
	fmt.Println("------")

	n := []int64{2, 3}
	fmt.Println("len of old n is ", len(n)) // 2
	fmt.Println("cap of old n is ", cap(n)) // 2
	fmt.Println("")

	n = append(n, 4)
	n = append(n, 5)
	n = append(n, 6)
	fmt.Println("len of n is ", len(n)) // 5
	fmt.Println("cap of n is ", cap(n)) // 8

	fmt.Println("------")

}
