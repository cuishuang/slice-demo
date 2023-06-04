package main

import "fmt"

func main() {

	// case1
	m := []int64{2, 3}
	fmt.Println("len of old m is ", len(m)) //2
	fmt.Println("cap of old m is ", cap(m)) //2
	fmt.Println("")

	m = append(m, 4, 5, 6)
	fmt.Println("len of m is ", len(m)) //5
	fmt.Println("cap of m is ", cap(m)) //! 6  如果要的容量是原来容量的两倍还要多, 那新的容量就是所要求的容量大小

	fmt.Println()
	fmt.Println("------")

	n := []int64{2, 3}
	fmt.Println("len of old n is ", len(n)) //2
	fmt.Println("cap of old n is ", cap(n)) //2
	fmt.Println("")

	// cas2
	fmt.Printf("切片n为:%#v,长度为%d,容量为%d,内存地址为%p\n", n, len(n), cap(n), n)
	n = append(n, 4)
	fmt.Printf("切片n为:%#v,长度为%d,容量为%d,内存地址为%p\n", n, len(n), cap(n), n)
	n = append(n, 5)
	fmt.Printf("切片n为:%#v,长度为%d,容量为%d,内存地址为%p\n", n, len(n), cap(n), n)
	n = append(n, 6)
	fmt.Printf("切片n为:%#v,长度为%d,容量为%d,内存地址为%p\n", n, len(n), cap(n), n)

	fmt.Println("len of n is ", len(n)) //5
	fmt.Println("cap of n is ", cap(n)) //! 8  如果要的容量没有原来容量两倍大, 那就扩充到原来容量的两倍.

	fmt.Println("------")

}
